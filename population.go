package main

type Population struct {
	Chromosomes []*Chromosome
}

func newPopulation(chromosomes []*Chromosome) *Population {
	return &Population{Chromosomes: chromosomes}
}

func GeneratePopulation(populationSize, numberOfGenes int) *Population {
	chromosomes := make([]*Chromosome, populationSize)

	for index := range populationSize {
		chromosomes[index] = GenerateChromosome(numberOfGenes)
	}

	return newPopulation(chromosomes)
}
