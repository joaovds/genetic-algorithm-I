package main

import (
	"log"

	"github.com/joaovds/genetic-algorithm-I/pkg"
)

func main() {
	geneticAlgorithm := pkg.NewGeneticAlgorithm("artificial", 1000)
	go geneticAlgorithm.Run()

	for generation := range geneticAlgorithm.GenerationCount {
		log.Println("Generation", generation)
	}
}
