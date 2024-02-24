package main

import (
	"bufio"
	"fmt"
	"os"
)

var i int = -1

func decode(s []rune, res int) int {
	i++

	switch s[i] {
	case '+':
		return decode(s, 0) + decode(s, 0)
	case '-':
		return decode(s, 0) - decode(s, 0)
	case '*':
		return decode(s, 0) * decode(s, 0)
	case '(':
		return decode(s, res)
	case ')':
		return decode(s, res)
	case ' ':
		return decode(s, res)
	default:
		return int(s[i] - '0')
	}

}

func main() {
	cin := bufio.NewScanner(os.Stdin)
	cin.Scan()
	s := []rune(cin.Text())
	res := decode(s, 0)
	fmt.Println(res)
}
