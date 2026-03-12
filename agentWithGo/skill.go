package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Skill struct {
	Name        string
	Description string
	Body        string
	Path        string
	Dir         string
}

type SkillLoader struct {
	skills map[string]Skill
}

func NewSkillLoader(skillsDir string) (*SkillLoader, error) {
	loader := &SkillLoader{
		skills: make(map[string]Skill),
	}

	if err := loader.loadSkills(skillsDir); err != nil {
		return nil, err
	}

	return loader, nil
}

func (s *SkillLoader) loadSkills(skillsDir string) error {
	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		return fmt.Errorf("failed to read skills directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		skillPath := filepath.Join(skillsDir, entry.Name(), "SKILL.md")
		if _, err := os.Stat(skillPath); os.IsNotExist(err) {
			continue
		}

		skill, err := parseSkillMD(skillPath)
		if err != nil {
			return fmt.Errorf("failed to parse skill %s: %w", entry.Name(), err)
		}

		s.skills[skill.Name] = skill
	}

	return nil
}

func parseSkillMD(path string) (Skill, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Skill{}, fmt.Errorf("failed to read file: %w", err)
	}

	re := regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---\s*\n(.*)$`)
	matches := re.FindStringSubmatch(string(content))

	if len(matches) < 3 {
		return Skill{}, fmt.Errorf("invalid SKILL.md format: missing frontmatter")
	}

	yamlContent := matches[1]
	body := matches[2]

	name := extractYAMLField(yamlContent, "name")
	description := extractYAMLField(yamlContent, "description")

	dir := filepath.Base(filepath.Dir(path))
	if name == "" {
		name = dir
	}

	return Skill{
		Name:        name,
		Description: description,
		Body:        body,
		Path:        path,
		Dir:         dir,
	}, nil
}

func extractYAMLField(content, field string) string {
	re := regexp.MustCompile(fmt.Sprintf(`(?m)^%s:\s*(.*)$`, field))
	matches := re.FindStringSubmatch(content)
	if len(matches) < 2 {
		return ""
	}
	return strings.TrimSpace(matches[1])
}

func (s *SkillLoader) GetDescriptions() string {
	var lines []string
	for name, skill := range s.skills {
		lines = append(lines, fmt.Sprintf("- %s: %s", name, skill.Description))
	}
	return strings.Join(lines, "\n")
}

func (s *SkillLoader) GetSkillContent(name string) (string, error) {
	skill, exists := s.skills[name]
	if !exists {
		return "", fmt.Errorf("skill not found: %s", name)
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# Skill: %s\n\n", skill.Name))
	sb.WriteString(skill.Body)
	sb.WriteString("\n\n")

	return sb.String(), nil
}

func (s *SkillLoader) GetSkill(name string) (Skill, bool) {
	skill, exists := s.skills[name]
	return skill, exists
}

func (s *SkillLoader) ListSkills() []string {
	names := make([]string, 0, len(s.skills))
	for name := range s.skills {
		names = append(names, name)
	}
	return names
}
