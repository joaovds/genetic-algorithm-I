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
		fmt.Println("Better:", population.Chromosomes[0].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[0].NormalizedFitness)
		fmt.Println("Middle:", population.Chromosomes[population.Size/2].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[population.Size/2].NormalizedFitness)
		fmt.Println("Worse:", population.Chromosomes[population.Size-1].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[population.Size-1].NormalizedFitness)
		println()

		if population.Chromosomes[0].GenesToString() == g.Target {
			fmt.Println("\n Target found! =>", population.Chromosomes[0].GenesToString())
			break
		}

		population = population.GenerateNextGeration()
	}
}
