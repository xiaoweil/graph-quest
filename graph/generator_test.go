package graph

import "testing"

func TestGraphWithInvalidInputs(t *testing.T) {
	g := GenerateGraph(-1, -1)

	requiredSize := 0
	size := len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 0
	mSize := len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}

	g = GenerateGraph(1, 100)

	requiredSize = 1
	size = len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize = 0
	mSize = len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}

func TestGraphWith1Vertex(t *testing.T) {
	g := GenerateGraph(1, 0)

	requiredSize := 1
	size := len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 0
	mSize := len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}

	g = GenerateGraph(1, 1)

	requiredSize = 1
	size = len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize = 0
	mSize = len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}

func TestGraphWith2Vertices(t *testing.T) {
	g := GenerateGraph(2, 0)

	requiredSize := 2
	size := len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize := 0
	mSize := len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}

	g = GenerateGraph(2, 1)

	requiredSize = 2
	size = len(g.vertices)
	if size != requiredSize {
		t.Fatalf(`Required vertices size %v does not match with findings length %v`, requiredSize, size)
	}

	requiredMSize = 2
	mSize = len(g.edges)
	if mSize != requiredMSize {
		t.Fatalf(`Required edge size %v does not match with findings length %v`, requiredMSize, mSize)
	}
}
