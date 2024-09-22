package pkg

import (
	"fmt"
	"log"
)

type geneticAlgorithm struct {
	Target              string
	geneQuantity        int
	numberOfChromosomes int
	MaxGenerations      int
}

func NewGeneticAlgorithm(target string, maxGenerations int) *geneticAlgorithm {
	return &geneticAlgorithm{
		Target:              target,
		geneQuantity:        len(target),
		numberOfChromosomes: len(target) * 10,
		MaxGenerations:      maxGenerations,
	}
}

func (g *geneticAlgorithm) Run() {
	var population *Population
	initialPopulation := InitialPopulation(g.numberOfChromosomes, g.geneQuantity)
	population = initialPopulation

	for generation := range g.MaxGenerations {
		population.EvaluateFitness(g.Target)

		log.Println("Generation:", generation+1)
		for _, c := range population.Chromosomes {
			fmt.Println(c.GenesToString(), "=>", c.NormalizedFitness)
		}
		fmt.Println()

		if population.Chromosomes[0].GenesToString() == g.Target {
			fmt.Println("\n Target found! =>", population.Chromosomes[0].GenesToString())
			break
		}

		population = population.GenerateNextGeration()
	}
}
