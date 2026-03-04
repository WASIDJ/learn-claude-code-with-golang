package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	MaxOutputSize  = 50000 // 50KB output limit
	CommandTimeout = 120 * time.Second
)

// WORKDIR 工作目录
var WORKDIR string

func init() {
	var err error
	WORKDIR, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

// ExecuteTool 执行工具调用
func ExecuteTool(toolName string, inputJSON string) string {
	switch toolName {
	case "bash":
		return executeBash(inputJSON)
	case "read_file":
		return executeReadFile(inputJSON)
	case "write_file":
		return executeWriteFile(inputJSON)
	case "edit_file":
		return executeEditFile(inputJSON)
	default:
		return fmt.Sprintf("未知工具: %s", toolName)
	}
}

// safePath 确保路径在 WORKDIR 内，防止逃逸
func safePath(p string) (string, error) {
	absPath := filepath.Join(WORKDIR, p)
	resolved, err := filepath.Abs(absPath)
	if err != nil {
		return "", err
	}

	workdirAbs, err := filepath.Abs(WORKDIR)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(resolved, workdirAbs) {
		return "", fmt.Errorf("path escapes workspace: %s", p)
	}
	return resolved, nil
}

func executeBash(inputJSON string) string {
	var input struct {
		Command string `json:"command"`
	}
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	command := input.Command

	// 安全检查：阻止危险命令
	dangerousPatterns := []string{
		"rm -rf /",
		"sudo",
		"shutdown",
		"reboot",
		"> /dev/",
	}
	for _, pattern := range dangerousPatterns {
		if strings.Contains(command, pattern) {
			return "Error: Dangerous command blocked"
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), CommandTimeout)
	defer cancel()

	out, err := exec.CommandContext(ctx, "bash", "-lc", command).CombinedOutput()
	res := string(out)
	if err != nil {
		res += "\n" + err.Error()
	}

	// 输出截断
	if len(res) > MaxOutputSize {
		res = res[:MaxOutputSize]
	}
	if res == "" {
		return "(no output)"
	}
	return res
}

func executeReadFile(inputJSON string) string {
	var input struct {
		Path  string `json:"path"`
		Limit int    `json:"limit"`
	}
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	// 安全路径检查
	safePath, err := safePath(input.Path)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	data, err := os.ReadFile(safePath)
	if err != nil {
		return fmt.Sprintf("读取文件失败: %v", err)
	}

	content := string(data)

	// 如果有 limit，限制行数
	if input.Limit > 0 {
		lines := strings.Split(content, "\n")
		if len(lines) > input.Limit {
			lines = lines[:input.Limit]
			lines = append(lines, fmt.Sprintf("... (%d more lines)", len(strings.Split(content, "\n"))-input.Limit))
			content = strings.Join(lines, "\n")
		}
	}

	// 输出截断
	if len(content) > MaxOutputSize {
		content = content[:MaxOutputSize]
	}
	return content
}

func executeWriteFile(inputJSON string) string {
	var input struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	// 安全路径检查
	safePath, err := safePath(input.Path)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	// 创建父目录
	if err := os.MkdirAll(filepath.Dir(safePath), 0755); err != nil {
		return fmt.Sprintf("创建目录失败: %v", err)
	}

	if err := os.WriteFile(safePath, []byte(input.Content), 0644); err != nil {
		return fmt.Sprintf("写入文件失败: %v", err)
	}
	return fmt.Sprintf("Wrote %d bytes to %s", len(input.Content), input.Path)
}

func executeEditFile(inputJSON string) string {
	var input struct {
		Path    string `json:"path"`
		OldText string `json:"old_text"`
		NewText string `json:"new_text"`
	}
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	// 安全路径检查
	safePath, err := safePath(input.Path)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	// 读取原文件
	data, err := os.ReadFile(safePath)
	if err != nil {
		return fmt.Sprintf("读取文件失败: %v", err)
	}
	content := string(data)

	// 检查旧文本是否存在
	if !strings.Contains(content, input.OldText) {
		return fmt.Sprintf("Error: Text not found in %s", input.Path)
	}

	// 只替换第一个匹配项（安全）
	newContent := strings.Replace(content, input.OldText, input.NewText, 1)

	// 写回文件
	if err := os.WriteFile(safePath, []byte(newContent), 0644); err != nil {
		return fmt.Sprintf("写入文件失败: %v", err)
	}
	return fmt.Sprintf("Edited %s", input.Path)
}
