package main

import "fmt"

func main() {
	var (
		carry, sum int
		a, b, c    [3]int
	)
	fmt.Scanf("%d.%d.%d %d.%d.%d\n", &a[0], &a[1], &a[2], &b[0], &b[1], &b[2])
	sum = a[2] + b[2] + carry
	c[2] = sum % 29
	carry = sum / 29
	sum = a[1] + b[1] + carry
	c[1] = sum % 17
	carry = sum / 17
	c[0] = a[0] + b[0] + carry
	fmt.Printf("%d.%d.%d", c[0], c[1], c[2])
}
