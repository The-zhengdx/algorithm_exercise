package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)

	s, e := make([]int, 0, n), 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &e)
		s = append(s, e)
	}

	m %= n
	reverse(s)
	reverse(s[:m])
	reverse(s[m:])

	l := len(s) - 1
	for i := 0; i < l; i++ {
		fmt.Printf("%d ", s[i])
	}
	fmt.Print(s[l])

}

func reverse(s []int) {
	for i, l := 0, len(s); i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
