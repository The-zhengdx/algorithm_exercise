package main

import "fmt"

func main() {
	var n, maxNum int
	fmt.Scanf("%d\n", &n)

	total := make([]int, n+1)
	for i := 0; i < n; i++ {
		var num, grade int
		fmt.Scanf("%d %d\n", &num, &grade)
		total[num] += grade
		if total[maxNum] < total[num] {
			maxNum = num
		}
	}

	fmt.Println(maxNum, total[maxNum])
}
