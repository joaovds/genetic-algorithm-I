package main

import (
	"fmt"

	"github.com/joaovds/genetic-algorithm-I/pkg"
)

const (
	WORD_TO_FIND          = "artificial"
	NUMBER_OF_GENES       = len(WORD_TO_FIND)
	NUMBER_OF_CHROMOSOMES = len(WORD_TO_FIND) * 10
)

func main() {
	population := pkg.GeneratePopulation(NUMBER_OF_CHROMOSOMES, NUMBER_OF_GENES)
	for _, chromosome := range population.Chromosomes {
		fmt.Println(chromosome.Genes)
	}
	fmt.Println(len(population.Chromosomes))
}
