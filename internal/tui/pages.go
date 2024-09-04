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

func wordsList(m model) string {
	var b strings.Builder

	var currentWords []string
	switch m.page {
	case SAVED: 
		currentWords = m.savedWords
	case RANDOM:
		currentWords = m.words
	case MODIFIED:
		currentWords = m.modifiedWords
	}

	for i, word := range currentWords {
		if i == m.selected {
			b.WriteString(selectedItemStyle(word))
		} else {
			b.WriteString(" " + word + " ")
		}

		if m.page != SAVED && contains(m.savedWords, word) {
			b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render(" ✔"))
		}
		b.WriteString("\n")
	}
	
	return b.String()
}

func _random(m model) string {
	return wordsList(m)
}

func _saved(m model) string {
	return wordsList(m)
}

func _modified(m model) string {
	return wordsList(m)
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
