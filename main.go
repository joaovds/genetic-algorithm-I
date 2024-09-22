package main

import (
	"github.com/joaovds/genetic-algorithm-I/pkg"
)

func main() {
	geneticAlgorithm := pkg.NewGeneticAlgorithm("artificial", 1)
	geneticAlgorithm.Run()
}
