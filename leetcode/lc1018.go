package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	l := len(A)
	n, ans := 0, make([]bool, l)
	for i := 0; i < l; i++ {
		n = 2*(n%5) + A[i]
		fmt.Println(n)
		if n%5 == 0 {
			ans[i] = true
		}
	}
	return ans
}
