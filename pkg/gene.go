package pkg

import (
	"math/rand"
	"time"
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
		newGeneValue := rnd.Intn(127)
		if newGeneValue < 32 {
			continue
		} else {
			geneValue = rune(newGeneValue)
			break
		}
	}
	return newGene(geneValue)
}
