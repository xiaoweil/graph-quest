# graph-quest

### Problem statement

We will represent a graph, G=(V, f), as a set of vertices, V and a function f(v1,v2) which returns true if and only if there is a (single) edge between v1 and v2.

The function f is symmetric: f(v1, v2) = f(v2, v1). There is no edge from v to itself: f(v, v)=0.

A graph is connected if there exists a path of edges between any two vertices in the graph. A graph with one vertex is considered connected. Define a connected subgraph, CSG=(SV,f) of G to be a graph such that

- SV is a subset of V
- CSG is connected
- There is no vertex, v, in V but not in SV such that the graph (SV+v, f) is connected. In other words CSG is as large as it can be and still be connected.

Write a Go function that takes as input V (as an array) and f and returns the number of connected subgraphs. Include one or more helper functions if that helps with clarity and organization.
Write a set of test functions that validate that the function in (1) is working correctly

### Solution Structure explained

`main.go` provides sample usage  
`/graph` contains core logic  
`/graph/*_test.go` contains test cases for core logic

run `go run .` to see the output of sample usage  
run `go test ./...` to see test results

### Logic Overview

- `Graph` object holds the following data:
  1. `vertices` - slice contains sample vertices, insert during generation
  2. `edges` - contains edge relationships, insert during generation, used for display details only
  3. `neighborFunc` - check to see if 2 given vertices are neighbors, created during generation, used for finding connected subgraphs
  4. `connectedSubgraphs` - tracks found connected subgraphs, populated on a method call
  5. `maximumConnectedSubgraphs` - tracks found connected subgraphs with maximum size, populated on a method call
- `Graph` methods:
  1. `PrintDetails()` method prints all details of the graph
  2. `GetConnectedSubgraphs()` method finds all the connected subgraphs within the graph
  3. `GetMaximumConnectedSubgraphs()` method finds all the connected subgraphs with the maximum size within the graph
- Comments on core methods to help follow along
