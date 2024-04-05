package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	name  int
	color int
}

type Subgraph struct {
	min_elem int
	n_elems  int
	n_edges  int
	elements []int
}

type ByMinElem []Subgraph

func (a ByMinElem) Len() int           { return len(a) }
func (a ByMinElem) Less(i, j int) bool { return a[j].min_elem < a[i].min_elem }
func (a ByMinElem) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByElems []Subgraph

func (a ByElems) Len() int           { return len(a) }
func (a ByElems) Less(i, j int) bool { return a[i].n_elems < a[j].n_elems }
func (a ByElems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByEdges []Subgraph

func (a ByEdges) Len() int           { return len(a) }
func (a ByEdges) Less(i, j int) bool { return a[i].n_edges < a[j].n_edges }
func (a ByEdges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func dfs(graph map[int][]int, vertexes []Vertex) []Subgraph {
	stack := []int{}
	subgraphs := []Subgraph{}
	for key, _ := range graph {

		if vertexes[key].color == 0 {
			subgraph := Subgraph{min_elem: 99999999999999, n_elems: 0, n_edges: 0, elements: []int{}}
			stack = append(stack, key)
			for len(stack) > 0 {
				var last_elem int
				last_elem, stack = stack[len(stack)-1], stack[:len(stack)-1]

				if vertexes[last_elem].color == 0 {
					vertexes[last_elem].color = 1

					subgraph.n_elems++
					subgraph.elements = append(subgraph.elements, last_elem)
					if subgraph.min_elem > last_elem {
						subgraph.min_elem = last_elem
					}

					stack = append(stack, last_elem)

					for _, elem := range graph[last_elem] {
						subgraph.n_edges++
						if vertexes[elem].color == 0 {
							stack = append(stack, elem)
						}
					}
				} else if vertexes[last_elem].color == 1 {
					vertexes[last_elem].color = 2
				}
			}
			subgraphs = append(subgraphs, subgraph)
		}
	}
	return subgraphs
}

func whoIsRed(subgraphs []Subgraph) Subgraph {
	sort.Sort(ByMinElem(subgraphs))
	sort.Sort(ByEdges(subgraphs))
	sort.Sort(ByElems(subgraphs))
	return subgraphs[len(subgraphs)-1]
}

func contains(arr []int, x int) bool {
	for _, i := range arr {
		if i == x {
			return true
		}
	}
	return false
}

func main() {
	var n_elements, n_edges int
	vertexes := []Vertex{}
	edges := []int{}
	graph := make(map[int][]int)
	fmt.Scan(&n_elements)
	fmt.Scan(&n_edges)
	for i := 0; i < n_elements; i++ {
		vertexes = append(vertexes, Vertex{i, 0})
	}
	for i := 0; i < n_edges; i++ {
		var a, b int

		fmt.Scan(&a)
		fmt.Scan(&b)
		edges = append(edges, a, b)
		if graph[a] == nil {
			graph[a] = []int{b}
		} else {
			graph[a] = append(graph[a], b)
		}
		if graph[b] == nil {
			graph[b] = []int{a}
		} else {
			graph[b] = append(graph[b], a)
		}
	}
	subgraphs := dfs(graph, vertexes)
	fmt.Println(subgraphs)
	red_elements := whoIsRed(subgraphs)
	fmt.Println("graph {")
	for i := 0; i < n_elements; i++ {
		if contains(red_elements.elements, i) {
			fmt.Printf("    %d [color = red]\n", i)
		} else {
			fmt.Printf("    %d\n", i)
		}
	}
	for i := 0; i < n_edges; i++ {
		if contains(red_elements.elements, edges[2*i]) || contains(red_elements.elements, edges[2*i+1]) {
			fmt.Printf("    %d -- %d [color = red]\n", edges[2*i], edges[2*i+1])
		} else {
			fmt.Printf("    %d -- %d\n", edges[2*i], edges[2*i+1])
		}
	}
	fmt.Println("}")
}
