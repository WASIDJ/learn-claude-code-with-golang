package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestTodoManagerMaxItems 测试最多20条任务限制
func TestTodoManagerMaxItems(t *testing.T) {
	// 创建超过20条任务
	items := []TodoItem{}
	for i := 0; i < 25; i++ {
		items = append(items, TodoItem{
			Content:    "Task " + string(rune('A'+i%26)),
			Status:     "pending",
			ActiveForm: "Doing task",
		})
	}

	input := TodoWriteInput{Items: items}
	inputJSON, _ := json.Marshal(input)

	result := executeTodoWriter(string(inputJSON))

	// 应该返回错误信息
	if !strings.Contains(result, "Max 20 todos allowed") {
		t.Errorf("Expected error for more than 20 items, got: %s", result)
	}
}

// TestTodoManagerSingleInProgress 测试同一时刻只允许一个 in_progress
func TestTodoManagerSingleInProgress(t *testing.T) {
	items := []TodoItem{
		{Content: "Task 1", Status: "in_progress", ActiveForm: "Doing task 1"},
		{Content: "Task 2", Status: "in_progress", ActiveForm: "Doing task 2"},
		{Content: "Task 3", Status: "pending", ActiveForm: "Doing task 3"},
	}

	input := TodoWriteInput{Items: items}
	inputJSON, _ := json.Marshal(input)

	result := executeTodoWriter(string(inputJSON))

	// 应该返回错误信息
	if !strings.Contains(result, "Only one task can be in_progress") {
		t.Errorf("Expected error for multiple in_progress items, got: %s", result)
	}
}

// TestTodoManagerRequiredFields 测试必填字段检查
func TestTodoManagerRequiredFields(t *testing.T) {
	tests := []struct {
		name        string
		items       []TodoItem
		expectError string
	}{
		{
			name: "missing content",
			items: []TodoItem{
				{Content: "", Status: "pending", ActiveForm: "Doing task"},
			},
			expectError: "content required",
		},
		{
			name: "missing activeForm",
			items: []TodoItem{
				{Content: "Task 1", Status: "pending", ActiveForm: ""},
			},
			expectError: "activeForm required",
		},
		{
			name: "invalid status",
			items: []TodoItem{
				{Content: "Task 1", Status: "invalid", ActiveForm: "Doing task"},
			},
			expectError: "invalid status",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := TodoWriteInput{Items: tt.items}
			inputJSON, _ := json.Marshal(input)

			result := executeTodoWriter(string(inputJSON))

			if !strings.Contains(result, tt.expectError) {
				t.Errorf("Expected error containing '%s', got: %s", tt.expectError, result)
			}
		})
	}
}

// TestTodoManagerRenderFormat 测试渲染输出格式
func TestTodoManagerRenderFormat(t *testing.T) {
	items := []TodoItem{
		{Content: "Completed task", Status: "completed", ActiveForm: "Done"},
		{Content: "In progress task", Status: "in_progress", ActiveForm: "Working on task"},
		{Content: "Pending task", Status: "pending", ActiveForm: "Will do task"},
	}

	input := TodoWriteInput{Items: items}
	inputJSON, _ := json.Marshal(input)

	result := executeTodoWriter(string(inputJSON))

	// 检查格式
	if !strings.Contains(result, "[x] Completed task") {
		t.Errorf("Expected '[x] Completed task' in output, got: %s", result)
	}
	if !strings.Contains(result, "[>] In progress task <- Working on task") {
		t.Errorf("Expected '[>] In progress task <- Working on task' in output, got: %s", result)
	}
	if !strings.Contains(result, "[ ] Pending task") {
		t.Errorf("Expected '[ ] Pending task' in output, got: %s", result)
	}
	if !strings.Contains(result, "(1/3 completed)") {
		t.Errorf("Expected '(1/3 completed)' in output, got: %s", result)
	}
}

// TestTodoManagerEmptyList 测试空列表
func TestTodoManagerEmptyList(t *testing.T) {
	items := []TodoItem{}

	input := TodoWriteInput{Items: items}
	inputJSON, _ := json.Marshal(input)

	result := executeTodoWriter(string(inputJSON))

	if !strings.Contains(result, "No todos") {
		t.Errorf("Expected 'No todos' for empty list, got: %s", result)
	}
}

// TestTodoManagerValidInput 测试有效输入
func TestTodoManagerValidInput(t *testing.T) {
	items := []TodoItem{
		{Content: "Refactor auth module", Status: "completed", ActiveForm: "Refactoring"},
		{Content: "Add unit tests", Status: "in_progress", ActiveForm: "Adding tests for auth module"},
		{Content: "Update documentation", Status: "pending", ActiveForm: "Updating docs"},
	}

	input := TodoWriteInput{Items: items}
	inputJSON, _ := json.Marshal(input)

	result := executeTodoWriter(string(inputJSON))

	// 验证不包含错误信息
	if strings.Contains(result, "Error") || strings.Contains(result, "error") {
		t.Errorf("Unexpected error for valid input: %s", result)
	}

	// 验证渲染格式正确
	if !strings.Contains(result, "[x]") || !strings.Contains(result, "[>]") || !strings.Contains(result, "[ ]") {
		t.Errorf("Missing expected status markers in output: %s", result)
	}
}

// TestTodoManagerStatusValidation 测试状态值验证
func TestTodoManagerStatusValidation(t *testing.T) {
	validStatuses := []string{"pending", "in_progress", "completed"}

	for _, status := range validStatuses {
		t.Run("valid_status_"+status, func(t *testing.T) {
			items := []TodoItem{
				{Content: "Task", Status: status, ActiveForm: "Doing"},
			}

			input := TodoWriteInput{Items: items}
			inputJSON, _ := json.Marshal(input)

			result := executeTodoWriter(string(inputJSON))

			if strings.Contains(result, "invalid status") {
				t.Errorf("Status '%s' should be valid, got: %s", status, result)
			}
		})
	}
}

// TestTodoManagerJSONParsing 测试 JSON 解析
func TestTodoManagerJSONParsing(t *testing.T) {
	// 测试 snake_case 字段名（API 规范）
	jsonInput := `{"items": [{"content": "Task 1", "status": "pending", "active_form": "Doing task 1"}]}`

	result := executeTodoWriter(jsonInput)

	// 应该能正常解析
	if strings.Contains(result, "解析输入失败") {
		t.Errorf("Failed to parse valid JSON: %s", result)
	}
}
