package pkg

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"
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

func (p *Population) GenerateNextGeration() *Population {
	nextGenerationChromosomes := make([]*Chromosome, p.Size)
	parentsCrossed := make(map[string]bool)

	parentsSelected := p.parentSelection(&parentsCrossed)
	fmt.Println(parentsSelected[0].GenesToString())
	fmt.Println(parentsSelected[1].GenesToString())

	return GeneratePopulation(nextGenerationChromosomes)
}

func (p *Population) parentSelection(parentsCrossed *map[string]bool) [2]*Chromosome {
	parent1Index := p.tournament()
	for {
		parent2Index := p.tournament()
		key := fmt.Sprintf("%d-%d", parent1Index, parent2Index)
		if _, ok := (*parentsCrossed)[key]; ok {
			log.Println("já cruzado", parent1Index, "&", parent2Index, parentsCrossed)
			continue
		} else if parent1Index == parent2Index {
			continue
		}

		(*parentsCrossed)[key] = true
		return [2]*Chromosome{p.Chromosomes[parent1Index], p.Chromosomes[parent2Index]}
	}
}

func (p *Population) tournament() (selectedParentIndex int) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	for {
		competitor1Index := rnd.Intn(len(p.Chromosomes))
		competitor2Index := rnd.Intn(len(p.Chromosomes))

		if competitor1Index == competitor2Index {
			continue
		}

		if p.Chromosomes[competitor1Index].NormalizedFitness < p.Chromosomes[competitor2Index].NormalizedFitness {
			selectedParentIndex = competitor1Index
		} else {
			selectedParentIndex = competitor2Index
		}
		return
	}
}
