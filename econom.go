package main

import (
	"bufio"
	"fmt"
	"os"
)

func econom(s []rune) int {
	set := map[string]struct{}{}
	set[string(s)] = struct{}{}
	if s[0] != '(' {
		return 0
	}
	first_i := 0
	i := 0
	for i < len(s) {
		i = first_i
		c := 1
		for j := i + 1; j < len(s); j++ {
			if s[j] == '(' {
				c++
				if first_i == i {
					first_i = j
				}
			}
			if s[j] == ')' {
				c--
			}
			if c == 0 {
				elem := s[i : j+1]
				set[string(elem)] = struct{}{}
				break
			}
		}
		if first_i == i {
			for j := i + 1; j < len(s); j++ {
				if s[j] == '(' {
					first_i = j
					break
				}
			}
		}
		if first_i == i {
			break
		}
	}
	return len(set)
}

func main() {
	cin := bufio.NewScanner(os.Stdin)
	cin.Scan()
	s := []rune(cin.Text())
	fmt.Println(econom(s))
}
