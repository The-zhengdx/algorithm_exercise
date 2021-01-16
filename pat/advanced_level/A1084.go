package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s\n%s", &s1, &s2)
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)
	i, j, l1, l2, m := 0, 0, len(s1), len(s2), make(map[byte]bool)
	for ; j < l2; i++ {
		if s1[i] == s2[j] {
			j++
		} else if output := m[s1[i]]; !output {
			m[s1[i]] = true
			fmt.Printf("%c", s1[i])
		}
	}
	for ; i < l1; i++ {
		if output := m[s1[i]]; !output {
			m[s1[i]] = true
			fmt.Printf("%c", s1[i])
		}
	}
}
