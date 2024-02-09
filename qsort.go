package main

import "fmt"

var a = []int{45, 7, 3, 7, 25, 872, 456, 27, 27, 5}

func less(i, j int) bool {
	if a[i] < a[j] {
		return true
	} else {
		return false
	}
}

func swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partition(f, l int, less func(i, j int) bool, swap func(i, j int)) int {
	i, j := f, f
	for i < l {
		if less(i, l) {
			swap(i, j)
			j++
		}
		i++
	}
	swap(j, l)
	return j
}

func quick_sort(f, l int, less func(i, j int) bool, swap func(i, j int)) {
	if f < l {
		q := partition(f, l, less, swap)
		quick_sort(f, q-1, less, swap)
		quick_sort(q+1, l, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quick_sort(0, n-1, less, swap)
}

func main() {
	qsort(len(a), less, swap)
	fmt.Println(a)
}
