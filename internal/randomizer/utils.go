package randomizer

import (
	"math/rand/v2"
)
func random(from, to int) int {
	if from >= to {
		panic("from must be less than to")
	}
	return rand.IntN(to-from) + from
}

