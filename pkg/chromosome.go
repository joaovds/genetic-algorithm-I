package pkg

import (
	"strings"
)

type Chromosome struct {
	Genes             []*Gene
	Fitness           int
	NormalizedFitness float64
}

func NewChromosome(genes []*Gene) *Chromosome {
	return &Chromosome{Genes: genes}
}

func GenerateRandomChromosome(geneQuantity int) *Chromosome {
	genes := make([]*Gene, geneQuantity)
	for index := range genes {
		genes[index] = GenerateRandomGene()
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
	if len(c.Genes) != len(target) {
		return
	}

	var fitness int
	for i, gene := range c.Genes {
		diff := int(target[i]) - int(gene.Value)
		fitness += diff * diff
	}

	c.Fitness = fitness
}
