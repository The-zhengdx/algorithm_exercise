package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	l := len(str)
	n := (l + 2) / 3
	// 打印上半截
	for i := 0; i < n-1; i++ {
		fmt.Printf("%c", str[i])
		for j := 0; j < l-2*n; j++ {
			fmt.Print(" ")
		}
		fmt.Printf("%c\n", str[l-i-1])
	}
	// 打印底行
	for i := n - 1; i < l-n+1; i++ {
		fmt.Printf("%c", str[i])
	}
}
