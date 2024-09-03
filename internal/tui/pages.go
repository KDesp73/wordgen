package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	RANDOM = 0
	SAVED = 1
)

func _random(m model) string {
	var b strings.Builder

	for i, word := range m.words {
		if i == m.selected {
			b.WriteString(selectedItemStyle(word) + "\n")
		} else {
			b.WriteString(" " + word + "\n")
		}
	}
	
	return b.String()
}

func _saved(m model) string {
	var b strings.Builder

	for i, word := range m.savedWords {
		if i == m.selected {
			b.WriteString(selectedItemStyle(word) + "\n")
		} else {
			b.WriteString(" " + word + "\n")
		}
	}
	
	return b.String()
}


func page(m model) string {
	var s string
	switch m.page {
	case RANDOM:
		s = _random(m)
	case SAVED:
		s = _saved(m)
	}

	return lipgloss.NewStyle().PaddingLeft(1).Render(s)
}
