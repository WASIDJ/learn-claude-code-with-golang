package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
)

func AgentLoop(messages []anthropic.MessageParam) ([]anthropic.MessageParam, string, error) {
	var output string
	for {
		resp, err := client.Messages.New(context.Background(), anthropic.MessageNewParams{
			Model:     "ark-code-latest",
			System:    SYSTEM,
			Messages:  messages,
			Tools:     TOOLS,
			MaxTokens: 8000,
		})
		if err != nil {
			return messages, output, err
		}

		var toolCalls []anthropic.ToolUseBlock
		for _, block := range resp.Content {
			if tb, ok := block.AsAny().(anthropic.TextBlock); ok {
				output += tb.Text
			}
			if tu, ok := block.AsAny().(anthropic.ToolUseBlock); ok {
				toolCalls = append(toolCalls, tu)
			}
		}

		messages = append(messages, resp.ToParam())

		if resp.StopReason != "tool_use" {
			return messages, output, nil
		}

		results := []map[string]any{}
		// 打印 工具调用信息并执行工具
		for _, tc := range toolCalls {
			inputJSON := tc.JSON.Input.Raw()
			inputMap := jsonToMap(inputJSON)
			output += fmt.Sprintf("\n😮‍💨> %s: %v\n", tc.Name, inputMap)

			result := ExecuteTool(tc.Name, inputJSON)
			preview := result
			// 使输出更具区分度
			if len(result) > 200 {
				preview = result[:200] + "..."
			}
			output += fmt.Sprintf("  %s\n", preview)

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

func jsonToMap(jsonStr string) map[string]any {
	var result map[string]any
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}

func Chat(prompt string, history []anthropic.MessageParam) ([]anthropic.MessageParam, string, error) {
	history = append(history, anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)))
	return AgentLoop(history)
}

func PrettyPrint(history []anthropic.MessageParam) {
	if len(history) == 0 {
		return
	}
	lastMsg := history[len(history)-1]
	for _, block := range lastMsg.Content {
		if block.OfText != nil {
			fmt.Println(block.OfText.Text)
		}
	}
}
