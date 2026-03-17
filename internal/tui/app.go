// Package tui implements the bubbletea-based terminal user interface.
package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tab represents the active content tab.
type Tab int

const (
	TabPreview Tab = iota
	TabDiff
	TabTerminal
)

func (t Tab) String() string {
	switch t {
	case TabPreview:
		return "Preview"
	case TabDiff:
		return "Diff"
	case TabTerminal:
		return "Terminal"
	default:
		return "Unknown"
	}
}

// tickMsg triggers periodic updates (metadata, capture).
type tickMsg time.Time

// App is the main bubbletea model.
type App struct {
	width  int
	height int

	activeTab     Tab
	selectedIndex int

	// Set to true when we need to quit
	quitting bool
}

// NewApp creates a new TUI application model.
func NewApp() App {
	return App{
		activeTab: TabPreview,
	}
}

// Init implements tea.Model.
func (a App) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update implements tea.Model.
func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			a.quitting = true
			return a, tea.Quit
		case "tab":
			a.activeTab = (a.activeTab + 1) % 3
		case "j", "down":
			a.selectedIndex++
		case "k", "up":
			if a.selectedIndex > 0 {
				a.selectedIndex--
			}
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height

	case tickMsg:
		return a, tickCmd()
	}

	return a, nil
}

// View implements tea.Model.
func (a App) View() string {
	if a.quitting {
		return ""
	}
	if a.width == 0 || a.height == 0 {
		return "Initializing..."
	}

	// Calculate layout dimensions
	leftWidth := max(int(float64(a.width)*LeftPanelRatio), MinLeftWidth)
	rightWidth := a.width - leftWidth - 2 // account for borders
	contentHeight := a.height - MenuHeight - 2 // menu + borders

	// Render panels
	left := a.renderLeftPanel(leftWidth, contentHeight)
	right := a.renderRightPanel(rightWidth, contentHeight)
	menu := a.renderMenu()

	// Compose layout
	panels := lipgloss.JoinHorizontal(lipgloss.Top, left, right)
	return lipgloss.JoinVertical(lipgloss.Left, panels, menu)
}

func (a App) renderLeftPanel(width, height int) string {
	content := "No instances yet.\n\nPress 'n' to create one."
	return LeftPanelStyle.
		Width(width).
		Height(height).
		Render(content)
}

func (a App) renderRightPanel(width, height int) string {
	// Tab bar
	tabs := a.renderTabBar()

	content := "Select an instance to preview."
	return RightPanelStyle.
		Width(width).
		Height(height).
		Render(tabs + "\n" + content)
}

func (a App) renderTabBar() string {
	var tabs []string
	for _, tab := range []Tab{TabPreview, TabDiff, TabTerminal} {
		if tab == a.activeTab {
			tabs = append(tabs, ActiveTabStyle.Render(tab.String()))
		} else {
			tabs = append(tabs, InactiveTabStyle.Render(tab.String()))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (a App) renderMenu() string {
	items := []struct{ key, desc string }{
		{"n", "new"},
		{"D", "kill"},
		{"Enter", "open"},
		{"p", "push"},
		{"tab", "switch tab"},
		{"?", "help"},
		{"q", "quit"},
	}

	var parts []string
	for _, item := range items {
		parts = append(parts,
			MenuKeyStyle.Render(item.key)+" "+MenuDescStyle.Render(item.desc))
	}

	return MenuStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top,
		joinWithSep(parts, " │ ")))
}

func joinWithSep(parts []string, sep string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += sep
		}
		result += p
	}
	return result
}

// Run starts the TUI application.
func Run() error {
	p := tea.NewProgram(
		NewApp(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	_, err := p.Run()
	return err
}
