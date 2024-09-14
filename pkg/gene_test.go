package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGene(t *testing.T) {
	var value rune = 74
	gene := newGene(value)
	assert.NotEmpty(t, gene)
	assert.NotNil(t, gene)
	assert.Equal(t, value, gene.Value)
}

func TestGenerateGene(t *testing.T) {
	for range 500 {
		t.Run("should return a gene with a rune value between MIN_ASCII_VALUE and MAX_ASCII_VALUE", func(t *testing.T) {
			gene := GenerateGene()
			assert.NotNil(t, gene)
			assert.NotEmpty(t, gene)
			assert.GreaterOrEqual(t, int(gene.Value), MIN_ASCII_VALUE)
			assert.LessOrEqual(t, int(gene.Value), MAX_ASCII_VALUE)
		})
	}
}
