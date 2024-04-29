package main

import "fmt"

type vertex struct {
	dom     int
	sdom    int
	anc     int
	label   int
	parent  int
	time    int
	inEdge  []int
	outEdge []int
	lst     map[int]struct{}
}

func main() {
	start, graph := scanGraph()
	tin := numAndDelete(start, graph)
	buildDominatorsTree(graph, tin)

	countLoops := 0
	for v := range graph {
		for _, u := range graph[v].inEdge {
			for !(u == -1 || v == u) {
				u = graph[u].dom
			}
			if v == u {
				countLoops += 1
				break
			}
		}
	}
	fmt.Println(countLoops)
}

var from, to int
var tin []int
var time int
var used map[int]struct{}
var cmd string

func scanGraph() (int, map[int]*vertex) {
	g := make(map[int]*vertex)
	n := 0
	fmt.Scan(&n)

	prevInd := -1
	lastCmd := ""
	start := -1

	if n > 0 {
		fmt.Scan(&from, &cmd)
		g[from] = &vertex{
			dom:     -1,
			sdom:    from,
			label:   from,
			anc:     -1,
			parent:  -1,
			time:    -1,
			inEdge:  []int{},
			outEdge: []int{},
			lst:     map[int]struct{}{},
		}
		if cmd != "ACTION" {
			fmt.Scan(&to)
			g[to] = &vertex{
				dom:     -1,
				sdom:    to,
				label:   to,
				anc:     -1,
				parent:  -1,
				time:    -1,
				inEdge:  []int{},
				outEdge: []int{},
				lst:     map[int]struct{}{},
			}
			g[from].outEdge = append(g[from].outEdge, to)
			g[to].inEdge = append(g[to].inEdge, from)
		}
		start = from
		prevInd = from
		lastCmd = cmd
	}

	for i := 0; i < n-1; i++ {
		fmt.Scan(&from, &cmd)
		if _, ok := g[from]; !ok {
			g[from] = &vertex{
				dom:     -1,
				sdom:    from,
				label:   from,
				anc:     -1,
				parent:  -1,
				time:    -1,
				inEdge:  []int{},
				outEdge: []int{},
				lst:     map[int]struct{}{},
			}
		}
		if cmd != "ACTION" {
			fmt.Scan(&to)
			if _, ok := g[to]; !ok {
				g[to] = &vertex{
					dom:     -1,
					sdom:    to,
					label:   to,
					anc:     -1,
					parent:  -1,
					time:    -1,
					inEdge:  []int{},
					outEdge: []int{},
					lst:     map[int]struct{}{},
				}
			}
			g[to].inEdge = append(g[to].inEdge, from)
			g[from].outEdge = append(g[from].outEdge, to)

		}
		if lastCmd != "JUMP" {
			g[from].inEdge = append(g[from].inEdge, prevInd)
			g[prevInd].outEdge = append(g[prevInd].outEdge, from)
		}
		prevInd = from
		lastCmd = cmd
	}
	return start, g
}

func dfs(g *map[int]*vertex, parent, v int) {
	used[v] = struct{}{}
	tin = append(tin, v)
	(*g)[v].parent = parent
	(*g)[v].time = time
	time += 1
	for _, to := range (*g)[v].outEdge {
		if _, ok := used[to]; !ok {
			dfs(g, v, to)
		}
	}
}

func numAndDelete(s int, g map[int]*vertex) []int {
	tin = []int{}
	used = map[int]struct{}{}
	time = 0

	dfs(&g, -1, s)

	for ind, v := range g {
		if _, ok := used[ind]; ok {
			newInEdge := []int{}
			newOutEdge := []int{}
			for _, w := range v.outEdge {
				if _, ok := used[w]; ok {
					newOutEdge = append(newOutEdge, w)
				}
			}
			for _, w := range v.inEdge {
				if _, ok := used[w]; ok {
					newInEdge = append(newInEdge, w)
				}
			}
			v.inEdge = newInEdge
			v.outEdge = newOutEdge
		} else {
			delete(g, ind)
		}

	}
	return tin
}

func findMin(v int, g map[int]*vertex) (min int) {
	ind := v
	stack := []int{}
	n := 0
	for g[v].anc != -1 {
		stack = append(stack, v)
		n++
		v = g[v].anc
	}
	min = v
	for n != 0 {
		n--
		v = stack[n]
		if g[g[g[g[v].anc].label].sdom].time < g[g[g[v].label].sdom].time {
			g[v].label = g[g[v].anc].label
		}
		g[v].anc = min
	}
	min = g[ind].label
	return
}

func buildDominatorsTree(g map[int]*vertex, tin []int) {
	for i := 0; i < len(g); i++ {
		w := tin[len(g)-1-i]
		min := g[w].sdom
		for _, v := range g[w].inEdge {
			u := findMin(v, g)
			if g[g[u].sdom].time < g[min].time {
				min = g[u].sdom
			}
		}
		g[min].lst[w] = struct{}{}
		g[w].sdom = min
		g[w].anc = g[w].parent

		if w != tin[0] {
			for v := range g[g[w].parent].lst {
				u := findMin(v, g)
				if g[g[u].sdom].time == g[g[v].sdom].time {
					u = g[v].sdom
				}
				g[v].dom = u
			}
			g[g[w].parent].lst = map[int]struct{}{}
		}
	}

	for i := 1; i < len(g); i++ {
		if g[g[tin[i]].dom].time != g[g[tin[i]].sdom].time {
			g[tin[i]].dom = g[g[tin[i]].dom].dom
		}
	}
}
