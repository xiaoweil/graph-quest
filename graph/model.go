package graph

// Define a new Vertex type with int as value
type Vertex int

// Define a new graph structure
type Graph struct {
	vertices                  []Vertex
	edges                     map[Vertex][]Vertex
	neighborFunc              func(Vertex, Vertex) bool
	connectedSubgraphs        [][]Vertex
	maximumConnectedSubgraphs [][]Vertex
}
