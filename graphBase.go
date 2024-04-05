package main

import (
	"fmt"
	"sort"
)

type Component struct {
	order    int
	vertices []*Vertex
	isKoren  bool
}

type Vertex struct {
	order int
	edges []*Vertex
	T1    int
	T2    int
	low   int
	comp  int
}

type Stack struct {
	vertices []*Vertex
}

func (s *Stack) InitStack() {
	s.vertices = []*Vertex{}
}

func (s *Stack) Push(v *Vertex) {
	s.vertices = append(s.vertices, v)
}

func (s *Stack) Pop() *Vertex {
	ret := s.vertices[len(s.vertices)-1]
	s.vertices = s.vertices[:len(s.vertices)-1]
	return ret
}

func Tarjan(n int) {
	s.InitStack()
	for _, v := range g {
		if v.T1 == 0 {
			VisitVertex_Tarjan(v)
		}
	}
}

func VisitVertex_Tarjan(v *Vertex) {
	v.T1 = time
	v.low = time
	time++
	s.Push(v)
	for _, u := range v.edges {

		if u.T1 == 0 {
			VisitVertex_Tarjan(u)
		}
		if u.comp == -1 && v.low > u.low {
			v.low = u.low
		}
	}
	if v.T1 == v.low {
		u := s.Pop()
		u.comp = count
		for u != v {
			u = s.Pop()
			u.comp = count
		}
		count++
	}
}

var N, M, u, v, time, count int
var s Stack
var g []*Vertex

func main() {
	time = 1
	count = 0
	g = []*Vertex{}
	s = Stack{}

	fmt.Scan(&N)
	fmt.Scan(&M)
	for i := 0; i < N; i++ {
		g = append(g, &Vertex{i, nil, 0, 0, 0, -1})
	}

	for i := 0; i < M; i++ {
		fmt.Scan(&u, &v)
		g[u].edges = append(g[u].edges, g[v])
	}

	Tarjan(N)

	condencation := []*Component{}

	for i := 0; i < count; i++ {
		v_arr := []*Vertex{}
		for _, v := range g {
			if v.comp == i {
				v_arr = append(v_arr, v)
			}
		}
		condencation = append(condencation, &Component{i, v_arr, true})
	}

	for _, v := range g {
		for _, u := range v.edges {
			if v.comp != u.comp {
				condencation[u.comp].isKoren = false
			}
		}
	}

	base := []*Component{}

	for _, c := range condencation {
		if c.isKoren {
			base = append(base, c)
		}
	}
	arr := []int{}
	for _, c := range base {
		arr = append(arr, c.vertices[0].order)
	}
	sort.Ints(arr)
	for _, i := range arr {
		fmt.Printf("%d ", i)
	}
}
