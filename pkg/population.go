package pkg

type Population struct {
	Chromosomes []Chromosome
	Size        int
}

func GeneratePopulation(chromosomes []Chromosome) *Population {
	return &Population{
		Chromosomes: chromosomes,
		Size:        len(chromosomes),
	}
}

func InitialPopulation(size, geneQuantity int) *Population {
	chromosomes := make([]Chromosome, size)

	for index := range size {
		chromosomes[index] = *GenerateRandomChromosome(geneQuantity)
	}

	return GeneratePopulation(chromosomes)
}
