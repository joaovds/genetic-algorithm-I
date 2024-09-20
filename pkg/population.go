package pkg

import (
	"sort"
	"sync"
)

type Population struct {
	Chromosomes  []*Chromosome
	Size         int
	TotalFitness int
}

func GeneratePopulation(chromosomes []*Chromosome) *Population {
	return &Population{
		Chromosomes:  chromosomes,
		Size:         len(chromosomes),
		TotalFitness: 0,
	}
}

func InitialPopulation(size, geneQuantity int) *Population {
	chromosomes := make([]*Chromosome, size)

	for index := range size {
		chromosomes[index] = GenerateRandomChromosome(geneQuantity)
	}

	return GeneratePopulation(chromosomes)
}

func (p *Population) EvaluateFitness(target string) {
	totalFitnessChannel := make(chan int)

	wg := &sync.WaitGroup{}
	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			c.CalculateFitness(target)
			totalFitnessChannel <- c.Fitness
		}(chromosome)
	}
	go func() {
		wg.Wait()
		close(totalFitnessChannel)
	}()

	for fitness := range totalFitnessChannel {
		p.TotalFitness += fitness
	}
	p.sortByFitness()
	p.normalizeFitness()
}

func (p *Population) normalizeFitness() {
	for _, chromosome := range p.Chromosomes {
		if p.TotalFitness > 0 && chromosome.Fitness > 0 {
			chromosome.NormalizedFitness = (1 / float64(chromosome.Fitness)) / float64(p.TotalFitness) * 100
		} else {
			chromosome.NormalizedFitness = 0
		}
	}
}

func (p *Population) sortByFitness() {
	sort.Slice(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness < p.Chromosomes[j].Fitness
	})
}

func (p *Population) isSortedByFitness() bool {
	return sort.SliceIsSorted(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness < p.Chromosomes[j].Fitness
	})
}
