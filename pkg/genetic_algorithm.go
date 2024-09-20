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

	for generation := range g.MaxGenerations {
		population.EvaluateFitness(g.Target)

		log.Println("Generation:", generation+1)
		fmt.Println("Better:", population.Chromosomes[0].GenesToString(), "=> Fitness:", population.Chromosomes[0].Fitness, "=> Fitness normalized:", population.Chromosomes[0].NormalizedFitness)
		fmt.Println("Middle:", population.Chromosomes[len(population.Chromosomes)/2].GenesToString(), "=> Fitness:", population.Chromosomes[len(population.Chromosomes)/2].Fitness, "=> Fitness normalized:", population.Chromosomes[len(population.Chromosomes)/2].NormalizedFitness)
		fmt.Println("Worse:", population.Chromosomes[len(population.Chromosomes)-1].GenesToString(), "=> Fitness:", population.Chromosomes[len(population.Chromosomes)-1].Fitness, "=> Fitness normalized:", population.Chromosomes[len(population.Chromosomes)-1].NormalizedFitness)

		if population.Chromosomes[0].GenesToString() == g.Target {
			break
		}

		population = population.GenerateNextGeration()
		fmt.Println("Nova População:")
		for _, c := range population.Chromosomes {
			fmt.Println(c.GenesToString())
		}
		fmt.Println()
	}
}
