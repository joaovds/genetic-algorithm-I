package pkg

import (
	"sort"
	"sync"
)

type Population struct {
	Chromosomes            []*Chromosome
	Size                   int
	TotalFitness           int
	TotalNormalizedFitness float64
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
			totalFitnessChannel <- c.CalculateFitness(target)
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
	totalNormalizedFitnessChannel := make(chan float64)
	wg := &sync.WaitGroup{}
	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			if p.TotalFitness > 0 && chromosome.Fitness > 0 {
				normalizedFitness := (1 / float64(chromosome.Fitness)) / float64(p.TotalFitness) * 100
				chromosome.NormalizedFitness = normalizedFitness
				totalNormalizedFitnessChannel <- normalizedFitness
			} else {
				chromosome.NormalizedFitness = 0
			}
		}(chromosome)
	}
	go func() {
		wg.Wait()
		close(totalNormalizedFitnessChannel)
	}()
	for fitness := range totalNormalizedFitnessChannel {
		p.TotalNormalizedFitness += fitness
	}

	for _, chromosome := range p.Chromosomes {
		wg.Add(1)
		go func(c *Chromosome) {
			defer wg.Done()
			if p.TotalFitness > 0 {
				chromosome.NormalizedFitness = chromosome.NormalizedFitness / p.TotalNormalizedFitness
			} else {
				chromosome.NormalizedFitness = 0
			}
		}(chromosome)
	}
	wg.Wait()
}

func (p *Population) sortByFitness() {
	sort.Slice(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness < p.Chromosomes[j].Fitness
	})
}
