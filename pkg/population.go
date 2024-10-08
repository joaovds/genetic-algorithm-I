package pkg

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

const (
	ELITISM_NUMBER = 4
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

func (p *Population) GenerateNextGeration() *Population {
	nextGenerationChromosomes := make([]*Chromosome, p.Size)
	numberOfNewChromosomes := 0
	// elitism
	for i := range ELITISM_NUMBER {
		if i == p.Size {
			break
		}
		nextGenerationChromosomes[i] = p.Chromosomes[i]
		numberOfNewChromosomes++
	}

	for {
		parents := p.parentSelection()
		children := Crossover(*parents[0], *parents[1])

		for _, child := range children {
			if numberOfNewChromosomes < len(nextGenerationChromosomes) {
				nextGenerationChromosomes[numberOfNewChromosomes] = child
				numberOfNewChromosomes++
			}
		}

		if numberOfNewChromosomes >= len(nextGenerationChromosomes) {
			break
		}
	}

	return GeneratePopulation(nextGenerationChromosomes)
}

func (p *Population) parentSelection() [2]*Chromosome {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	parent1Random := rnd.Float64()
	parent2Random := rnd.Float64()
	parents := [2]*Chromosome{}

	fitnessAccumulated := 0.0
	for i := range p.Size {
		fitnessAccumulated += p.Chromosomes[i].NormalizedFitness
		if fitnessAccumulated >= parent1Random {
			parents[0] = p.Chromosomes[i]
			break
		}
	}

	fitnessAccumulated = 0.0
	for i := range p.Size {
		fitnessAccumulated += p.Chromosomes[i].NormalizedFitness
		if fitnessAccumulated >= parent2Random {
			parents[1] = p.Chromosomes[i]
			break
		}
	}

	return parents
}
