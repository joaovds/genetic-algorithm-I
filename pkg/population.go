package pkg

import "sync"

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

func (p *Population) EvaluateFitness(target string) {
	wg := sync.WaitGroup{}
	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func() {
			chromosome.CalculateFitness(target)
			wg.Done()
		}()
	}
	wg.Wait()
}
