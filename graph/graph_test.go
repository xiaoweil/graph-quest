package graph

import "testing"

func TestGraphWith0Subgraphs(t *testing.T) {
	g := GenerateGraphWithEdges(0, map[Vertex][]Vertex{})

	requiredSize := 0
	size, _ := g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 0
	mSize, _ := g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}

func TestGraphWithFullyConnected(t *testing.T) {
	g := GenerateGraphWithEdges(4, map[Vertex][]Vertex{
		1: {2},
		2: {1, 3},
		3: {2, 4},
		4: {3},
	})

	requiredSize := 1
	size, _ := g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 1
	mSize, _ := g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}

func TestGraphWith3Subgraphs(t *testing.T) {
	g := GenerateGraphWithEdges(7, map[Vertex][]Vertex{
		1: {6, 7},
		3: {6},
		5: {6},
		6: {1, 5, 3},
		7: {1},
	})

	requiredSize := 3
	size, _ := g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 1
	mSize, _ := g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}

	g = GenerateGraphWithEdges(7, map[Vertex][]Vertex{
		1: {6},
		2: {7},
		3: {6},
		5: {6},
		6: {1, 5, 3},
		7: {2},
	})

	requiredSize = 3
	size, _ = g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize = 1
	mSize, _ = g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}

	g = GenerateGraphWithEdges(6, map[Vertex][]Vertex{
		1: {2},
		2: {1},
		3: {4},
		4: {3},
		5: {6},
		6: {5},
	})

	requiredSize = 3
	size, _ = g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize = 3
	mSize, _ = g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}

func TestGraphWithFullyDisconnected(t *testing.T) {
	g := GenerateGraph(7, 0)

	requiredSize := 7
	size, _ := g.GetConnectedSubgraphs()
	if size != requiredSize {
		t.Fatalf(`Required size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 7
	mSize, _ := g.GetMaximumConnectedSubgraphs()
	if mSize != requiredMSize {
		t.Fatalf(`Required maximum size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}
