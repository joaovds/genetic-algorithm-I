package pkg

type geneticAlgorithm struct {
	Target              string
	numberOfGenes       int
	numberOfChromosomes int
	MaxGenerations      int
}

func NewGeneticAlgorithm(target string, maxGenerations int) *geneticAlgorithm {
	return &geneticAlgorithm{
		Target:              target,
		numberOfGenes:       len(target),
		numberOfChromosomes: len(target) * 10,
		MaxGenerations:      maxGenerations,
	}
}

func (g *geneticAlgorithm) GetNumberOfGenes() int {
	return g.numberOfGenes
}

func (g *geneticAlgorithm) GetNumberOfChromosomes() int {
	return g.numberOfChromosomes
}
