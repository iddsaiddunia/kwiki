package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kwiki/kwiki/installer"
	"github.com/kwiki/kwiki/tools"
)

// ── styles ────────────────────────────────────────────────────────────────────

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).MarginBottom(1)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	dimStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	boxStyle      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("205")).Padding(0, 1)
)

// ── stages ────────────────────────────────────────────────────────────────────

type stage int

const (
	stageSelect  stage = iota // checklist of tools
	stageVersion              // pick version for each selected tool
	stageConfirm              // confirm before install
	stageDone                 // finished
)

// ── list item ─────────────────────────────────────────────────────────────────

type toolItem struct {
	tool     tools.Tool
	selected bool
}

func (t toolItem) Title() string {
	if t.selected {
		return selectedStyle.Render("✔ " + t.tool.Name)
	}
	return normalStyle.Render("  " + t.tool.Name)
}
func (t toolItem) Description() string {
	return dimStyle.Render("versions: " + strings.Join(t.tool.Versions, ", "))
}
func (t toolItem) FilterValue() string { return t.tool.Name }

// ── model ─────────────────────────────────────────────────────────────────────

type Model struct {
	stage        stage
	list         list.Model
	items        []toolItem
	versionIdx   int   // which selected tool we're picking version for
	versionCur   int   // cursor in version list
	versions     []installer.Selection
	width        int
	height       int
	err          string
}

func New() Model {
	items := make([]toolItem, len(tools.Registry))
	listItems := make([]list.Item, len(tools.Registry))
	for i, t := range tools.Registry {
		items[i] = toolItem{tool: t}
		listItems[i] = items[i]
	}

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = selectedStyle
	delegate.Styles.SelectedDesc = dimStyle

	l := list.New(listItems, delegate, 60, 20)
	l.Title = "Select tools to install"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle

	return Model{list: l, items: items}
}

func (m Model) Init() tea.Cmd { return nil }

// ── update ────────────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.list.SetSize(msg.Width-4, msg.Height-8)

	case tea.KeyMsg:
		switch m.stage {

		case stageSelect:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case " ":
				i := m.list.Index()
				m.items[i].selected = !m.items[i].selected
				m.list.SetItem(i, m.items[i])
				return m, nil
			case "enter", "i":
				if !m.anySelected() {
					m.err = "Please select at least one tool (SPACE to select)"
					return m, nil
				}
				m.err = ""
				m.stage = stageVersion
				m.versionIdx = 0
				m.versionCur = 0
				m.versions = nil
				return m, nil
			case "a":
				all := !m.anySelected()
				for i := range m.items {
					m.items[i].selected = all
					m.list.SetItem(i, m.items[i])
				}
				return m, nil
			}

		case stageVersion:
			selected := m.selectedTools()
			cur := selected[m.versionIdx]
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.versionCur > 0 {
					m.versionCur--
				}
			case "down", "j":
				if m.versionCur < len(cur.tool.Versions)-1 {
					m.versionCur++
				}
			case "enter", " ":
				m.versions = append(m.versions, installer.Selection{
					Tool:    cur.tool,
					Version: cur.tool.Versions[m.versionCur],
				})
				m.versionIdx++
				m.versionCur = 0
				if m.versionIdx >= len(selected) {
					m.stage = stageConfirm
				}
			case "b":
				if m.versionIdx > 0 {
					m.versionIdx--
					m.versions = m.versions[:len(m.versions)-1]
					m.versionCur = 0
				} else {
					m.stage = stageSelect
				}
			}

		case stageConfirm:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "y", "enter":
				m.stage = stageDone
				return m, tea.Quit
			case "n", "b":
				m.stage = stageSelect
				m.versions = nil
				m.versionIdx = 0
			}
		}
	}

	if m.stage == stageSelect {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
	return m, nil
}

// ── view ──────────────────────────────────────────────────────────────────────

func (m Model) View() string {
	switch m.stage {

	case stageSelect:
		help := helpStyle.Render("SPACE select • A toggle all • ENTER confirm • / filter • Q quit")
		errMsg := ""
		if m.err != "" {
			errMsg = "\n" + errorStyle.Render("⚠  "+m.err)
		}
		return "\n" + m.list.View() + errMsg + "\n" + help

	case stageVersion:
		selected := m.selectedTools()
		cur := selected[m.versionIdx]
		var sb strings.Builder
		sb.WriteString(titleStyle.Render(fmt.Sprintf("Pick version for %s (%d/%d)", cur.tool.Name, m.versionIdx+1, len(selected))))
		sb.WriteString("\n\n")
		for i, v := range cur.tool.Versions {
			if i == m.versionCur {
				sb.WriteString(selectedStyle.Render("  ▶ "+v) + "\n")
			} else {
				sb.WriteString(normalStyle.Render("    "+v) + "\n")
			}
		}
		sb.WriteString(helpStyle.Render("\n↑/↓ navigate • ENTER select • B back • Q quit"))
		return boxStyle.Render(sb.String())

	case stageConfirm:
		var sb strings.Builder
		sb.WriteString(titleStyle.Render("Ready to install:") + "\n\n")
		for _, s := range m.versions {
			sb.WriteString(selectedStyle.Render(fmt.Sprintf("  ✔ %s %s", s.Tool.Name, s.Version)) + "\n")
		}
		sb.WriteString(helpStyle.Render("\nY/ENTER to install • N/B to go back • Q quit"))
		return boxStyle.Render(sb.String())

	case stageDone:
		return ""
	}
	return ""
}

// ── helpers ───────────────────────────────────────────────────────────────────

func (m Model) anySelected() bool {
	for _, it := range m.items {
		if it.selected {
			return true
		}
	}
	return false
}

func (m Model) selectedTools() []toolItem {
	var out []toolItem
	for _, it := range m.items {
		if it.selected {
			out = append(out, it)
		}
	}
	return out
}

func (m Model) Selections() []installer.Selection { return m.versions }
