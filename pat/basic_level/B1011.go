package main

import "fmt"

func main() {
	var a, b, c int64
	t := 0
	fmt.Scanf("%d\n", &t)
	for i := 1; i <= t; i++ {
		fmt.Scanf("%d %d %d\n", &a, &b, &c)
		fmt.Printf("Case #%d: %v\n", i, a+b > c)
	}
}
