package pkg

import (
	"fmt"
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
		errorScore += float32(math.Abs(float64(gene.Value) - float64(target[i])))
	}
	normalizedScorePercent := 1 / (1 + errorScore) * 100
	fmt.Println(c.Genes, "=>", normalizedScorePercent)
	c.Fitness = normalizedScorePercent
}
