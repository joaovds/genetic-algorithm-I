package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type (
	geneticAlgorithm struct {
		Target              string
		stats               []generationStats
		geneQuantity        int
		numberOfChromosomes int
		MaxGenerations      int
	}

	generationStats struct {
		betterFitness int
		middleFitness int
		worseFitness  int
	}
)

func NewGeneticAlgorithm(target string, maxGenerations int) *geneticAlgorithm {
	return &geneticAlgorithm{
		Target:              target,
		geneQuantity:        len(target),
		numberOfChromosomes: len(target) * 10,
		MaxGenerations:      maxGenerations,
	}
}

func (g *geneticAlgorithm) Run() {
	var population *Population
	initialPopulation := InitialPopulation(g.numberOfChromosomes, g.geneQuantity)
	population = initialPopulation

	for generation := range g.MaxGenerations {
		population.EvaluateFitness(g.Target)

		log.Println("Generation:", generation+1)
		fmt.Println("Better:", population.Chromosomes[0].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[0].NormalizedFitness)
		fmt.Println("Middle:", population.Chromosomes[population.Size/2].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[population.Size/2].NormalizedFitness)
		fmt.Println("Worse:", population.Chromosomes[population.Size-1].GenesToString(), "=> Fitness Normalized:", population.Chromosomes[population.Size-1].NormalizedFitness)
		println()
		g.stats = append(g.stats, generationStats{
			betterFitness: population.Chromosomes[0].Fitness,
			middleFitness: population.Chromosomes[population.Size/2].Fitness,
			worseFitness:  population.Chromosomes[population.Size-1].Fitness,
		})

		if population.Chromosomes[0].GenesToString() == g.Target {
			fmt.Println("\n Target found! =>", population.Chromosomes[0].GenesToString())
			break
		}

		population = population.GenerateNextGeration()
	}
}

func (g *geneticAlgorithm) RenderChart() {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Convergence"}), charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}))
	generations := make([]int, len(g.stats))
	betterData := make([]opts.LineData, len(g.stats))
	middleData := make([]opts.LineData, len(g.stats))
	worseData := make([]opts.LineData, len(g.stats))
	for i := range len(g.stats) {
		generations[i] = i + 1
		betterData[i] = opts.LineData{Value: g.stats[i].betterFitness}
		middleData[i] = opts.LineData{Value: g.stats[i].middleFitness}
		worseData[i] = opts.LineData{Value: g.stats[i].worseFitness}
	}
	line.SetXAxis(generations).
		AddSeries("Better", betterData).
		AddSeries("Middle", middleData).
		AddSeries("Worse", worseData).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(true), ShowSymbol: opts.Bool(true), SymbolSize: 10, Symbol: "diamond"},
		), charts.WithAreaStyleOpts(opts.AreaStyle{
			Opacity: 0.1,
		}))

	f, err := os.Create("convergence_graph.html")
	if err != nil {
		panic(err)
	}
	line.Render(f)
}
