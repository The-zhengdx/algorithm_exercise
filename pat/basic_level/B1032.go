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

	n, maxNum, inputStrs := 0, 0, make([]string, 2)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	total := make([]int, n+1)
	for i := 0; i < n; i++ {
		var num, grade int
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		num, _ = strconv.Atoi(inputStrs[0])
		grade, _ = strconv.Atoi(inputStrs[1])
		total[num] += grade
		if total[maxNum] < total[num] {
			maxNum = num
		}
	}

	fmt.Println(maxNum, total[maxNum])
}
