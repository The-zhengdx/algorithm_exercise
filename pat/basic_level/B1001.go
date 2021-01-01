package main

import "fmt"

func main() {
	var n, step int
	fmt.Scanf("%d", &n)
	for n != 1 {
		if n%2 == 1 {
			n = 3*n + 1
		}
		n /= 2
		step++
	}
	fmt.Println(step)
}
