package graph

const EDGES_THRESHOLD = 10

// Define a new Vertex type with int as value
type Vertex int

// Define a new graph structure
type Graph struct {
	vertices                  []Vertex
	edges                     map[Vertex][]Vertex
	maxEdges                  int
	neighborFunc              func(Vertex, Vertex) bool
	connectedSubgraphs        [][]Vertex
	maximumConnectedSubgraphs [][]Vertex
}
