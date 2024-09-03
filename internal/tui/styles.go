package tui

import "github.com/charmbracelet/lipgloss"

var (
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Background(lipgloss.Color("57")).Render
	selectedStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(lipgloss.Color("57")).Render
	inactiveStyle = lipgloss.NewStyle().PaddingLeft(1).PaddingRight(1).Foreground(lipgloss.Color("#393939")).Render
)

