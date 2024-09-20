package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePopulation(t *testing.T) {
	t.Run("should return a population with the passed chromosomes", func(t *testing.T) {
		chromosomes := []Chromosome{*GenerateRandomChromosome(2), *GenerateRandomChromosome(2)}
		population := GeneratePopulation(chromosomes)
		assert.Equal(t, chromosomes, population.Chromosomes)
		assert.Equal(t, len(chromosomes), population.Size)
	})
}

func TestInitialPopulation(t *testing.T) {
	t.Run("should return a new, random population", func(t *testing.T) {
		newPopulationSize := 100
		randomPopulation := InitialPopulation(newPopulationSize, 10)
		assert.Equal(t, newPopulationSize, len(randomPopulation.Chromosomes))
		assert.Equal(t, newPopulationSize, randomPopulation.Size)
	})
}
