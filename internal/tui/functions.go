package tui

import "wordgen/internal/randomizer"

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

func getWord(m model) string {
	var word string
	switch m.page {
	case RANDOM:
		word = m.words[m.selected]
	case SAVED:
		word = m.savedWords[m.selected]
	case MODIFIED:
		word = m.modifiedWords[m.selected]
	}

	return word
}

func (m *model) enter() {
	if len(m.pages) == 2{
		m.pages = append(m.pages, "Modified")
	}
	
	word := getWord(*m)

	m.modifiedWords = []string{}
	m.modifiedWords = append(m.modifiedWords, word)
	m.modifiedWords = append(m.modifiedWords, randomizer.GetModifiedWords(word, 10)...)

	m.page = MODIFIED
}

func (m *model) save() {
	word := getWord(*m)

	if contains(m.savedWords, word) {
		return
	}

	m.savedWords = append(m.savedWords, word)

	if m.page == SAVED {
		m.selected = len(m.savedWords) - 1
	}
}

func (m *model) delete() {
	if m.page == SAVED {
		m.savedWords = append(m.savedWords[:m.selected], m.savedWords[m.selected+1:]...)
	}
}
