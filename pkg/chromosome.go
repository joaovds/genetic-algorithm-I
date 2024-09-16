package pkg

import (
	"math"
	"strings"
)

type Chromosome struct {
	Genes   []Gene
	Fitness float32
}

func NewChromosome(genes []Gene) *Chromosome {
	return &Chromosome{Genes: genes}
}

func GenerateChromosome(numberOfGenes int) *Chromosome {
	genes := make([]Gene, numberOfGenes)
	for index := range numberOfGenes {
		genes[index] = *GenerateGene()
	}
	return NewChromosome(genes)
}

func (c *Chromosome) GenesToString() string {
	var genesStrBuilder strings.Builder
	for _, gene := range c.Genes {
		genesStrBuilder.WriteRune(gene.Value)
	}
	return genesStrBuilder.String()
}

func (c *Chromosome) CalculateFitness(target string) {
	var errorScore float32
	for i, gene := range c.Genes {
		if gene.Value == rune(target[i]) {
			continue
		}
		errorScore += float32(math.Abs(float64(gene.Value) - float64(target[i])))
		errorScore += float32(i) * 0.2
	}
	normalizedScore := 1 / (1 + errorScore) * 100
	c.Fitness = normalizedScore
}
