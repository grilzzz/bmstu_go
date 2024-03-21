package main

import (
	"fmt"
	"math"
	"sort"
)

// Vertex представляет собой вершину графа с именем и цветом
type Vertex struct {
	name  int
	color int
}

// ByColor предоставляет методы для сортировки массива вершин по цвету
type ByColor []Vertex

func (a ByColor) Len() int           { return len(a) }
func (a ByColor) Less(i, j int) bool { return a[i].name < a[j].name }
func (a ByColor) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func deviders(x int) []Vertex {
	divs := []Vertex{}
	sqrt := math.Ceil(math.Pow(float64(x), 0.5))
	for i := 1; i < int(sqrt); i++ {
		if x%i == 0 {
			divs = append(divs, Vertex{name: i, color: 0}, Vertex{name: x / i, color: 0})
		}
	}
	sort.Sort(ByColor(divs))
	return divs
}

func dfs(lst []Vertex) {
	for i := len(lst) - 1; i > 0; i-- {
		if lst[i].color == 0 {
			visitVertex(lst, i)
		}
	}
}

func visitVertex(lst []Vertex, i int) {
	// fmt.Println(lst)
	lst[i].color = 1
	counter := 0
	for j := i - 1; j >= 0; j-- {
		if (lst[j].color == 0 || lst[j].name > lst[i].name/lst[j].name) && lst[i].name%lst[j].name == 0 {
			fmt.Printf("    %d -- %d;\n", lst[i].name, lst[j].name)
			counter++
			if lst[j].color == 0 {
				visitVertex(lst, j)
			}

		}
	}
	if counter == 0 && lst[i].name > 1 && lst[i].color != 2 {
		fmt.Printf("    %d -- %d;\n", lst[i].name, lst[0].name)
	}
	lst[i].color = 2
}

func main() {
	var x int
	fmt.Scanf("%d", &x)
	divs := deviders(x)
	fmt.Println("graph deviders {")
	for i := range divs {
		fmt.Printf("    %d;\n", divs[i].name)
	}
	dfs(divs)
	fmt.Println("}")
}
