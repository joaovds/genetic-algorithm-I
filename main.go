package main

import (
	"fmt"

	"github.com/joaovds/genetic-algorithm-I/pkg"
)

const (
	WORD_TO_FIND          = "artificial"
	NUMBER_OF_GENES       = len(WORD_TO_FIND)
	NUMBER_OF_CHROMOSOMES = len(WORD_TO_FIND) * 10
)

func main() {
	population := pkg.GeneratePopulation(NUMBER_OF_CHROMOSOMES, NUMBER_OF_GENES)

	artificialGenes := make([]pkg.Gene, len(WORD_TO_FIND))
	for i, value := range WORD_TO_FIND {
		artificialGenes[i].Value = value
	}
	artificialGenes[1].Value = rune(107)
	mockChromosome := pkg.NewChromosome(artificialGenes)
	population.Chromosomes[NUMBER_OF_CHROMOSOMES-1] = mockChromosome

	population.EvaluateFitness(WORD_TO_FIND)
	for _, chromosome := range population.Chromosomes {
		fmt.Println(chromosome.Genes, "=>", chromosome.GenesToString(), "=> fitness:", chromosome.Fitness)
	}
	fmt.Println(len(population.Chromosomes))
	theBest := population.Chromosomes[0]
	for _, chromosome := range population.Chromosomes {
		if theBest.Fitness <= chromosome.Fitness {
			theBest = chromosome
		}
	}
	fmt.Println("Melhor:", theBest.GenesToString())
	population.IndividualExists(WORD_TO_FIND)
}
