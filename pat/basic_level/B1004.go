package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	highestName, highestNo, highestGrade := "", "", 0
	lowestName, lowestNo, lowestGrade := "", "", 101
	var (
		inputStrs []string
		grade     int
	)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		grade, _ = strconv.Atoi(inputStrs[2])
		if grade > highestGrade {
			highestGrade = grade
			highestName = inputStrs[0]
			highestNo = inputStrs[1]
		}
		if grade < lowestGrade {
			lowestGrade = grade
			lowestName = inputStrs[0]
			lowestNo = inputStrs[1]
		}
	}
	fmt.Println(highestName, highestNo)
	fmt.Println(lowestName, lowestNo)
}
