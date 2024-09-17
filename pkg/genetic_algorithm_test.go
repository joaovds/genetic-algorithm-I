package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGeneticAlgorithm(t *testing.T) {
	geneticAlgorithm := NewGeneticAlgorithm("test", 1000)
	assert.Equal(t, 4, geneticAlgorithm.GetNumberOfGenes())
	assert.Equal(t, 40, geneticAlgorithm.GetNumberOfChromosomes())
	assert.Equal(t, 1000, geneticAlgorithm.MaxGenerations)
	assert.Equal(t, "test", geneticAlgorithm.Target)
}
