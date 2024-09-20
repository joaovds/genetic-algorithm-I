package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenesToString(t *testing.T) {
	expected := "test"
	genes := make([]*Gene, len(expected))
	genes[0].Value = rune(116)
	genes[1].Value = rune(101)
	genes[2].Value = rune(115)
	genes[3].Value = rune(116)
	chromosome := NewChromosome(genes)
	assert.Equal(t, expected, chromosome.GenesToString())
}
