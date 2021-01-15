package main

import "fmt"

func main() {
	var (
		c   byte
		ans [10]int
	)
	for fmt.Scanf("%c", &c); c != '\r'; fmt.Scanf("%c", &c) {
		ans[c-'0']++
	}
	for i := 0; i < 10; i++ {
		if ans[i] > 0 {
			fmt.Printf("%d:%d\n", i, ans[i])
		}
	}
}
