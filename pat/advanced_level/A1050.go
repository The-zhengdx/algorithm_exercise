package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s1 := sc.Text()
	sc.Scan()
	s2 := sc.Text()
	m1, m2 := make([]bool, len(s1)), [128]bool{}
	for i, l := 0, len(s2); i < l; i++ {
		m2[s2[i]] = true
	}
	for i, l := 0, len(s1); i < l; i++ {
		if m2[s1[i]] {
			m1[i] = true
		}
	}
	for i, l := 0, len(s1); i < l; i++ {
		if !m1[i] {
			fmt.Printf("%c", s1[i])
		}
	}
}
