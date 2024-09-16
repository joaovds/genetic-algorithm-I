package pkg

import (
	"math"
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

func TestGenesToString(t *testing.T) {
	expected := "test"
	genes := make([]Gene, len(expected))
	genes[0].Value = rune(116)
	genes[1].Value = rune(101)
	genes[2].Value = rune(115)
	genes[3].Value = rune(116)
	chromosome := NewChromosome(genes)
	assert.Equal(t, expected, chromosome.GenesToString())
}

func TestCalculateFitness(t *testing.T) {
	t.Run("no matching", func(t *testing.T) {
		chromosomeSize := 4
		genes := make([]Gene, chromosomeSize)
		genes[0].Value = rune(116)
		genes[1].Value = rune(101)
		genes[2].Value = rune(115)
		genes[3].Value = rune(117)
		chromosome := NewChromosome(genes)
		chromosome.CalculateFitness("test")
		assert.Equal(t,
			1/(1+float32(3)*0.2+(117-116))*100,
			chromosome.Fitness,
		)
	})

	t.Run("no matching 3", func(t *testing.T) {
		chromosomeSize := 4
		genes := make([]Gene, chromosomeSize)
		genes[0].Value = rune(100)
		genes[1].Value = rune(101)
		genes[2].Value = rune(190)
		genes[3].Value = rune(117)
		chromosome := NewChromosome(genes)
		chromosome.CalculateFitness("test")
		assert.Equal(t,
			1/(1+(float32(3)+float32(2))*0.2+(117-116)+(190-115)+float32(math.Abs(float64(100-116))))*100,
			chromosome.Fitness,
		)
	})

	t.Run("matching", func(t *testing.T) {
		chromosomeSize := 4
		genes := make([]Gene, chromosomeSize)
		genes[0].Value = rune(116)
		genes[1].Value = rune(101)
		genes[2].Value = rune(115)
		genes[3].Value = rune(116)
		chromosome := NewChromosome(genes)
		chromosome.CalculateFitness("test")
		assert.Equal(t,
			float32(100),
			chromosome.Fitness,
		)
	})
}
