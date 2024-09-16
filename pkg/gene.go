package pkg

import (
	"math/rand"
	"time"
)

const (
	MIN_ASCII_VALUE = 97
	MAX_ASCII_VALUE = 122
	MUTATION_RATE   = 0.4
)

type Gene struct {
	Value rune
}

func newGene(value rune) *Gene {
	return &Gene{Value: value}
}

func GenerateGene() *Gene {
	geneValue := newGeneValue()
	return newGene(geneValue)
}

func newGeneValue() rune {
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

func (g *Gene) Mutate() {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	mutationProb := rnd.Float64()
	if mutationProb < MUTATION_RATE {
		g.Value = newGeneValue()
	}
}
