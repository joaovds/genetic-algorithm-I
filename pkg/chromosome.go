package pkg

import "strings"

type Chromosome struct {
	Genes []*Gene
}

func NewChromosome(genes []*Gene) *Chromosome {
	chromosome := new(Chromosome)
	chromosome.Genes = genes
	return chromosome
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
