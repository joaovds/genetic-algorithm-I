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
	p.SortByFitness()
}

func (p *Population) SortByFitness() {
	sort.Slice(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness > p.Chromosomes[j].Fitness
	})
}

func (p *Population) IsSortedByFitness() bool {
	return sort.SliceIsSorted(p.Chromosomes, func(i, j int) bool {
		return p.Chromosomes[i].Fitness > p.Chromosomes[j].Fitness
	})
}

func (p *Population) IndividualExists(target string) bool {
	if !p.IsSortedByFitness() {
		p.SortByFitness()
	}
	return p.Chromosomes[0].GenesToString() == target
}

func (p *Population) GenerateNextGeneration() *Population {
	log.Println("gerando proxima geração")
	parentsCrossed := make(map[string]bool)
	nextGenerationChromosomes := make([]*Chromosome, len(p.Chromosomes))

	numberOfNewChromosomes := 0
	for {
		parentsIndex := p.selectParentsIndex(&parentsCrossed)
		children := p.Chromosomes[parentsIndex[0]].Crossover(*p.Chromosomes[parentsIndex[1]])

		for _, child := range children {
			if numberOfNewChromosomes < len(p.Chromosomes) {
				nextGenerationChromosomes[numberOfNewChromosomes] = child
				numberOfNewChromosomes += 1
			}
		}

		if numberOfNewChromosomes >= len(p.Chromosomes) {
			break
		}
	}

	nextGenerationChromosomes[0] = p.Chromosomes[0]
	nextGenerationChromosomes[1] = p.Chromosomes[1]

	return newPopulation(nextGenerationChromosomes)
}

func (p *Population) selectParentsIndex(parentsCrossed *map[string]bool) [2]int {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	for {
		firstParentIndex := rnd.Intn(len(p.Chromosomes))
		secondParentIndex := rnd.Intn(len(p.Chromosomes))

		if firstParentIndex > secondParentIndex {
			firstParentIndex, secondParentIndex = secondParentIndex, firstParentIndex
		}

		key := fmt.Sprintf("%d-%d", firstParentIndex, secondParentIndex)
		if _, ok := (*parentsCrossed)[key]; ok {
			log.Println("já cruzado", firstParentIndex, "&", secondParentIndex, parentsCrossed)
			continue
		} else if firstParentIndex == secondParentIndex {
			continue
		}

		(*parentsCrossed)[key] = true
		return [2]int{firstParentIndex, secondParentIndex}
	}
}
