package installer

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kwiki/kwiki/tools"
	"gopkg.in/yaml.v3"
)

type envFile struct {
	Selections []struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"selections"`
}

func Export(path string) error {
	// nothing to export without a prior run — write a template instead
	var ef envFile
	for _, t := range tools.Registry {
		ef.Selections = append(ef.Selections, struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{Name: t.Name, Version: t.Versions[0]})
	}
	data, err := yaml.Marshal(ef)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Import(path string) ([]Selection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var ef envFile
	if err := yaml.Unmarshal(data, &ef); err != nil {
		return nil, err
	}
	toolMap := map[string]tools.Tool{}
	for _, t := range tools.Registry {
		toolMap[strings.ToLower(t.Name)] = t
	}
	var selections []Selection
	for _, s := range ef.Selections {
		t, ok := toolMap[strings.ToLower(s.Name)]
		if !ok {
			fmt.Printf("⚠  Unknown tool: %s — skipping\n", s.Name)
			continue
		}
		selections = append(selections, Selection{Tool: t, Version: s.Version})
	}
	return selections, nil
}

func ListTools() {
	nameStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	verStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	fmt.Println()
	for _, t := range tools.Registry {
		fmt.Printf("  %s  %s\n",
			nameStyle.Render(fmt.Sprintf("%-20s", t.Name)),
			verStyle.Render(strings.Join(t.Versions, " | ")),
		)
	}
	fmt.Println()
}
