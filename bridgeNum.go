package main

import (
	"fmt"
)

type Vertex struct {
	name  int
	color int
}

func dfs(graph map[int][]int, vertexes []Vertex, a, b int) bool {
	stack := []int{}
	ch_color := []int{}
	key := a
	if vertexes[key].color == 0 {
		stack = append(stack, key)
		for len(stack) > 0 {
			var last_elem int
			last_elem, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if vertexes[last_elem].color == 0 {
				vertexes[last_elem].color = 1

				ch_color = append(ch_color, last_elem)
				if last_elem == b {
					// for i := 0; i < len(ch_color); i++ {
					// 	vertexes[ch_color[i]].color = 0
					// }
					return true
				}

				stack = append(stack, last_elem)

				for _, elem := range graph[last_elem] {
					if elem == b && last_elem == a {
						continue
					}
					if elem == a && last_elem == b {
						continue
					}
					if vertexes[elem].color == 0 {
						stack = append(stack, elem)
					}
				}
			} else if vertexes[last_elem].color == 1 {
				vertexes[last_elem].color = 2
			}
		}
	}
	// for i := 0; i < len(ch_color); i++ {
	// 	vertexes[ch_color[i]].color = 0
	// }
	return false
}

func main() {
	var n_elements, n_edges int
	graph := make(map[int][]int)
	vertexes := []Vertex{}
	fmt.Scan(&n_elements)
	fmt.Scan(&n_edges)
	edges := []int{}
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
	// fmt.Println(subgraphs)
	counter := 0
	for i := 0; i < n_edges; i++ {
		copySlice := make([]Vertex, len(vertexes))
		copy(copySlice, vertexes)
		ans := dfs(graph, copySlice, edges[2*i], edges[2*i+1])
		if !ans {
			counter++
		}
	}
	fmt.Println(counter)
}
