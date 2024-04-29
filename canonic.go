package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	s   *State
	out string
}

type States []*State

func (states States) Len() int           { return len(states) }
func (states States) Less(i, j int) bool { return states[i].order < states[j].order }
func (states States) Swap(i, j int)      { states[i], states[j] = states[j], states[i] }

type State struct {
	order  int
	marked bool
	edges  []*Edge
}

func dfs(s *State) {
	s.order = num
	num++
	s.marked = true

	for _, elem := range s.edges {
		if !elem.s.marked {
			dfs(elem.s)
		}
	}
}

var n, m, q0, x1, num int
var x2 string
var transition_matr [][]int
var out_matr [][]string
var states States

func main() {
	num = 0
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q0)

	transition_matr = [][]int{}
	for i := 0; i < n; i++ {
		transition_matr = append(transition_matr, []int{})
		for j := 0; j < m; j++ {
			fmt.Scan(&x1)
			transition_matr[i] = append(transition_matr[i], x1)
		}
	}

	out_matr = [][]string{}
	for i := 0; i < n; i++ {
		out_matr = append(out_matr, []string{})
		for j := 0; j < m; j++ {
			fmt.Scan(&x2)
			out_matr[i] = append(out_matr[i], x2)
		}
	}

	for i := 0; i < n; i++ {
		states = append(states, &State{i, false, []*Edge{}})
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			states[i].edges = append(states[i].edges, &Edge{states[transition_matr[i][j]], out_matr[i][j]})
		}
	}

	dfs(states[q0])

	sort.Sort(states)

	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(q0)

	for i := 0; i < n; i++ {
		for _, e := range states[i].edges {
			fmt.Printf("%d ", e.s.order)
		}
		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for _, e := range states[i].edges {
			fmt.Printf("%s ", e.out)
		}
		fmt.Println()
	}
}

// 7
// 3
// 0

// 5 1 4
// 4 6 1
// 3 6 0
// 1 6 4
// 0 1 0
// 2 5 1
// 6 3 3

// g g s
// a y g
// c s t
// b z s
// g n b
// b t z
// c z b
