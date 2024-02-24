package main

import (
	"fmt"
)

type elem struct {
	num, den int
}

func sum_elems(x1 elem, x2 elem) elem {
	x1.num, x1.den = x1.num*x2.den+x1.den*x2.num, x1.den*x2.den
	return x1
}

func mul_elems(x1 elem, x2 elem) elem {
	x1.num, x1.den = x1.num*x2.num, x1.den*x2.den
	return x1
}

func div_elems(x1 elem, x2 elem) elem {
	x1.num, x1.den = x1.num*x2.den, x1.den*x2.num
	return x1
}

func dif_elems(x1 elem, x2 elem) elem {
	x1.num, x1.den = x1.num*x2.den-x1.den*x2.num, x1.den*x2.den
	return x1
}

func gauss(matrix [][]elem, answers []elem) []elem {
	for i := range matrix {

		if matrix[i][i].num == 0 {
			for j := i + 1; j < len(matrix); j++ {
				if matrix[j][i].num != 0 {
					matrix[i], matrix[j] = matrix[j], matrix[i]
					answers[i], answers[j] = answers[j], answers[i]
					break
				}
			}
		}
		x := matrix[i][i]
		for j := range matrix {
			matrix[i][j] = reduce(div_elems(matrix[i][j], x))
		}
		answers[i] = reduce(div_elems(answers[i], x))

		for j := range matrix {
			if i == j {
				continue
			}
			x := matrix[j][i]
			for k := range matrix {
				matrix[j][k] = reduce(dif_elems(matrix[j][k], mul_elems(matrix[i][k], x)))
			}
			answers[j] = reduce(dif_elems(answers[j], mul_elems(answers[i], x)))
		}
	}

	return answers
}

func nod(a, b int) int {
	if a > b {
		a, b = b, a
	}
	for a != 0 {
		b, a = a, b%a
	}
	return b
}

func reduce(x1 elem) elem {
	if x1.den == 0 {
		return x1
	}
	x := nod(x1.den, x1.num)
	x1.num = x1.num / x
	x1.den = x1.den / x
	if x1.den < 0 {
		x1.num, x1.den = -x1.num, -x1.den
	}
	return x1
}

func check(answers []elem) bool {
	ret := true
	for i := range answers {
		if answers[i].den == 0 {
			ret = false
		}
	}
	return ret
}

func main() {
	var n, x int
	fmt.Scan(&n)
	matrix := make([][]elem, n)
	for i := range matrix {
		matrix[i] = make([]elem, n)
	}

	answers := make([]elem, n)

	for i := range matrix {
		for j := range matrix {
			fmt.Scan(&x)
			matrix[i][j] = elem{x, 1}
		}
		fmt.Scan(&x)
		answers[i] = elem{x, 1}
	}

	answers = gauss(matrix, answers)
	if check(answers) {
		for i := range answers {
			answers[i] = reduce(answers[i])
			fmt.Printf("%d/%d\n", answers[i].num, answers[i].den)
		}
	} else {
		fmt.Println("No solution")
	}
}
