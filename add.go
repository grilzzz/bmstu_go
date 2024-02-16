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
	for i < len(a) {
		ans = append(ans, c/int32(p))
		c = a[i] + b[i]
		ans[i] += c % int32(p)
		i++
	}
	if c/int32(p) > 0 {
		ans = append(ans, c/int32(p))
	}
	return ans
}

func main() {
	a := []int32{5, 2}
	b := []int32{6, 0, 5}

	fmt.Println(add(a, b, 10))
}
