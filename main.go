package main

import (
	"github.com/xiaoweil/graph-quest/graph"
)

func main() {
	// Create a graph with 7 vertices and 4 random edges
	g := graph.GenerateGraph(7, 4)
	g.PrintDetails()

	// Create a graph with 6 vertices and preset edges
	g2 := graph.GenerateGraphWithEdges(6, map[graph.Vertex][]graph.Vertex{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {6},
		6: {5},
	})
	g2.PrintDetails()
}
