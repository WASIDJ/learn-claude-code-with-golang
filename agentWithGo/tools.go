package main

import (
	"context"
	"encoding/json"
	"errors"
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
	MaxTodoItems   = 20 // 最多20条任务
)

type TodoItem struct {
	Content    string `json:"content"`
	Status     string `json:"status"`
	ActiveForm string `json:"active_form"`
}

type TodoWriteInput struct {
	Items []TodoItem `json:"items"`
}

// WORKDIR 工作目录
var WORKDIR string

// SKILLS 技能加载器
var SKILLS *SkillLoader

func init() {
	var err error
	WORKDIR, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	// 初始化技能加载器
	skillsDir := filepath.Join(WORKDIR, ".skills")
	SKILLS, err = NewSkillLoader(skillsDir)
	if err != nil {
		// 技能加载失败只输出警告，不中断程序
		fmt.Fprintf(os.Stderr, "Warning: Failed to load skills: %v\n", err)
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
	case "todo_writer":
		return executeTodoWriter(inputJSON)
	case "Task":
		return executeTask(inputJSON)
	case "Skill":
		return executeSkill(inputJSON)
	default:
		return fmt.Sprintf("未知工具: %s", toolName)
	}
}

func executeTask(inputJSON string) string {
	var input TaskInput
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	result, err := RunTask(input)
	if err != nil {
		return fmt.Sprintf("子代理执行失败: %v", err)
	}
	return result
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

func executeTodoWriter(inputJSON string) string {
	var input TodoWriteInput
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	if len(input.Items) == 0 {
		return "No todos."
	}

	if len(input.Items) > MaxTodoItems {
		return "Error: Max 20 todos allowed"
	}

	inProgressCount := 0
	for i, item := range input.Items {
		content := strings.TrimSpace(item.Content)
		status := strings.ToLower(strings.TrimSpace(item.Status))
		activeForm := strings.TrimSpace(item.ActiveForm)

		if content == "" {
			return fmt.Sprintf("Error: Item %d: content required", i)
		}
		if status == "" {
			return fmt.Sprintf("Error: Item %d: status required", i)
		}
		if activeForm == "" {
			return fmt.Sprintf("Error: Item %d: activeForm required", i)
		}
		if status != "pending" && status != "in_progress" && status != "completed" {
			return fmt.Sprintf("Error: Item %d: invalid status '%s'", i, status)
		}
		if status == "in_progress" {
			inProgressCount++
		}
		input.Items[i].Content = content
		input.Items[i].Status = status
		input.Items[i].ActiveForm = activeForm
	}

	if inProgressCount > 1 {
		return "Error: Only one task can be in_progress at a time"
	}

	return renderTodos(input.Items)
}

func renderTodos(items []TodoItem) string {
	var output strings.Builder
	completedCount := 0

	for _, item := range items {
		var statusSymbol string
		switch item.Status {
		case "completed":
			statusSymbol = "[x]"
			completedCount++
		case "in_progress":
			statusSymbol = "[>]"
		default:
			statusSymbol = "[ ]"
		}

		output.WriteString(fmt.Sprintf("%s %s", statusSymbol, item.Content))
		if item.Status == "in_progress" {
			output.WriteString(fmt.Sprintf(" <- %s", item.ActiveForm))
		}
		output.WriteString("\n")
	}

	output.WriteString(fmt.Sprintf("\n(%d/%d completed)\n", completedCount, len(items)))
	return output.String()
}

// executeSkill 执行技能调用
func executeSkill(inputJSON string) string {
	var input struct {
		Skill string `json:"skill"`
	}
	if err := json.Unmarshal([]byte(inputJSON), &input); err != nil {
		return fmt.Sprintf("解析输入失败: %v", err)
	}

	result, err := runSkill(input.Skill)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return result
}

// runSkill 加载并运行技能
// 完整内容作为 tool_result 返回
// 它会成为对话历史的一部分（user message）
func runSkill(skillName string) (string, error) {
	if SKILLS == nil {
		return "", errors.New("skills not initialized")
	}

	content, err := SKILLS.GetSkillContent(skillName)
	if err != nil {
		return "", fmt.Errorf("load skill '%s': %w", skillName, err)
	}

	result := fmt.Sprintf(`<skill-loaded name="%s">
%s
</skill-loaded>

Follow the instructions in the skill above.`, skillName, content)
	return result, nil
}
