package main

import "fmt"

func main() {
	var n, b int
	fmt.Scanf("%d %d\n", &n, &b)
	if n < b {
		fmt.Println("Yes")
		fmt.Println(n)
		return
	}
	nInB := make([]int, 0, 30)
	for n > 0 {
		nInB = append(nInB, n%b)
		n /= b
	}
	if isPalindromic(nInB) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	for i := len(nInB) - 1; i >= 0; i-- {
		fmt.Print(nInB[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
}
func isPalindromic(s []int) bool {
	for i, l := 0, len(s); i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
