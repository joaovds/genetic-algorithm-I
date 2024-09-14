package pkg

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewChromosome(t *testing.T) {
	n_genes := 10
	genes := make([]Gene, n_genes)
	for index := range n_genes {
		genes[index] = *GenerateGene()
	}
	chromosome := NewChromosome(genes)
	assert.NotEmpty(t, chromosome)
	assert.NotNil(t, chromosome)
	assert.Equal(t, n_genes, len(chromosome.Genes))
}

func TestGenerateChromosome(t *testing.T) {
	randSource := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(randSource)
	for range 500 {
		n_genes := rnd.Intn(100)
		t.Run("should return a chromosome with x number of genes", func(t *testing.T) {
			chromosome := GenerateChromosome(n_genes)
			assert.NotNil(t, chromosome)
			assert.NotEmpty(t, chromosome)
		})
	}
}
