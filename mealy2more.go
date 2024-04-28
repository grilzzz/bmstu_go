package main

import "fmt"

type mealy struct {
	numQ, numIn, numOut int
	alphaIn, alphaOut   []string
	QInToQ              [][]int
	QInToOut            [][]int
}
type mooreQ struct {
	qIndMealy int
	outValue  int
}
type moore struct {
	numQ, numIn, numOut int
	alphaIn, alphaOut   []string
	QInToQ              [][]int
	QInToOut            []int
	mealyQOutToMoorQ    map[mooreQ]int
	moorQToMealy        map[int]mooreQ
}

func main() {
	autMealy := mealy{}
	fmt.Scan(&autMealy.numIn)
	for i := 0; i < autMealy.numIn; i++ {
		v := ""
		fmt.Scan(&v)
		autMealy.alphaIn = append(autMealy.alphaIn, v)
	}
	fmt.Scan(&autMealy.numOut)
	for i := 0; i < autMealy.numOut; i++ {
		v := ""
		fmt.Scan(&v)
		autMealy.alphaOut = append(autMealy.alphaOut, v)
	}
	fmt.Scan(&autMealy.numQ)
	for i := 0; i < autMealy.numQ; i++ {
		autMealy.QInToQ = append(autMealy.QInToQ, []int{})
		for j := 0; j < autMealy.numIn; j++ {
			v := 0
			fmt.Scan(&v)
			autMealy.QInToQ[i] = append(autMealy.QInToQ[i], v)
		}
	}
	for i := 0; i < autMealy.numQ; i++ {
		autMealy.QInToOut = append(autMealy.QInToOut, []int{})
		for j := 0; j < autMealy.numIn; j++ {
			v := 0
			fmt.Scan(&v)
			autMealy.QInToOut[i] = append(autMealy.QInToOut[i], v)
		}
	}

	m := mealyToMoore(&autMealy)

	fmt.Println("digraph {")
	fmt.Println("	rankdir = LR")
	for i := 0; i < m.numQ; i++ {
		fmt.Printf(`	%d [label = "(%d,%s)"]`, i, m.moorQToMealy[i].qIndMealy, m.alphaOut[m.QInToOut[i]])
		fmt.Println()
	}
	for i := 0; i < m.numQ; i++ {
		for j := 0; j < m.numIn; j++ {
			fmt.Printf(`	%d -> %d [label = "%s"]`, i, m.QInToQ[i][j], m.alphaIn[j])
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func mealyToMoore(autMealy *mealy) moore {
	mealyQOutToMoorQ := map[mooreQ]int{}
	moorQToMealy := map[int]mooreQ{}
	mooreNumQ := 0

	for i := 0; i < autMealy.numQ; i++ {
		for j := 0; j < autMealy.numIn; j++ {
			mQ := mooreQ{autMealy.QInToQ[i][j], autMealy.QInToOut[i][j]}
			if _, ok := mealyQOutToMoorQ[mQ]; !ok {
				mooreNumQ += 1
				moorQToMealy[mooreNumQ-1] = mQ
				mealyQOutToMoorQ[mQ] = mooreNumQ - 1
			}

		}
	}

	autMoore := moore{
		numQ:             mooreNumQ,
		numIn:            autMealy.numIn,
		numOut:           autMealy.numOut,
		alphaIn:          autMealy.alphaIn,
		alphaOut:         autMealy.alphaOut,
		QInToQ:           nil,
		QInToOut:         nil,
		mealyQOutToMoorQ: mealyQOutToMoorQ,
		moorQToMealy:     moorQToMealy,
	}

	autMoore.QInToOut = make([]int, mooreNumQ)
	for i := range autMoore.QInToOut {
		autMoore.QInToOut[i] = moorQToMealy[i].outValue
	}

	autMoore.QInToQ = make([][]int, autMoore.numQ)
	for i := range autMoore.QInToQ {
		autMoore.QInToQ[i] = make([]int, autMoore.numIn)
		for j := range autMoore.QInToQ[i] {
			mealyQ := moorQToMealy[i].qIndMealy
			mQ := mooreQ{qIndMealy: autMealy.QInToQ[mealyQ][j], outValue: autMealy.QInToOut[mealyQ][j]}
			autMoore.QInToQ[i][j] = mealyQOutToMoorQ[mQ]
		}
	}
	return autMoore
}
