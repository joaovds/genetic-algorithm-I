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

func (g *geneticAlgorithm) GetGeneQuantity() int {
	return g.geneQuantity
}

func (g *geneticAlgorithm) GetNumberOfChromosomes() int {
	return g.numberOfChromosomes
}

func (g *geneticAlgorithm) Run() {
	var population *Population
	initialPopulation := InitialPopulation(g.numberOfChromosomes, g.geneQuantity)
	population = initialPopulation
	targetFound := false

	for generation := range g.MaxGenerations {
		log.Println("Generation:", generation+1)
		fmt.Println("Population:")
		for _, chromosome := range population.Chromosomes {
			fmt.Println(chromosome.GenesToString())
		}

		if targetFound {
			break
		}
	}
}
