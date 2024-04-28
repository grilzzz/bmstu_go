package main

import (
	"fmt"
)

type Edge struct {
	s   *State
	out string
}

type State struct {
	order  int
	marked bool
	edges  []*Edge
}

var n, m, q0, x1, num int
var x2 string
var transition_matr [][]int
var out_matr [][]string

func main() {
	alphabet := []string{} // Инициализация пустой строки для алфавита
	for c := 'a'; c <= 'z'; c++ {
		alphabet = append(alphabet, string(c)) // Добавление каждой буквы к строке
	}
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

	fmt.Println("digraph {")
	fmt.Println("	rankdir = LR")

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf(`	%d -> %d [label = "%s(%s)"]`, i, transition_matr[i][j], alphabet[j], out_matr[i][j])
			fmt.Println()
		}
	}
	fmt.Println("}")
}
