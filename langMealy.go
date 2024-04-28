package main

import "fmt"

var numQ, numIn, startQ, m int
var QInToQ [][]int
var QInToOut [][]string

func main() {
	numIn = 2
	fmt.Scan(&numQ)
	for i := 0; i < numQ; i++ {
		QInToQ = append(QInToQ, []int{})
		for j := 0; j < numIn; j++ {
			QInToQ[i] = append(QInToQ[i], 0)
		}
		for j := 0; j < numIn; j++ {
			fmt.Scan(&QInToQ[i][j])
		}
	}

	for i := 0; i < numQ; i++ {
		QInToOut = append(QInToOut, []string{})
		for j := 0; j < numIn; j++ {
			QInToOut[i] = append(QInToOut[i], "")
		}
		for j := 0; j < numIn; j++ {
			fmt.Scan(&QInToOut[i][j])
		}
	}
	fmt.Scan(&startQ, &m)

	ans := generateLanguage(startQ, QInToQ, QInToOut, m)

	for key := range ans {
		fmt.Printf("%s ", key)
	}

}

type state struct {
	word string
	curQ int
}

type edge struct {
	from int
	to   int
}

var recoursiveGenerate func(curLen int, word string, curQ int, lambdaQ map[edge]struct{})

func generateLanguage(startQ int, QInToQ [][]int, QInToOut [][]string, maxLen int) map[string]struct{} {
	language := map[string]struct{}{}

	memo := map[state]struct{}{}

	recoursiveGenerate = func(curLen int, word string, curQ int, lambdaQ map[edge]struct{}) {
		memo[state{word, curQ}] = struct{}{}

		if curLen == maxLen || word != "" {
			language[word] = struct{}{}
			if curLen == maxLen {
				return
			}
		}

		for i := range QInToQ[curQ] {

			newQ := QInToQ[curQ][i]
			newSymbol := QInToOut[curQ][i]
			newEdge := edge{curQ, newQ}

			if _, ok := memo[state{word + newSymbol, newQ}]; !ok {
				if _, ok := lambdaQ[newEdge]; !ok {
					if newSymbol != "-" {
						recoursiveGenerate(curLen+1, word+newSymbol, newQ, map[edge]struct{}{})
					} else if newQ != curQ {
						lambdaQ[newEdge] = struct{}{}
						recoursiveGenerate(curLen, word, newQ, lambdaQ)
						delete(lambdaQ, newEdge)
					}
				}
			}
		}
	}

	recoursiveGenerate(0, "", startQ, map[edge]struct{}{})
	return language
}
