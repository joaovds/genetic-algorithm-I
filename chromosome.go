package main

type Chromosome struct {
	Genes []Gene
}

func newChromosome(genes []Gene) *Chromosome {
	return &Chromosome{Genes: genes}
}

func GenerateChromosome(numberOfGenes int) *Chromosome {
	genes := make([]Gene, numberOfGenes)
	for index := range numberOfGenes {
		genes[index] = *GenerateGene()
	}
	return newChromosome(genes)
}
