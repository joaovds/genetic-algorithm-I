package main

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
		t.Run("should return a gene with a rune value between 32 and 127", func(t *testing.T) {
			gene := GenerateGene()
			assert.NotNil(t, gene)
			assert.NotEmpty(t, gene)
			assert.GreaterOrEqual(t, int(gene.Value), 32)
			assert.LessOrEqual(t, int(gene.Value), 127)
		})
	}
}
