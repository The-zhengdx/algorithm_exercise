package main

import (
	"fmt"
	"math"
)

func main() {
	n, c := 0, ""
	fmt.Scanf("%d %s\n", &n, &c)

	for i, l := 0, int(math.Floor((float64(n)+1)/2)); i < l; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || j == n-1 {
				// 第一列或最后一列输出字符
				fmt.Print(c)
			} else {
				// 中间列
				if i == 0 || i == l-1 {
					// 第一行或最后一行输出字符
					fmt.Print(c)
				} else {
					// 中间行输出空格
					fmt.Print(" ")
				}
			}
			if j == n-1 {
				fmt.Println()
			}
		}
	}
}
