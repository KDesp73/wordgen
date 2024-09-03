package tui

import "randomname/internal/randomizer"

func (m *model) up() {
	switch m.page {
	case SAVED:
		m.selected = eumod(m.selected-1, len(m.savedWords))
	case RANDOM:
		m.selected = eumod(m.selected-1, len(m.words))
	case MODIFIED:
		m.selected = eumod(m.selected-1, len(m.modifiedWords))
	} 
}

func (m *model) down() {
	switch m.page {
	case SAVED:
		m.selected = eumod(m.selected+1, len(m.savedWords))
	case RANDOM:
		m.selected = eumod(m.selected+1, len(m.words))
	case MODIFIED:
		m.selected = eumod(m.selected+1, len(m.modifiedWords))
	} 
}

func (m *model) enter() {
	if len(m.pages) == 2{
		m.pages = append(m.pages, "Modified")
	}
	
	var word string
	switch m.page {
	case RANDOM:
		word = m.words[m.selected]
	case SAVED:
		word = m.savedWords[m.selected]
	case MODIFIED:
		word = m.modifiedWords[m.selected]
	}

	m.modifiedWords = []string{}
	m.modifiedWords = append(m.modifiedWords, word)
	m.modifiedWords = append(m.modifiedWords, randomizer.GetModifiedWords(word, 10)...)

	m.page = MODIFIED
}

func (m *model) save() {
	word := m.table.SelectedRow()[0]

	m.savedWords = append(m.savedWords, word)
}

func (m *model) delete() {
	if m.page == SAVED {
		m.savedWords = append(m.savedWords[:m.selected], m.savedWords[m.selected+1:]...)
	}
}
