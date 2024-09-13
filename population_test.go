package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPopulation(t *testing.T) {
	n_chromosomes := 10
	chromosomes := make([]*Chromosome, n_chromosomes)
	for index := range n_chromosomes {
		chromosomes[index] = GenerateChromosome(1)
	}
	population := newPopulation(chromosomes)
	assert.NotEmpty(t, population)
	assert.NotNil(t, population)
	assert.Equal(t, n_chromosomes, len(population.Chromosomes))
}

func TestGeneratePopulation(t *testing.T) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	for range 100 {
		population_size := rnd.Intn(200)
		n_genes := rnd.Intn(10)
		t.Run("should return a population with x number of chromosomes", func(t *testing.T) {
			population := GeneratePopulation(population_size, n_genes)
			assert.NotNil(t, population)
			assert.NotEmpty(t, population)
			assert.Equal(t, population_size, len(population.Chromosomes))
		})
	}
}
