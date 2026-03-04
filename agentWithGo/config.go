package main

import (
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

// SYSTEM 系统提示词
var SYSTEM []anthropic.TextBlockParam

// TOOLS 可用工具列表
var TOOLS []anthropic.ToolUnionParam

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
- Prefer tools over prose. Act, don't just explain.
- Never invent file paths. Use bash ls/find first if unsure.
- Make minimal changes. Don't over-engineer.
- After finishing, summarize what changed.`, cwd)}}
}
