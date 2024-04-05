package main

import (
	"fmt"
)

type Vertex struct {
	name    int
	color   int
	counter int
	dist    int
	isBase  int
}

type Queue struct {
	arr []int
	i   int
}

func enqueue(x int) {
	q.arr = append(q.arr, x)
}

func dequeue() int {
	if q.i >= len(q.arr) {
		return -1
	}
	q.i++
	return q.arr[q.i-1]
}

func isNotEmpty() bool {
	return q.i < len(q.arr)
}

var q = Queue{[]int{}, 0}
var n_elements, n_edges, k_base int
var vertexes = []Vertex{}
var graph = make(map[int][]int)
var base_list = []int{}

func bfs(vertex int) {
	for i := range vertexes {
		vertexes[i].color = 0
	}
	q.arr = []int{}
	q.i = 0

	key := vertex
	if vertexes[key].color == 0 {
		vertexes[key].color = 2
		enqueue(key)
		for isNotEmpty() {
			elem := dequeue()

			for _, next_elem := range graph[elem] {
				if vertexes[next_elem].isBase == 1 {
					continue
				}
				if vertexes[next_elem].color == 0 {
					vertexes[next_elem].color = 2
					if vertexes[next_elem].dist == -1 {
						vertexes[next_elem].dist = vertexes[elem].dist + 1
						vertexes[next_elem].counter++
						enqueue(next_elem)
					} else if vertexes[next_elem].dist > 0 {
						if vertexes[next_elem].dist == vertexes[elem].dist+1 {
							vertexes[next_elem].counter++
						}
						vertexes[next_elem].dist = vertexes[elem].dist + 1
						enqueue(next_elem)
					} else {
						fmt.Println("error")
					}
				}
			}
		}
	}
}

func main() {
	fmt.Scan(&n_elements)
	fmt.Scan(&n_edges)
	for i := 0; i < n_elements; i++ {
		vertexes = append(vertexes, Vertex{color: 0, counter: 0, name: i, dist: -1, isBase: 0})

	}

	for i := 0; i < n_edges; i++ {
		var a, b int

		fmt.Scan(&a)
		fmt.Scan(&b)

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

	fmt.Scan(&k_base)

	for i := 0; i < k_base; i++ {
		var x int
		fmt.Scan(&x)
		base_list = append(base_list, x)
		vertexes[x].dist = 0
		vertexes[x].isBase = 1
	}

	for _, vertex := range base_list {
		bfs(vertex)
		// fmt.Println(vertexes)
	}
	flag := false
	for _, elem := range vertexes {
		if elem.counter == k_base {
			flag = true
			fmt.Printf("%d ", elem.name)
		}
	}
	if !flag {
		fmt.Println("-")
	}
}
