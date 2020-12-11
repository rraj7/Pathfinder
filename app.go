package main

import (	"fmt"	;"sort";	"strconv");

type Graph struct {	Edges []*Edge	
	Nodes []*Node}

type Edge struct {	Source *Node
		Destination  *Node
		Weight   int
	}

type Node struct {	Name string }

const Infinity = int(^uint(0) >> 1)

func (g *Graph) AddEdge(Source, Destination *Node, Weight int) {
	edge := &Edge{
		Source: Source,
		Destination:  Destination,
		Weight:   Weight,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(Source)
	g.AddNode(Destination)
}


func (g *Graph) AddNode(node *Node) {
	var ifNodeExists bool
	for _, n := range g.Nodes {
		if n == node {
			ifNodeExists = true
		}
	}

	if !ifNodeExists {
		g.Nodes = append(g.Nodes, node)
	}
}


func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Source.Name + " -> " + edge.Destination.Name + " = " + strconv.Itoa(edge.Weight)
		s += "\n"
	}
	s += "\n"

	s += "Nodes: "
	for i, node := range g.Nodes {
		if i == len(g.Nodes)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
	}
	s += "\n"

	return s
}


func (g *Graph) Dijkstra(startNode *Node) (shortestPathTable string) {


	WeightTable := g.NewWeightTable(startNode)

	
	var visited []*Node

	
	for len(visited) != len(g.Nodes) {

	
		node := getClosestNonVisitedNode(WeightTable, visited)

	
		visited = append(visited, node)

	
		nodeEdges := g.GetNodeEdges(node)

		for _, edge := range nodeEdges {
		distanceToNeighbor := WeightTable[node] + edge.Weight


			if distanceToNeighbor < WeightTable[edge.Destination] {
				WeightTable[edge.Destination] = distanceToNeighbor
			}
		}
	}

	for node, Weight := range WeightTable {
		shortestPathTable += fmt.Sprintf("Shortest Distances from points %s to %s = %d\n", startNode.Name, node.Name, Weight)
	}

	return shortestPathTable
}
func (g *Graph) NewWeightTable(startNode *Node) map[*Node]int {
	WeightTable := make(map[*Node]int)
	WeightTable[startNode] = 0

	for _, node := range g.Nodes {
		if node != startNode {
			WeightTable[node] = Infinity
		}
	}

	return WeightTable
}

func (g *Graph) GetNodeEdges(node *Node) (edges []*Edge) {
	for _, edge := range g.Edges {
		if edge.Source == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

func getClosestNonVisitedNode(WeightTable map[*Node]int, visited []*Node) *Node {
	type WeightTableToSort struct {
		Node *Node
		Weight int
	}
	var sorted []WeightTableToSort


	for node, Weight := range WeightTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}

		if !isVisited {
			sorted = append(sorted, WeightTableToSort{node, Weight})
		}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Weight < sorted[j].Weight
	})

	return sorted[0].Node
}

func main() {
	a := &Node{Name: "Kruthika's abode"}
	b := &Node{Name: "Brian's apartment"}
	c := &Node{Name: "Greg's casac"}
	d := &Node{Name: "Wesley's condo"}
	e := &Node{Name: "Matt's pad"}
	f := &Node{Name: "Mark's crib"}
	g := &Node{Name: "Bryce's den"}
	h := &Node{Name: "Mike's digs"}
	i := &Node{Name: "Cam's dwelling"} 
	j := &Node{Name: "Nathan's flat"}
	k := &Node{Name: "Kirk's farm"}

	graph := Graph{}
	graph.AddEdge(a, c, 9)
	graph.AddEdge(a, e, 4)
	graph.AddEdge(c, b, 18)
	graph.AddEdge(c, d, 8)
	graph.AddEdge(b, d, 4)
	graph.AddEdge(d, e, 2)
	graph.AddEdge(d, g, 30)
	graph.AddEdge(d, f, 10)
	graph.AddEdge(f, g, 1)
	graph.AddEdge(g, f, 2)
	graph.AddEdge(h, i, 4)
	graph.AddEdge(j, k, 3)


	fmt.Println(graph.Dijkstra(a))
	
	/*fmt.Println(graph.Dijkstra(b))
	 fmt.Println(graph.Dijkstra(c))
	fmt.Println(graph.Dijkstra(d))
	fmt.Println(graph.Dijkstra(e))
	fmt.Println(graph.Dijkstra(f))
	fmt.Println(graph.Dijkstra(g))
	fmt.Println(graph.Dijkstra(h))
	fmt.Println(graph.Dijkstra(i))
	fmt.Println(graph.Dijkstra(j))
	fmt.Println(graph.Dijkstra(k)) */

}