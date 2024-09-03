package tui

import (
	"fmt"
	"os"
	"wordgen/internal/randomizer"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width int
	height int
	
	saveFile string
	keys keyMap
	help help.Model
	selected int
	page int
	pages []string
	words []string
	savedWords []string
	modifiedWords []string
}

func NewModel() model {
	const saveFile = "save.json"
	words, err := loadFromFile(saveFile)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not load save file")
	}
	m := model {
		page: 0,
		pages: []string { 
			"Random",
			"Saved",
		},
		savedWords: words,
		saveFile: saveFile,
		help: help.New(),
		keys: keys,
		words: randomizer.GetWords(10),
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		// cmd tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.NextPage):
			m.selected = 0
			m.page = eumod(m.page+1, len(m.pages))
		case key.Matches(msg, m.keys.PrevPage):
			m.selected = 0
			m.page = eumod(m.page-1, len(m.pages))
		case key.Matches(msg, m.keys.Up):
			m.up()
		case key.Matches(msg, m.keys.Down):
			m.down()
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			saveToFile(m.saveFile, m.savedWords)
			return m, tea.Quit
		case key.Matches(msg, m.keys.New):
			m.words = randomizer.GetWords(10)
		case key.Matches(msg, m.keys.Save):
			m.save()
		case key.Matches(msg, m.keys.Enter):
			m.enter()
		case key.Matches(msg, m.keys.Delete):
			m.delete()
		}
	}
	
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i, page := range m.pages {
		if i == m.page {
			b.WriteString(selectedStyle(page) + " ")
		} else {
			b.WriteString(inactiveStyle(page) + " ")
		}
	}
	b.WriteString("\n\n")

	b.WriteString(page(m))

	b.WriteString("\n\n" + m.help.View(m.keys))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.NewStyle().PaddingLeft(3).PaddingTop(1).Render(b.String()),
	)
}
