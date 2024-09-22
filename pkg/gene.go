package pkg

import (
	"math/rand"
	"time"
)

const (
	MIN_ASCII_VALUE = 97
	MAX_ASCII_VALUE = 122
)

type Gene struct {
	Value rune
}

func NewGene(value rune) *Gene {
	return &Gene{Value: value}
}

func GenerateRandomGene() *Gene {
	return NewGene(RandomGeneValue())
}

func RandomGeneValue() rune {
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
	return geneValue
}
