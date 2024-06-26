package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	ans := []int32{}
	if len(a) > len(b) {
		a, b = b, a
	}
	for len(a) < len(b) {
		a = append(a, 0)
	}
	i := 0
	var c int32 = 0
	var c1 int32 = 0
	for i < len(a) {
		ans = append(ans, c1)
		c = a[i] + b[i]
		c1 = (ans[i] + c) / int32(p)
		ans[i] = (ans[i] + c) % int32(p)
		i++
	}
	if c1 > 0 {
		ans = append(ans, c1)
	}
	return ans
}

func main() {
	a := []int32{5, 2}
	b := []int32{6, 7, 5}

	fmt.Println(add(a, b, 10))
}
