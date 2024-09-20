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

	artificialGenes := make([]*Gene, len(g.Target))
	for i, value := range g.Target {
		artificialGenes[i] = NewGene(value)
	}
	artificialGenes2 := make([]*Gene, len(g.Target))
	for i, value := range g.Target {
		artificialGenes2[i] = NewGene(value)
	}
	mockChromosome := NewChromosome(artificialGenes)
	artificialGenes2[1].Value = rune(107)
	mockChromosome2 := NewChromosome(artificialGenes2)
	population.Chromosomes[g.numberOfChromosomes-1] = mockChromosome
	population.Chromosomes[48] = mockChromosome2

	for generation := range g.MaxGenerations {
		population.EvaluateFitness(g.Target)

		log.Println("Generation:", generation+1)
		fmt.Println("Better:", population.Chromosomes[0].GenesToString(), "=> Fitness:", population.Chromosomes[0].Fitness, "=> Fitness normalized:", population.Chromosomes[0].NormalizedFitness)
		fmt.Println("Middle:", population.Chromosomes[len(population.Chromosomes)/2].GenesToString(), "=> Fitness:", population.Chromosomes[len(population.Chromosomes)/2].Fitness, "=> Fitness normalized:", population.Chromosomes[len(population.Chromosomes)/2].NormalizedFitness)
		fmt.Println("Worse:", population.Chromosomes[len(population.Chromosomes)-1].GenesToString(), "=> Fitness:", population.Chromosomes[len(population.Chromosomes)-1].Fitness, "=> Fitness normalized:", population.Chromosomes[len(population.Chromosomes)-1].NormalizedFitness)

		for _, a := range population.Chromosomes {
			fmt.Println(a.GenesToString())
		}

		if targetFound {
			break
		}
	}
}
