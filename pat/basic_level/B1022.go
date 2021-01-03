package main

import "fmt"

func main() {
	var a, b, d int
	fmt.Scanf("%d %d %d\n", &a, &b, &d)

	DSum := make([]int, 0, 30)
	for sum := a + b; ; {
		DSum = append(DSum, sum%d)
		sum /= d
		if sum == 0 {
			break
		}
	}

	for i := len(DSum) - 1; i >= 0; i-- {
		fmt.Print(DSum[i])
	}
}
