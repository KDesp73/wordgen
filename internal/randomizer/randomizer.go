package randomizer

import (
	"math/rand"
	"strings"
)

const vowels = "aeiouy"
const consonants = "qwrtpsfghjdklzxcvbnmy"

func randVow() string {
	return string(vowels[rand.Intn(len(vowels))])
}

func randCon() string {
	return string(consonants[rand.Intn(len(consonants))])
}

func RandSyllable() string {
	switch rand.Intn(6) { // Increased to 6 for more variety
	case 0:
		return randCon() + randVow()
	case 1:
		return randVow() + randCon()
	case 2:
		return randVow()
	case 3:
		return randCon() + randVow() + randCon()
	case 4:
		return randCon() + randVow() + randCon()
	case 5:
		return randVow() + randCon() + randVow()
	}
	return ""
}

func RandWord(numLetters int) string {
	var s strings.Builder
	totalLetters := 0

	for totalLetters < numLetters {
		syllable := RandSyllable()
		syllableLength := len(syllable)

		if totalLetters+syllableLength > numLetters {
			syllable = syllable[:numLetters-totalLetters]
		}

		s.WriteString(syllable)
		totalLetters += len(syllable)
	}

	return s.String()
}

func FetchWords(size int) []string {
	words := []string{}

	for i := 0; i< size; i++ {
		words = append(words, RandWord(random(4, 10)))
	}
	
	return words
}

func ModifyWord(word string) string {
	index := rand.Intn(len(word))
	if strings.Index(vowels, string(word[index])) != -1 {
		return word[:index] + randVow() + word[index+1:]
	} else {
		return word[:index] + randCon() + word[index+1:]
	}
}
