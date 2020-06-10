package dijkstra

import (
	"errors"
	"fmt"
	"math"
)

// Node reprensets a Node in a Graph
type Node string

// Edge represents an edge between Source and Dest with a given Cost in a Graph
type Edge struct {
	Source Node
	Dest   Node
	Cost   float64
}

// Graph directed graph which consists of a set of edges
type Graph struct {
	Edges []Edge
	Nodes []Node
}

// Route represents a path from two nodes
type Route struct {
	Nodes []Node
	Cost  float64
}

// NewGraph returns a Graph object from a list of edges
func NewGraph(edges []Edge) *Graph {
	nodesMap := map[Node]bool{}
	for _, edge := range edges {
		nodesMap[edge.Source] = true
		nodesMap[edge.Dest] = true
	}
	nodes := []Node{}
	for node := range nodesMap {
		nodes = append(nodes, node)
	}
	return &Graph{Edges: edges, Nodes: nodes}
}

// Contains checks if the Node a is in the Graph g
func (g *Graph) Contains(a Node) bool {
	found := false
	for _, node := range g.Nodes {
		if a == node {
			found = true
			break
		}
	}
	return found
}

// EdgesFrom returns the edges that start in Node a
func (g *Graph) EdgesFrom(a Node) []Edge {
	edges := []Edge{}
	for _, e := range g.Edges {
		if e.Source == a {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *Graph) getNextNode(nodeRoutes map[Node]Route, visitedNodes map[Node]bool) Node {
	min := math.Inf(1)
	node := Node("")
	for n, r := range nodeRoutes {
		if visitedNodes[n] {
			continue
		}
		if r.Cost < min {
			min = r.Cost
			node = n
		}
	}
	return node
}

// ErrUnreachable is returned when a node can not be reached from another node
var ErrUnreachable = errors.New("unreachable node")

// Distance returns the distance and the route between Node a and b in Graph g
func (g *Graph) Distance(a, b Node) (*Route, error) {
	if !(g.Contains(a) && g.Contains(b)) {
		return nil, fmt.Errorf("Nodes %s and %s must belong to graph", a, b)
	}
	nodeRoutes := map[Node]Route{
		a: Route{
			Nodes: []Node{},
			Cost:  0,
		},
	}
	visitedNodes := map[Node]bool{}
	currentNode := a
	for currentNode != Node("") {
		currentRoute := nodeRoutes[currentNode]
		for _, edge := range g.EdgesFrom(currentNode) {
			newRoute := Route{
				Nodes: append(currentRoute.Nodes, currentNode),
				Cost:  currentRoute.Cost + edge.Cost,
			}
			if oldRoute, ok := nodeRoutes[edge.Dest]; !ok || oldRoute.Cost > newRoute.Cost {
				nodeRoutes[edge.Dest] = newRoute
			}
		}
		visitedNodes[currentNode] = true
		currentNode = g.getNextNode(nodeRoutes, visitedNodes)
		if currentNode == b {
			break
		}
	}
	route, ok := nodeRoutes[b]
	if !ok {
		return nil, ErrUnreachable
	}
	return &route, nil
}
