package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	M := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	var n int
	fmt.Scanf("%d\n", &n)

	idStr, idSlice, j, sum, flag := "", make([]string, 0), 0, 0, true
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &idStr)
		idSlice = strings.Split(idStr, "")
		sum = 0
		for j = 0; j < 17; j++ {
			num, err := strconv.Atoi(idSlice[j])
			if err != nil {
				fmt.Println(idStr)
				flag = false
				break
			}
			sum += weight[j] * num
		}
		if j == 17 {
			if idSlice[17] != M[sum%11] {
				fmt.Println(idStr)
				flag = false
			}
		}
	}

	if flag {
		fmt.Println("All passed")
	}
}
