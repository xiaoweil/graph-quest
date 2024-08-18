package graph

import (
	"math/rand"
	"slices"
	"sort"
)

/*
Generate a random graph with a total number of Vertices and a total number of Edges

Returns a pointer of the generated graph
*/
func GenerateGraph(numOfVertices, numOfEdges int) *Graph {
	return generate(numOfVertices, numOfEdges, nil)
}

/*
Generate a random graph with a total number of Vertices and a preset of Edges

Returns a pointer of the generated graph
*/
func GenerateGraphWithEdges(numOfVertices int, edges map[Vertex][]Vertex) *Graph {
	return generate(numOfVertices, len(edges), edges)
}

/*
Generate the graph

# The edges map is a key value map, where key would be a Vertex and values would be all its connected neighbors

If edges map is nil, then randomly generates edges based on number of edges parameter
*/
func generate(numOfVertices int, numOfEdges int, edges map[Vertex][]Vertex) *Graph {
	// Checking edging cases, where only 1 or 0 vertex
	// In this case, our neighbor function should always return false
	if numOfVertices <= 1 {
		vertices := []Vertex{}
		if numOfVertices == 1 {
			vertices = append(vertices, 1)
		}

		return &Graph{
			vertices: vertices,
			edges:    map[Vertex][]Vertex{},
			neighborFunc: func(v1, v2 Vertex) bool {
				return false
			},
		}
	}

	// Generate a list of vertex based on the total number provided
	vertices := []Vertex{}
	for i := range numOfVertices {
		vertices = append(vertices, Vertex(i+1))
	}

	// Initialize the isNeighbor function, for any 2 vertices, we will loop through the edges map
	// and check to see if they are neighbors
	neighborFunc := func(v1, v2 Vertex) bool {
		return slices.Contains(edges[v1], v2)
	}

	// If no preset edges provided, we will automatically generate the edges up to the number of edges requested
	if edges == nil {
		edges = make(map[Vertex][]Vertex)

		// Ensure max edges based on number of vertex
		maxEdges := numOfVertices * (numOfVertices - 1) / 2
		if numOfEdges > maxEdges {
			numOfEdges = maxEdges
		}

		edgeCounter := 0
		for edgeCounter < numOfEdges {
			// Randomly pick 2 vertex from the range, +1 is to ensure one based value
			v1 := Vertex(rand.Intn(numOfVertices) + 1)
			v2 := Vertex(rand.Intn(numOfVertices) + 1)

			// Create bidirectional edge relationship if they are not the same vertex and are not already neighbors
			if v1 != v2 && !neighborFunc(v1, v2) {
				edges[v1] = append(edges[v1], v2)
				edges[v2] = append(edges[v2], v1)

				edgeCounter++
			}
		}
	}

	maxEdgeLength := 0
	for _, val := range edges {
		if len(val) > maxEdgeLength {
			maxEdgeLength = len(val)
		}

		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}

	return &Graph{
		vertices:     vertices,
		edges:        edges,
		maxEdges:     maxEdgeLength,
		neighborFunc: neighborFunc,
	}
}
