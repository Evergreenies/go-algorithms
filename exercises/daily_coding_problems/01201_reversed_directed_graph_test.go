package dailycodingproblems

import (
	"fmt"
	"reflect"
	"testing"
)

/**
* Write an algorithm that computes the reversal of a directed graph.
* For example, if a graph consists of A -> B -> C, it should become A <- B <- C.
* **/

type Graph map[interface{}][]interface{}

func reversedDirectedGraph(graph Graph) Graph {
	reversedDirected := make(Graph)
	for vertex, neighbors := range graph {
		for _, neighbor := range neighbors {
			_, ok := reversedDirected[vertex]
			if !ok {
				reversedDirected[vertex] = make([]interface{}, 0)
			}

			reversedDirected[neighbor] = append(reversedDirected[neighbor], vertex)
		}
	}

	return reversedDirected
}

func assert1201(t *testing.T, actual interface{}, expected interface{}) {
	if reflect.DeepEqual(actual, expected) {
		fmt.Printf("PASSED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
	} else {
		t.Fatalf("FAILED: \n\tActual: %v\n\tExpected: %v\n", actual, expected)
		t.Fail()
	}
}

func TestReversedDirectedGraph(t *testing.T) {
	originalGraph := Graph{
		"A": {"B", "C"},
		"B": {"C"},
		"C": {},
	}
	expectedGraph := Graph{
		"A": {},
		"B": {"A"},
		"C": {"A", "B"},
	}
	assert1201(t, reversedDirectedGraph(originalGraph), expectedGraph)
}
