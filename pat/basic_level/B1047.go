package main

import "fmt"

func main() {
	var n, teamNo, personNo, grade int
	var totalGrade [10001]int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d-%d %d\n", &teamNo, &personNo, &grade)
		totalGrade[teamNo] += grade
	}
	max := 0
	for i := 0; i < 10001; i++ {
		if totalGrade[i] > totalGrade[max] {
			max = i
		}
	}
	fmt.Printf("%d %d", max, totalGrade[max])
}
