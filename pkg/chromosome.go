package pkg

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	CROSSOVER_RATE = 1
	MUTATION_RATE  = 0.1
)

type Chromosome struct {
	Genes             []*Gene
	Fitness           int
	NormalizedFitness float64
}

func NewChromosome(genes []*Gene) *Chromosome {
	return &Chromosome{Genes: genes}
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

func (c *Chromosome) CalculateFitness(target string) {
	if len(c.Genes) != len(target) {
		return
	}

	var fitness int
	for i, gene := range c.Genes {
		diff := int(target[i]) - int(gene.Value)
		fitness += diff * diff
	}

	c.Fitness = fitness
}

func (c *Chromosome) Crossover(partner *Chromosome) (children [2]*Chromosome) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	fmt.Println("p1", c.GenesToString())
	fmt.Println("p2", partner.GenesToString())

	child1, child2 := new(Chromosome), new(Chromosome)
	children = [2]*Chromosome{
		child1, child2,
	}
	if rnd.Float32() < CROSSOVER_RATE {
		crossoverPoint := rnd.Intn(len(c.Genes) - 1)
		child1.Genes = append([]*Gene{}, c.Genes[:crossoverPoint]...)
		child1.Genes = append(child1.Genes, partner.Genes[crossoverPoint:]...)
		child2.Genes = append([]*Gene{}, partner.Genes[:crossoverPoint]...)
		child2.Genes = append(child2.Genes, c.Genes[crossoverPoint:]...)
	} else {
		child1.Genes = c.Genes
		child2.Genes = partner.Genes
	}

	child1.mutation()
	child2.mutation()

	return children
}

func (c *Chromosome) mutation() {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)

	for _, gene := range c.Genes {
		if rnd.Float64() < MUTATION_RATE {
			gene.Mutate()
		}
	}
}
