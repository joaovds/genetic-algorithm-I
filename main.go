package main

import (
	"github.com/joaovds/genetic-algorithm-I/pkg"
)

func main() {
	geneticAlgorithm := pkg.NewGeneticAlgorithm("artificial", 600)
	geneticAlgorithm.Run()
	geneticAlgorithm.RenderChart()
}
