package dijkstra_test

import (
	"testing"

	"github.com/efueyo/dijkstra"
)

func TestContainsEmpty(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{})
	if g.Contains(dijkstra.Node("D")) {
		t.Errorf("Graph %v should not contain edge D", g)
	}
}
func TestContainsOk(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("B"),
			Dest:   dijkstra.Node("C"),
			Cost:   4,
		},
	})
	all := g.Contains(dijkstra.Node("A")) && g.Contains(dijkstra.Node("B")) && g.Contains(dijkstra.Node("C"))
	if !all {
		t.Errorf("Graph %v should contain edges A, B & C", g)
	}
}
func TestContainsNotOk(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("B"),
			Dest:   dijkstra.Node("C"),
			Cost:   4,
		},
	})
	if g.Contains(dijkstra.Node("D")) {
		t.Errorf("Graph %v should not contain edge D", g)
	}
}

func TestDijsktraStraight(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
	})
	d, err := g.Distance(dijkstra.Node("A"), dijkstra.Node("B"))
	if err != nil {
		t.Error(err)
	}
	if d.Cost != 3 {
		t.Errorf("Expected distance of %v, got %v", 3, d.Cost)
	}
}

func TestDijsktraInvalidNodes(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
	})
	_, err := g.Distance(dijkstra.Node("C"), dijkstra.Node("B"))
	if err == nil {
		t.Error("Expected error with non existing source")
	}
	_, err = g.Distance(dijkstra.Node("A"), dijkstra.Node("C"))
	if err == nil {
		t.Error("Expected error with non existing dest")
	}
}

func TestDijsktraUnreachable(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("C"),
			Dest:   dijkstra.Node("B"),
			Cost:   5,
		},
	})
	_, err := g.Distance(dijkstra.Node("A"), dijkstra.Node("C"))
	if err != dijkstra.ErrUnreachable {
		t.Errorf("Expected error ErrUnreachable, got %v", err)
	}
}

func TestDijsktraAsymetric(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   3,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("B"),
			Dest:   dijkstra.Node("A"),
			Cost:   5,
		},
	})
	d, err := g.Distance(dijkstra.Node("A"), dijkstra.Node("B"))
	if err != nil {
		t.Error(err)
	}
	if d.Cost != 3 {
		t.Errorf("Expected distance of %v, got %v", 3, d.Cost)
	}
	d, err = g.Distance(dijkstra.Node("B"), dijkstra.Node("A"))
	if err != nil {
		t.Error(err)
	}
	if d.Cost != 5 {
		t.Errorf("Expected distance of %v, got %v", 5, d.Cost)
	}
}

func TestDijsktraHops(t *testing.T) {
	g := dijkstra.NewGraph([]dijkstra.Edge{
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("B"),
			Cost:   5,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("A"),
			Dest:   dijkstra.Node("C"),
			Cost:   10,
		},
		dijkstra.Edge{
			Source: dijkstra.Node("B"),
			Dest:   dijkstra.Node("C"),
			Cost:   4,
		},
	})
	d, err := g.Distance(dijkstra.Node("A"), dijkstra.Node("C"))
	if err != nil {
		t.Error(err)
	}
	if d.Cost != 9 {
		t.Errorf("Expected distance of %v, got %v", 9, d.Cost)
	}
	if !equal([]dijkstra.Node{dijkstra.Node("A"), dijkstra.Node("B")}, d.Nodes) {
		t.Errorf("Expected distance of %v, got %v", []dijkstra.Node{dijkstra.Node("A"), dijkstra.Node("B")}, d.Nodes)
	}
}

func equal(a, b []dijkstra.Node) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
