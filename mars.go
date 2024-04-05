package main

import (
	"fmt"
	"math"
	"sort"
)

type Vertex struct {
	name  int
	group int
	color int
	deep  []int
}

type Group struct {
	min_elem int
	n_elems  int
	elems    []int
}

type Subgraph struct {
	group1 Group
	group2 Group
	diff   int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type ByMinElem []Subgraph

func (a ByMinElem) Len() int { return len(a) }
func (a ByMinElem) Less(i, j int) bool {
	return min(a[i].group1.min_elem, a[i].group2.min_elem) < min(a[j].group1.min_elem, a[j].group2.min_elem)
}
func (a ByMinElem) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByDiff []Subgraph

func (a ByDiff) Len() int { return len(a) }
func (a ByDiff) Less(i, j int) bool {
	return math.Abs(float64(a[i].group1.n_elems-a[i].group2.n_elems)) < math.Abs(float64(a[i].group1.n_elems-a[i].group2.n_elems))
}
func (a ByDiff) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func checkColor(deep []int, group1, group2 int) bool {
	if group1 == -1 {
		return true
	}
	for _, i := range deep {
		if i%2 == 0 && group1 != group2 {
			return false
		}
		if i%2 == 1 && group1 != group2+1 {
			return false
		}
	}
	return true
}

func addVertex(deep []int, last_elem, group int, subgraph Subgraph, info []Vertex) (Subgraph, []Vertex) {
	if deep[0]%2 == 0 {
		info[last_elem].group = group
		subgraph.group1.n_elems++
		subgraph.group1.elems = append(subgraph.group1.elems, last_elem)
		if last_elem < subgraph.group1.min_elem {
			subgraph.group1.min_elem = last_elem
		}
	} else {
		info[last_elem].group = group + 1
		subgraph.group2.n_elems++
		subgraph.group2.elems = append(subgraph.group2.elems, last_elem)
		if last_elem < subgraph.group2.min_elem {
			subgraph.group2.min_elem = last_elem
		}
	}
	return subgraph, info
}

func dfs1(graph map[int][]int, group int, info []Vertex) (bool, []Subgraph, []Vertex) {
	isPossible := true
	stack := []int{}
	subgraphs := []Subgraph{}
	for key, _ := range graph {
		info[key].deep = append(info[key].deep, 0)
		if info[key].color == 0 {
			subgraph := Subgraph{group1: Group{min_elem: 99999999999, n_elems: 0, elems: []int{}}, group2: Group{min_elem: 99999999999, n_elems: 0, elems: []int{}}}
			stack = append(stack, key)
			for len(stack) > 0 {
				var last_elem int
				last_elem, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if isPossible {
					isPossible = checkColor(info[last_elem].deep, info[last_elem].group, group)
				}
				if info[last_elem].color == 0 {
					info[last_elem].color = 1

					if info[last_elem].group == -1 {
						subgraph, info = addVertex(info[last_elem].deep, last_elem, group, subgraph, info)
					}

					stack = append(stack, last_elem)
					for _, elem := range graph[last_elem] {
						for _, i := range info[last_elem].deep {
							info[elem].deep = append(info[elem].deep, i+1)
						}
						if info[elem].color == 0 {
							stack = append(stack, elem)
						}
					}
				} else if info[last_elem].color == 1 {
					info[last_elem].color = 2
				}
			}
			subgraphs = append(subgraphs, subgraph)
			group += 2
		}
	}
	return isPossible, subgraphs, info
}

func getMinDiff(subgraphs []Subgraph) int {
	var diff int
	sort.Sort(ByDiff(subgraphs))
	for i := len(subgraphs) - 1; i >= 0; i-- {
		diff = min(int(math.Abs(float64(diff-(subgraphs[i].group1.n_elems-subgraphs[i].group2.n_elems)))), int(math.Abs(float64(diff+(subgraphs[i].group1.n_elems-subgraphs[i].group2.n_elems)))))
	}
	return int(math.Abs(float64(diff)))
}

func getBinN(n int, length int) []int {
	ans := []int{}
	for n > 0 {
		ans = append(ans, n%2)
		n = n / 2
	}
	for length > len(ans) {
		ans = append(ans, 0)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func getFirstGroup(subgraphs []Subgraph, n int) []int {
	sort.Sort(ByMinElem(subgraphs))
	for i := 0; i < int(math.Pow(2, float64(len(subgraphs)))); i++ {
		ans := []int{}
		arr := getBinN(i, len(subgraphs))
		left := n
		for ind, val := range arr {
			if val == 0 {
				if subgraphs[ind].group1.min_elem < subgraphs[ind].group2.min_elem {
					left -= subgraphs[ind].group1.n_elems
					ans = append(ans, subgraphs[ind].group1.elems...)
				} else {
					left -= subgraphs[ind].group2.n_elems
					ans = append(ans, subgraphs[ind].group2.elems...)
				}
			} else {
				if subgraphs[ind].group2.min_elem < subgraphs[ind].group1.min_elem {
					left -= subgraphs[ind].group1.n_elems
					ans = append(ans, subgraphs[ind].group1.elems...)
				} else {
					left -= subgraphs[ind].group2.n_elems
					ans = append(ans, subgraphs[ind].group2.elems...)
				}
			}
		}
		if left == 0 {
			return ans
		}
	}
	return []int{}
}

func printPlusOne(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i]++
		fmt.Printf("%d ", arr[i])
	}
}

func main() {
	var x int
	fmt.Scanf("%d", &x)
	graph := make(map[int][]int)

	var info []Vertex
	for i := 0; i < x; i++ {
		info = append(info, Vertex{color: 0, name: i, group: -1})
	}
	for i := 0; i < x; i++ {
		graph[i] = []int{}
		for j := 0; j < x; j++ {
			var a string
			fmt.Scan(&a)
			if a == "+" {
				graph[i] = append(graph[i], j)
			}
		}
	}
	var isPossible bool
	var subgraphs []Subgraph
	isPossible, subgraphs, info = dfs1(graph, 0, info)
	if isPossible {
		a := getMinDiff(subgraphs)
		n := x/2 - a/2
		ans := getFirstGroup(subgraphs, n)
		sort.Ints(ans)
		printPlusOne(ans)
	} else {
		fmt.Println("No solution")
	}
}
