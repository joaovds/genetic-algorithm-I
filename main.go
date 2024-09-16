package main

import (
	"log"

	"github.com/joaovds/genetic-algorithm-I/pkg"
)

const (
	WORD_TO_FIND          = "artificial"
	NUMBER_OF_GENES       = len(WORD_TO_FIND)
	NUMBER_OF_CHROMOSOMES = len(WORD_TO_FIND) * 10
	MAX_GENERATIONS       = 1000
)

func main() {
	population := pkg.GeneratePopulation(NUMBER_OF_CHROMOSOMES, NUMBER_OF_GENES)

	artificialGenes := make([]pkg.Gene, len(WORD_TO_FIND))
	for i, value := range WORD_TO_FIND {
		artificialGenes[i].Value = value
	}
	// artificialGenes[1].Value = rune(107)
	// mockChromosome := pkg.NewChromosome(artificialGenes)
	// population.Chromosomes[NUMBER_OF_CHROMOSOMES-1] = mockChromosome

	geration := 0
	for {
		population.EvaluateFitness(WORD_TO_FIND)
		theBest := population.Chromosomes[0]
		for _, chromosome := range population.Chromosomes {
			if theBest.Fitness <= chromosome.Fitness {
				theBest = chromosome
			}
		}
		log.Println("Melhor:", theBest.GenesToString())

		if population.IndividualExists(WORD_TO_FIND) {
			log.Println("Achou")
		} else {
			population = population.GenerateNextGeneration()
			geration++
		}

		if geration >= MAX_GENERATIONS {
			log.Println("Stop! Maximum number of generations exceeded!")
			break
		}
	}
}
