package main

import (
	"github.com/jkao1/spring-curves/display"
	"github.com/jkao1/spring-curves/parser"
)

func main() {
	screen := display.NewScreen()
	transform := make([][]float64, 0)
	edges := make([][]float64, 4)

	parser.ParseFile("script", transform, edges, screen)
}
