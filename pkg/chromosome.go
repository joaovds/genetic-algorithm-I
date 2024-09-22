package pkg

import (
	"math/rand"
	"strings"
	"time"
)

const CROSSOVER_RATE = 0.8 // 80%

type Chromosome struct {
	Genes             []*Gene
	Fitness           int
	NormalizedFitness float64
}

func NewChromosome(genes []*Gene) *Chromosome {
	chromosome := new(Chromosome)
	chromosome.Genes = genes
	return chromosome
}

func GenerateRandomChromosome(geneQuantity int) *Chromosome {
	genes := make([]*Gene, geneQuantity)
	for index := range genes {
		genes[index] = GenerateRandomGene()
	}
	return NewChromosome(genes)
}

func (c *Chromosome) GenesToString() string {
	var genesStrBuilder strings.Builder
	for _, gene := range c.Genes {
		genesStrBuilder.WriteRune(gene.Value)
	}
	return genesStrBuilder.String()
}

func (c *Chromosome) CalculateFitness(target string) int {
	if len(c.Genes) != len(target) {
		return 0
	}

	var fitness int
	for i, gene := range c.Genes {
		diff := int(target[i]) - int(gene.Value)
		fitness += diff * diff
	}
	c.Fitness = fitness
	return fitness
}

func Crossover(parent1, parent2 Chromosome) [2]*Chromosome {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	child1 := new(Chromosome)
	child2 := new(Chromosome)
	children := [2]*Chromosome{child1, child2}

	crossoverProbability := rnd.Float64()
	println("Crossover probability:", crossoverProbability)
	if crossoverProbability < CROSSOVER_RATE {
		crossoverPoint := rnd.Intn(len(parent1.Genes))

		child1.Genes = append([]*Gene{}, parent1.Genes[:crossoverPoint]...)
		child1.Genes = append(child1.Genes, parent2.Genes[crossoverPoint:]...)
		child2.Genes = append([]*Gene{}, parent2.Genes[:crossoverPoint]...)
		child2.Genes = append(child2.Genes, parent1.Genes[crossoverPoint:]...)
	} else {
		child1.Genes = parent1.Genes
		child2.Genes = parent2.Genes
	}
	return children
}
