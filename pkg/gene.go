package pkg

import (
	"math/rand"
	"time"
)

const (
	MIN_ASCII_VALUE = 65
	MAX_ASCII_VALUE = 127
)

type Gene struct {
	Value rune
}

func newGene(value rune) *Gene {
	return &Gene{Value: value}
}

func GenerateGene() *Gene {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	var geneValue rune
	for {
		newGeneValue := rnd.Intn(MAX_ASCII_VALUE)
		if newGeneValue < MIN_ASCII_VALUE {
			continue
		} else {
			geneValue = rune(newGeneValue)
			break
		}
	}
	return newGene(geneValue)
}
