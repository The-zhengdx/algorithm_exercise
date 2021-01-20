// 八皇后问题
package main

import (
	"fmt"
	"math"
)

func main() {
	generateP(0, 8, make([]int, 8), make([]bool, 8))
}

func generateP(index, n int, p []int, flag []bool) {
	if index == n {
		for i := 0; i < n; i++ {
			fmt.Print(p[i])
		}
		fmt.Println()
	}
	for i := 1; i <= n; i++ {
		if !flag[i-1] {
			ok := true
			for j := 0; j < index; j++ {
				if math.Abs(float64(index-j)) == math.Abs(float64(i-p[j])) {
					ok = false
					break
				}
			}
			if ok {
				p[index] = i
				flag[i-1] = true
				generateP(index+1, n, p, flag)
				flag[i-1] = false
			}
		}
	}
}
