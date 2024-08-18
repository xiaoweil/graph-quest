package graph

import (
	"fmt"
	"sort"
)

// Helper method to print graph basic information
func (g *Graph) PrintDetails() {
	fmt.Printf("\nGraph details:\n\n")
	if len(g.vertices) > 20 {
		fmt.Printf("Vertices: %v ... %v\n", g.vertices[0], g.vertices[len(g.vertices)-1])
	} else {
		fmt.Printf("Vertices: %v\n", g.vertices)
	}

	// Sort vertex from small to large
	keys := []int{}
	for k := range g.edges {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// Print edges
	fmt.Println("Edges:")
	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, g.edges[Vertex(k)])
	}

	// Print found connected subgraphs
	size, csgs := g.GetConnectedSubgraphs()
	fmt.Printf("Found %v connected subgraphs:\n", size)
	for _, csg := range csgs {
		fmt.Println(csg)
	}
	fmt.Printf("Total: %v\n\n", size)

	// Print found maximum subgraphs
	mSize, mcsgs := g.GetMaximumConnectedSubgraphs()
	fmt.Printf("Found %v maximum connected subgraphs:\n", mSize)
	for _, mcsg := range mcsgs {
		fmt.Println(mcsg)
	}
	fmt.Printf("Total: %v\n", mSize)
}

// Find all connected subgraphs
func (g *Graph) GetConnectedSubgraphs() (int, [][]Vertex) {
	numOfVertices := len(g.vertices)
	// Only do the finding the first time, else return found subgraphs
	if numOfVertices > 0 && len(g.connectedSubgraphs) == 0 {
		if g.maxEdges > EDGES_THRESHOLD {
			g.dfsSearch()
		} else {
			g.bfsSearch()
		}
	}

	return len(g.connectedSubgraphs), g.connectedSubgraphs
}

func (g *Graph) bfsSearch() {
	visited := make(map[Vertex]bool)

	// Here we loop through all vertices in the graph
	// We use a map to track all visited vertex
	// We pass the current vertex and tracked visited to a bfs
	// Add the return value from the bfs to connected subgraphs
	for _, v := range g.vertices {
		if !visited[v] {
			connectedSubgraph := bfs(g, v, visited)

			// Sort the values in the connected subgraphs
			sort.Slice(connectedSubgraph, func(i, j int) bool {
				return connectedSubgraph[i] < connectedSubgraph[j]
			})

			g.connectedSubgraphs = append(g.connectedSubgraphs, connectedSubgraph)
		}
	}
}

func bfs(graph *Graph, current Vertex, visited map[Vertex]bool) []Vertex {
	// Queue for BFS
	queue := []Vertex{current}
	// List to keep track of connected vertices
	connectedSubgraph := []Vertex{}

	// Mark the start vertex as visited
	visited[current] = true

	// Process the queue
	for len(queue) > 0 {
		// Dequeue the first vertex
		vertex := queue[0]
		queue = queue[1:]

		// Add it to the connected subgraph
		connectedSubgraph = append(connectedSubgraph, vertex)

		// Explore all neighbors
		for _, neighbor := range graph.vertices {
			if vertex != neighbor && !visited[neighbor] && graph.neighborFunc(vertex, neighbor) {
				// Mark neighbor as visited and enqueue it
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return connectedSubgraph
}

func (g *Graph) dfsSearch() {
	visited := make(map[Vertex]bool)

	// Here we loop through all vertices in the graph
	// We use a map to track all visited vertex
	// We use a slice to track all connected vertex
	// We pass the current vertex and tracked visited and connected slice to a dfs
	for _, v := range g.vertices {
		if !visited[v] {
			var connectedSubgraph []Vertex
			dfs(g, v, visited, &connectedSubgraph)

			// Sort the values in the connected subgraphs
			sort.Slice(connectedSubgraph, func(i, j int) bool {
				return connectedSubgraph[i] < connectedSubgraph[j]
			})

			g.connectedSubgraphs = append(g.connectedSubgraphs, connectedSubgraph)
		}
	}
}

func dfs(graph *Graph, current Vertex, visited map[Vertex]bool, connectedSubgraph *[]Vertex) {
	// For each vertex, we mark it as visited and track in a map
	visited[current] = true
	// We track the current vertex in a connected subgraph slice
	*connectedSubgraph = append(*connectedSubgraph, current)

	// Here we loop through all the vertices, and see if its one of the current's neighbors
	// and for each found neighbor we will put it through the dfs to track all its neighbors.
	// We track all connected neighbors in the connectedSubgraph.
	// Here we assumed that graph.edges cannot be accessed, only the neighborFunc could be used
	// to align with requirements.
	// If graph.edges could be used, then we just need to loop through graph.edges[current]
	// Such:
	// for _, neighbor := range graph.edges[current] {
	//     if !visited[neighbor] {
	//         ...
	//     }
	for _, neighbor := range graph.vertices {
		if current != neighbor && !visited[neighbor] && graph.neighborFunc(current, neighbor) {
			dfs(graph, neighbor, visited, connectedSubgraph)
		}
	}
}

func (g *Graph) GetMaximumConnectedSubgraphs() (int, [][]Vertex) {
	// Find all the connected subgraphs first before finding the maximums
	if len(g.vertices) > 0 && len(g.connectedSubgraphs) == 0 {
		g.GetConnectedSubgraphs()
	}

	// Only do the finding the first time, else return found maximum subgraphs
	if len(g.vertices) > 0 && len(g.maximumConnectedSubgraphs) == 0 {
		maxLength := 0

		// Loop through all the found connected subgraphs and track the ones with largest size
		for _, subgraphs := range g.connectedSubgraphs {
			subgraphsLength := len(subgraphs)
			if subgraphsLength > maxLength {
				maxLength = subgraphsLength
				// Keep replacing our maximum slices when new max length found
				g.maximumConnectedSubgraphs = [][]Vertex{subgraphs}
			} else if len(subgraphs) == maxLength {
				g.maximumConnectedSubgraphs = append(g.maximumConnectedSubgraphs, subgraphs)
			}
		}
	}

	return len(g.maximumConnectedSubgraphs), g.maximumConnectedSubgraphs
}
