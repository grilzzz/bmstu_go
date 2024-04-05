package main

import (
	"fmt"
	"math"
	"sort"
)

var sets = []map[Vertex]struct{}{}
var n_elements, a, b int
var ans = 0.0
var v Vertex
var edges []Edge
var elements []Vertex

type ByLength []Edge

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Less(i, j int) bool { return a[i].dist < a[j].dist }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Vertex struct {
	x int
	y int
}

type Edge struct {
	a    Vertex
	b    Vertex
	dist int
}

func getSet(a Vertex) int {
	for ind, set := range sets {
		_, ok := set[a]
		if ok {
			return ind
		}
	}
	return -1
}

func union(a, b int) {
	for elem := range sets[a] {
		_, ok := sets[b][elem]
		if !ok {
			sets[b][elem] = struct{}{}
		}
	}
	sets[a] = map[Vertex]struct{}{}
}

func f() {
	for i := 0; i < len(edges); i++ {
		n_set := getSet(edges[i].a)
		_, ok := sets[n_set][edges[i].b]
		if ok {
			continue
		}
		ans += math.Pow(float64(edges[i].dist), 0.5)

		union(n_set, getSet(edges[i].b))
	}
}

func main() {
	fmt.Scan(&n_elements)

	for i := 0; i < n_elements; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		v = Vertex{x: a, y: b}
		for _, elem := range elements {

			edge := Edge{a: v, b: elem, dist: (v.x-elem.x)*(v.x-elem.x) + (v.y-elem.y)*(v.y-elem.y)}

			edges = append(edges, edge)

		}
		elements = append(elements, v)
		sets = append(sets, map[Vertex]struct{}{})
		sets[i][v] = struct{}{}
	}
	sort.Sort(ByLength(edges))
	f()
	fmt.Printf("%.2f\n", ans)
}
