package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var client = anthropic.NewClient(
	option.WithAPIKey(os.Getenv("LARK_API_KEY")),
	option.WithBaseURL("https://ark.cn-beijing.volces.com/api/coding"),
)
var (
	SYSTEM []anthropic.TextBlockParam
	TOOLS  []anthropic.ToolUnionParam
)

func init() {
	tools := []anthropic.ToolParam{
		{
			Name: "bash",
			Description: anthropic.String(`执行 shell 命令。模式：
- 读取: cat/grep/find/ls
- 写入: echo '...' > file
- 子代理: go run v0_bash_agent_mini.go 'task description`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"command": map[string]any{
						"type":        "string",
						"description": "要执行的 shell 命令",
					},
				},
				Required: []string{"command"},
			},
		},
	}
	for _, t := range tools {
		TOOLS = append(TOOLS, anthropic.ToolUnionParam{OfTool: &t})
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	SYSTEM = []anthropic.TextBlockParam{{Text: fmt.Sprintf("CLI agent at %s. Use bash. Spawn subagent for complex tasks.", cwd)}}
}
func chat(prompt string, history []anthropic.MessageParam) ([]anthropic.MessageParam, error) {
	history = append(history, anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)))

	for {
		resp, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
			Model:     "ark-code-latest",
			System:    SYSTEM,
			Messages:  history,
			Tools:     TOOLS,
			MaxTokens: 8000,
		})
		if err != nil {
			return history, err
		}

		var reply string
		for _, block := range resp.Content {
			if tb, ok := block.AsAny().(anthropic.TextBlock); ok {
				reply += tb.Text
			}
		}
		history = append(history, resp.ToParam())

		if resp.StopReason != "tool_use" {
			return history, nil
		}

		var toolResults []anthropic.ContentBlockParamUnion
		for _, block := range resp.Content {
			if tb, ok := block.AsAny().(anthropic.ToolUseBlock); ok {
				inputJSON := tb.JSON.Input.Raw()
				var input struct {
					Command string `json:"command"`
				}
				json.Unmarshal([]byte(inputJSON), &input)
				ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
				out, err := exec.CommandContext(ctx, "bash", "-lc", input.Command).CombinedOutput()
				cancel()
				res := string(out)
				if err != nil {
					res += "\n" + err.Error()
				}
				toolResults = append(toolResults, anthropic.NewToolResultBlock(tb.ID, res, false))
			}
		}
		if len(toolResults) > 0 {
			history = append(history, anthropic.NewUserMessage(toolResults...))
		}
	}
}

func main() {
	hist := []anthropic.MessageParam{}
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" || input == "exit" || input == "quit" {
			break
		}
		hist, err := chat(input, hist)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		lastMsg := hist[len(hist)-1]
		for _, block := range lastMsg.Content {
			if block.OfText != nil {
				fmt.Println(block.OfText.Text)
			}
		}
		fmt.Print("> ")
	}
}
