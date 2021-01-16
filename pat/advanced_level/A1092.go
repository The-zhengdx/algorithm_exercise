package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sale := sc.Text()
	sc.Scan()
	need := sc.Text()
	var m [129]int
	for i, l := 0, len(sale); i < l; i++ {
		m[sale[i]]++
	}
	for i, l := 0, len(need); i < l; i++ {
		m[need[i]]--
	}
	lackCnt, remainCnt := 0, 0
	for i, l := 0, len(m); i < l; i++ {
		if m[i] >= 0 {
			remainCnt += m[i]
		} else {
			lackCnt -= m[i]
		}
	}
	if lackCnt > 0 {
		fmt.Printf("No %d", lackCnt)
	} else {
		fmt.Printf("Yes %d", remainCnt)
	}
}
