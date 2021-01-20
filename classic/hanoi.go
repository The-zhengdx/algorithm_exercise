// 汉诺塔
package main

import "fmt"

func main() {
	hanoi("A", "B", "C", 3)
}

func hanoi(a, b, c string, n int) {
	if n > 1 {
		hanoi(a, c, b, n-1)
		fmt.Println(a + "->" + c)
		hanoi(b, a, c, n-1)
	} else {
		fmt.Println(a + "->" + c)
	}
}
