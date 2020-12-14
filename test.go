package main

import (	"reflect"	"testing")

func Test_Graph(t *testing.T) {
	a := &Node{Name: "Kruthika's abode"}
	b := &Node{Name: "Brian's apartment"}
	c := &Node{Name: "Greg's casa"}
	d := &Node{Name: "Wesley's condo"}

	g := Graph{}
	g.AddEdge(a, b, 1)
	g.AddEdge(a, c, 2)
	g.AddEdge(b, d, 9)
	g.AddEdge(c, b, 10)
	g.AddEdge(c, d, 5)

	t.Run("AddNode", func(t *testing.T) {
		if len(g.Nodes) != 4 {
			t.Errorf("expected %d, got %d", 4, len(g.Nodes))
		}

		nodes := make(map[*Node]bool)
		nodes[a] = true
		nodes[b] = true
		nodes[c] = true
		nodes[d] = true
		if !reflect.DeepEqual(g.Nodes, nodes) {
			t.Errorf("expected %v, got %v", nodes, g.Nodes)
		}
	})

	t.Run("AddEdge", func(t *testing.T) {
		if len(g.Edges) != 5 {
			t.Errorf("expected %d, got %d", 5, len(g.Edges))
		}

		edges := []*Edge{
			{
				Source: a,
				Destination:  b,
				Weight:   1,
			},
			{
				Source: a,
				Destination:  c,
				Weight:   2,
			},
			{
				Source: b,
				Destination:  d,
				Weight:   9,
			},
			{
				Source: c,
				Destination:  b,
				Weight:   10,
			},
			{
				Source: c,
				Destination:  d,
				Weight:   5,
			},
		}
		if !reflect.DeepEqual(g.Edges, edges) {
			t.Errorf("expected %v, got %v", edges, g.Edges)
		}
	})

	t.Run("NewWeightTable", func(t *testing.T) {
		WeightTable := g.NewWeightTable(a)

		expectedWeightTable := make(map[*Node]int)
		expectedWeightTable[a] = 0
		expectedWeightTable[b] = Infinity
		expectedWeightTable[c] = Infinity
		expectedWeightTable[d] = Infinity

		if !reflect.DeepEqual(WeightTable, expectedWeightTable) {
			t.Errorf("expected %v, got %v", expectedWeightTable, WeightTable)
		}

	})
	t.Run("GedNodeEdges", func(t *testing.T) {
		edges := g.GetNodeEdges(a)

		expectedEdges := []*Edge{
			{
				Source: a,
				Destination:  b,
				Weight:   1,
			},
			{
				Source: a,
				Destination:  c,
				Weight:   2,
			},
		}

		if !reflect.DeepEqual(edges, expectedEdges) {
			t.Errorf("expected %v, got %v", expectedEdges, edges)
		}

	})
	t.Run("getClosestVisitedNode", func(t *testing.T) {
		WeightTable := g.NewWeightTable(a)
		node := getClosestNonVisitedNode(WeightTable, []*Node{})

		if node != a {
			t.Errorf("expected %v, got %v", a, node)
		}

	})
	t.Run("Dijkstra", func(t *testing.T) {
		WeightTable := g.Dijkstra(a)

		if WeightTable[a] != 0 {
			t.Errorf("expected %d, got %d", 0, WeightTable[a])
		}
		if WeightTable[b] != 1 {
			t.Errorf("expected %d, got %d", 1, WeightTable[b])
		}
		if WeightTable[c] != 2 {
			t.Errorf("expected %d, got %d", 2, WeightTable[c])
		}
		if WeightTable[d] != 7 {
			t.Errorf("expected %d, got %d", 7, WeightTable[d])
		}
	})
}