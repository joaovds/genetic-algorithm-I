package pkg

type geneticAlgorithm struct {
	GenerationCount     chan int
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
		GenerationCount:     make(chan int),
	}
}

func (g *geneticAlgorithm) GetNumberOfGenes() int {
	return g.numberOfGenes
}

func (g *geneticAlgorithm) GetNumberOfChromosomes() int {
	return g.numberOfChromosomes
}

func (g *geneticAlgorithm) Run() {
	for generation := range g.MaxGenerations {
		g.GenerationCount <- generation + 1
	}
	close(g.GenerationCount)
}
