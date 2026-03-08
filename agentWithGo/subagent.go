package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
)

type TaskInput struct {
	Description  string `json:"description"`
	Prompt       string `json:"prompt"`
	SubagentType string `json:"subagent_type"`
}

type SubagentProgress struct {
	Description string
	ToolCalls   int
	StartTime   time.Time
}

func (p *SubagentProgress) Update(toolCall string) {
	p.ToolCalls++
	elapsed := time.Since(p.StartTime).Seconds()
	fmt.Printf("\r@ Task(%s): %s... (%d tool calls, %.1fs)", p.Description, toolCall, p.ToolCalls, elapsed)
}

func (p *SubagentProgress) Complete() {
	elapsed := time.Since(p.StartTime).Seconds()
	fmt.Printf("\r@ Task(%s): completed: %d tool calls in %.1fs\n", p.Description, p.ToolCalls, elapsed)
}

func RunTask(input TaskInput) (string, error) {
	agentConfig, exists := AGENT_TYPES[input.SubagentType]
	if !exists {
		return "", fmt.Errorf("unknown subagent type: %s", input.SubagentType)
	}

	progress := &SubagentProgress{
		Description: input.Description,
		StartTime:   time.Now(),
	}

	subSystem := fmt.Sprintf("You are a %s subagent.\n\n%s\n\nWork autonomously. Complete the task and report results concisely.", input.SubagentType, agentConfig.SystemPrompt)

	subTools := getToolsForAgent(input.SubagentType, TOOLS)

	subMessages := []anthropic.MessageParam{
		anthropic.NewUserMessage(anthropic.NewTextBlock(input.Prompt)),
	}

	result, err := subagentLoop(subMessages, subSystem, subTools, progress)
	if err != nil {
		return "", err
	}

	progress.Complete()
	return result, nil
}

func subagentLoop(messages []anthropic.MessageParam, system string, tools []anthropic.ToolUnionParam, progress *SubagentProgress) (string, error) {
	var finalText strings.Builder

	for {
		resp, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
			Model:     "ark-code-latest",
			System:    []anthropic.TextBlockParam{{Text: system}},
			Messages:  messages,
			Tools:     tools,
			MaxTokens: 4096,
		})
		if err != nil {
			return "", err
		}

		var toolCalls []anthropic.ToolUseBlock
		for _, block := range resp.Content {
			if tb, ok := block.AsAny().(anthropic.TextBlock); ok {
				finalText.WriteString(tb.Text)
			}
			if tu, ok := block.AsAny().(anthropic.ToolUseBlock); ok {
				toolCalls = append(toolCalls, tu)
			}
		}

		messages = append(messages, resp.ToParam())

		if resp.StopReason != "tool_use" {
			return finalText.String(), nil
		}

		results := []map[string]any{}
		for _, tc := range toolCalls {
			progress.Update(tc.Name)

			result := ExecuteTool(tc.Name, tc.JSON.Input.Raw())
			results = append(results, map[string]any{
				"type":        "tool_result",
				"tool_use_id": tc.ID,
				"content":     result,
			})
		}

		messages = append(messages, anthropic.NewUserMessage(
			anthropic.NewTextBlock(fmt.Sprintf("%v", results)),
		))
	}
}
