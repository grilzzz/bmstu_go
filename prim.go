package main

import "fmt"

type Vetrex struct {
	index int
	key   int
	value *Vetrex
	edges []*Edge
}

type Edge struct {
	v   *Vetrex
	len int
}

type PQueue struct {
	lst   []*Vetrex
	count int
}

func InitPriorityQueue(pq *PQueue, n int) {
	pq.lst = make([]*Vetrex, n)
	pq.count = 0
}

func QueueEmpty(pq *PQueue) bool {
	return pq.count == 0
}

func Insert(pq *PQueue, v *Vetrex) {
	i := pq.count
	pq.count++
	pq.lst[i] = v
	for i > 0 && pq.lst[(i-1)/2].key > pq.lst[i].key {
		pq.lst[(i-1)/2], pq.lst[i] = pq.lst[i], pq.lst[(i-1)/2]
		pq.lst[i].index = i
		i = (i - 1) / 2
	}
	pq.lst[i].index = i
}

func ExtractMin(pq *PQueue) *Vetrex {
	res := pq.lst[0]
	pq.count--
	if pq.count > 0 {
		pq.lst[0] = pq.lst[pq.count]
		pq.lst[0].index = 0
		i := 0
		for true {
			l := 2*i + 1
			r := l + 1
			j := i
			if l < pq.count && pq.lst[i].key > pq.lst[l].key {
				i = l
			}
			if r < pq.count && pq.lst[i].key > pq.lst[r].key {
				i = r
			}
			if i == j {
				break
			}
			pq.lst[i], pq.lst[j] = pq.lst[j], pq.lst[i]
			pq.lst[i].index = i
			pq.lst[j].index = j
		}
	}
	return res
}

func DecreaseKey(pq *PQueue, v *Vetrex, newKey int) {
	i := v.index
	v.key = newKey
	for i > 0 && pq.lst[(i-1)/2].key > newKey {
		pq.lst[i], pq.lst[(i-1)/2] = pq.lst[(i-1)/2], pq.lst[i]
		pq.lst[i].index = i
		i = (i - 1) / 2
	}
	v.index = i
}

func prim(vlist []*Vetrex) int {
	var pq PQueue
	InitPriorityQueue(&pq, len(vlist))
	ans := 0
	v := vlist[0]
	for {
		v.index = -2
		for _, c := range v.edges {
			if c.v.index == -1 {
				c.v.key = c.len
				c.v.value = v
				Insert(&pq, c.v)
			} else if c.v.index != -2 && c.len < c.v.key {
				c.v.value = v
				DecreaseKey(&pq, c.v, c.len)
			}
		}
		if QueueEmpty(&pq) {
			break
		}
		v = ExtractMin(&pq)
		ans += v.key
	}
	return ans
}

func main() {
	var n, m, v, u, len int
	fmt.Scan(&n, &m)

	vlist := []*Vetrex{}
	for i := 0; i < n; i++ {
		vlist = append(vlist, &Vetrex{-1, 0, nil, nil})
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&v, &u, &len)
		vlist[v].edges = append(vlist[v].edges, &Edge{vlist[u], len})
		vlist[u].edges = append(vlist[u].edges, &Edge{vlist[v], len})
	}
	fmt.Printf("%d\n", prim(vlist))
}
