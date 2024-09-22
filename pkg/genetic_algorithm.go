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

	artificialGenes := make([]*Gene, len(g.Target))
	for i, value := range g.Target {
		artificialGenes[i] = NewGene(value)
	}
	artificialGenes2 := make([]*Gene, len(g.Target))
	for i, value := range g.Target {
		artificialGenes2[i] = NewGene(value)
	}
	mockChromosome := NewChromosome(artificialGenes)
	artificialGenes2[1].Value = rune(113)
	mockChromosome2 := NewChromosome(artificialGenes2)
	population.Chromosomes[g.numberOfChromosomes-1] = mockChromosome
	population.Chromosomes[48] = mockChromosome2

	for generation := range g.MaxGenerations {
		population.EvaluateFitness(g.Target)

		log.Println("Generation:", generation+1)
		for _, c := range population.Chromosomes {
			fmt.Println(c.GenesToString(), "=>", c.NormalizedFitness)
		}
		fmt.Println()
	}
}
