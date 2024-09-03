package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	RANDOM = 0
	SAVED = 1
	MODIFIED = 2
)

func wordsList(selected int, list []string) string {
	var b strings.Builder

	for i, word := range list {
		if i == selected {
			b.WriteString(selectedItemStyle(word) + "\n")
		} else {
			b.WriteString(" " + word + "\n")
		}
	}
	
	return b.String()
}

func _random(m model) string {
	return wordsList(m.selected, m.words)
}

func _saved(m model) string {
	return wordsList(m.selected, m.savedWords)
}

func _modified(m model) string {
	return wordsList(m.selected, m.modifiedWords)
}


func page(m model) string {
	var s string
	switch m.page {
	case RANDOM:
		s = _random(m)
	case SAVED:
		s = _saved(m)
	case MODIFIED:
		s = _modified(m)
	}

	return lipgloss.NewStyle().PaddingLeft(1).Render(s)
}
