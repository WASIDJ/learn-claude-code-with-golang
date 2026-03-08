package main

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestGetToolsForAgent(t *testing.T) {
	tests := []struct {
		name          string
		agentType     string
		expectedTools []string
	}{
		{
			name:          "explore agent has read-only tools",
			agentType:     "explore",
			expectedTools: []string{"bash", "read_file"},
		},
		{
			name:          "code agent has all tools",
			agentType:     "code",
			expectedTools: []string{"bash", "read_file", "write_file", "edit_file"},
		},
		{
			name:          "plan agent has read-only tools",
			agentType:     "plan",
			expectedTools: []string{"bash", "read_file"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := getToolsForAgent(tt.agentType, TOOLS)

			toolNames := make(map[string]bool)
			for _, tool := range filtered {
				if tool.OfTool != nil {
					toolNames[tool.OfTool.Name] = true
				}
			}

			for _, expected := range tt.expectedTools {
				if !toolNames[expected] {
					t.Errorf("Expected tool %s for agent type %s", expected, tt.agentType)
				}
			}

			for name := range toolNames {
				found := false
				for _, expected := range tt.expectedTools {
					if name == expected {
						found = true
						break
					}
				}
				if !found && name != "todo_writer" && name != "Task" {
					t.Errorf("Unexpected tool %s for agent type %s", name, tt.agentType)
				}
			}
		})
	}
}

func TestAgentTypeConfig(t *testing.T) {
	expectedTypes := []string{"explore", "code", "plan"}

	for _, agentType := range expectedTypes {
		t.Run("agent_type_"+agentType, func(t *testing.T) {
			config, exists := AGENT_TYPES[agentType]
			if !exists {
				t.Errorf("Agent type %s not found in AGENT_TYPES", agentType)
				return
			}

			if config.Description == "" {
				t.Errorf("Agent type %s missing description", agentType)
			}
			if len(config.Tools) == 0 {
				t.Errorf("Agent type %s has no tools", agentType)
			}
			if config.SystemPrompt == "" {
				t.Errorf("Agent type %s missing system prompt", agentType)
			}
		})
	}
}

func TestTaskInputJSONParsing(t *testing.T) {
	jsonInput := `{
		"description": "explore project",
		"prompt": "Find all Go files in the project",
		"subagent_type": "explore"
	}`

	var input TaskInput
	if err := json.Unmarshal([]byte(jsonInput), &input); err != nil {
		t.Errorf("Failed to parse Task input JSON: %v", err)
	}

	if input.Description != "explore project" {
		t.Errorf("Expected description 'explore project', got '%s'", input.Description)
	}
	if input.SubagentType != "explore" {
		t.Errorf("Expected subagent_type 'explore', got '%s'", input.SubagentType)
	}
}

func TestExecuteTask(t *testing.T) {
	tests := []struct {
		name        string
		input       TaskInput
		expectError string
	}{
		{
			name: "unknown agent type returns error",
			input: TaskInput{
				Description:  "test",
				Prompt:       "test prompt",
				SubagentType: "unknown_type",
			},
			expectError: "unknown subagent type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputJSON, _ := json.Marshal(tt.input)
			result := executeTask(string(inputJSON))

			if tt.expectError != "" && !strings.Contains(result, tt.expectError) {
				t.Errorf("Expected error containing '%s', got '%s'", tt.expectError, result)
			}
		})
	}
}

func TestSubagentProgress(t *testing.T) {
	progress := &SubagentProgress{
		Description: "test task",
		StartTime:   time.Now(),
	}

	if progress.ToolCalls != 0 {
		t.Errorf("Expected initial ToolCalls to be 0, got %d", progress.ToolCalls)
	}

	progress.Update("read_file")
	if progress.ToolCalls != 1 {
		t.Errorf("Expected ToolCalls to be 1 after update, got %d", progress.ToolCalls)
	}

	progress.Update("bash")
	if progress.ToolCalls != 2 {
		t.Errorf("Expected ToolCalls to be 2 after second update, got %d", progress.ToolCalls)
	}
}

func TestToolsWhiteListFiltering(t *testing.T) {
	exploreTools := getToolsForAgent("explore", TOOLS)

	for _, tool := range exploreTools {
		if tool.OfTool == nil {
			continue
		}
		name := tool.OfTool.Name
		if name == "write_file" || name == "edit_file" {
			t.Errorf("explore agent should not have write tool %s", name)
		}
	}

	codeTools := getToolsForAgent("code", TOOLS)
	toolNames := make(map[string]bool)
	for _, tool := range codeTools {
		if tool.OfTool != nil {
			toolNames[tool.OfTool.Name] = true
		}
	}

	for _, expected := range []string{"bash", "read_file", "write_file", "edit_file"} {
		if !toolNames[expected] {
			t.Errorf("code agent should have tool %s", expected)
		}
	}
}
