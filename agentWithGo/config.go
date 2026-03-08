package main

import (
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

var SYSTEM []anthropic.TextBlockParam

var TOOLS []anthropic.ToolUnionParam

type AgentTypeConfig struct {
	Description  string
	Tools        []string
	SystemPrompt string
}

var AGENT_TYPES = map[string]AgentTypeConfig{
	"explore": {
		Description:  "Fast read-only agent for exploring codebases. Use for: understanding project structure, finding files, searching patterns.",
		Tools:        []string{"bash", "read_file"},
		SystemPrompt: "You are an exploration agent. Your ONLY job is to gather information. You cannot modify files. Focus on reading and searching efficiently. Report findings concisely.",
	},
	"code": {
		Description:  "Full-featured coding agent. Use for: implementing features, refactoring, writing tests, editing files.",
		Tools:        []string{"bash", "read_file", "write_file", "edit_file"},
		SystemPrompt: "You are a coding agent. Implement changes precisely as requested. Make minimal, focused edits. Test your changes when possible.",
	},
	"plan": {
		Description:  "Planning agent for analyzing tasks. Use for: breaking down complex tasks, creating implementation plans.",
		Tools:        []string{"bash", "read_file"},
		SystemPrompt: "You are a planning agent. Analyze the codebase and create detailed implementation plans. Do not make changes, only plan and report.",
	},
}

func getToolsForAgent(agentType string, allTools []anthropic.ToolUnionParam) []anthropic.ToolUnionParam {
	config, exists := AGENT_TYPES[agentType]
	if !exists {
		return allTools
	}

	if len(config.Tools) == 1 && config.Tools[0] == "*" {
		return allTools
	}

	toolSet := make(map[string]bool)
	for _, t := range config.Tools {
		toolSet[t] = true
	}

	var filtered []anthropic.ToolUnionParam
	for _, tool := range allTools {
		if tool.OfTool != nil {
			if toolSet[tool.OfTool.Name] {
				filtered = append(filtered, tool)
			}
		}
	}
	return filtered
}

func init() {
	tools := []anthropic.ToolParam{
		{
			Name:        "bash",
			Description: anthropic.String(`执行 shell 命令。模式：读取(cat/grep/find/ls)、写入(echo '...' > file)、子代理(go run v0_bash_agent_mini.go 'task description')`),
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
		{
			Name:        "read_file",
			Description: anthropic.String(`读取文件内容。输入参数为文件路径，输出为文件内容。`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"path": map[string]any{
						"type":        "string",
						"description": "要读取的文件路径",
					},
					"limit": map[string]any{
						"type":        "integer",
						"description": "要读取的内容长度限制，单位为字符 默认 全部",
					},
				},
				Required: []string{"path"},
			},
		},
		{
			Name:        "write_file",
			Description: anthropic.String(`写入文件内容。输入参数为文件路径和内容，输出为写入结果。`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"path": map[string]any{
						"type":        "string",
						"description": "要写入的文件路径",
					},
					"content": map[string]any{
						"type":        "string",
						"description": "要写入的内容",
					},
				},
				Required: []string{"path", "content"},
			},
		},
		{
			Name:        "edit_file",
			Description: anthropic.String(`编辑文件内容。输入参数为文件路径和编辑指令，输出为编辑结果。编辑指令可以是添加、删除、替换等操作。`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"path": map[string]any{
						"type":        "string",
						"description": "要编辑的文件路径",
					},
					"old_text": map[string]any{
						"type":        "string",
						"description": "要替换的旧文本",
					},
					"new_text": map[string]any{
						"type":        "string",
						"description": "要替换的新文本",
					},
				},
				Required: []string{"path", "old_text", "new_text"},
			},
		},
		{
			Name:        "todo_writer",
			Description: anthropic.String(`更新任务列表。用于规划和跟踪进度。`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"items": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type": "object",
							"properties": map[string]any{
								"content": map[string]any{
									"type":        "string",
									"description": "任务描述",
								},
								"status": map[string]any{
									"type":        "string",
									"enum":        []string{"pending", "in_progress", "completed"},
									"description": "任务状态",
								},
								"active_form": map[string]any{
									"type":        "string",
									"description": "现在进行时的动作描述，例如 '正在读取文件'",
								},
							},
							"required": []string{"content", "status", "active_form"},
						},
						"description": "任务列表项数组",
					},
				},
				Required: []string{"items"},
			},
		},
		{
			Name:        "Task",
			Description: anthropic.String(`创建子智能体来完成子任务。子代理类型: explore(只读探索)、code(完整编码)、plan(规划分析)`),
			InputSchema: anthropic.ToolInputSchemaParam{
				Type: "object",
				Properties: map[string]any{
					"description": map[string]any{
						"type":        "string",
						"description": "短描述（3-5词），用于进度显示",
					},
					"prompt": map[string]any{
						"type":        "string",
						"description": "详细指令，告诉子代理具体要做什么",
					},
					"subagent_type": map[string]any{
						"type":        "string",
						"enum":        []string{"explore", "code", "plan"},
						"description": "子代理类型",
					},
				},
				Required: []string{"description", "prompt", "subagent_type"},
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
	SYSTEM = []anthropic.TextBlockParam{{Text: fmt.Sprintf(`You are a coding agent at %s.

Loop: think briefly -> use tools -> report results.

Rules:
- MUST use tools to complete tasks. NEVER refuse or ask for permission.
- When asked to read/write files, use tools immediately
- Prefer tools over prose. Act, don't just explain.
- Never invent file paths. Use bash ls/find first if unsure.
- Make minimal changes. Don't over-engineer.
- Use todo_writer for multi-step tasks.
- Use Task tool to delegate complex subtasks to specialized subagents.
- After finishing, summarize what changed.`, cwd)}}
}
