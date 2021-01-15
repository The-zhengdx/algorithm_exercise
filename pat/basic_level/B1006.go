package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	var char = []string{"S", "B"}
	fmt.Scanf("%s\n", &n)
	nSlice := strings.Split(n, "")
	l := len(nSlice)
	for i := 0; i < l; i++ {
		num, _ := strconv.Atoi(nSlice[i])
		if i != l-1 {
			for j := 0; j < num; j++ {
				fmt.Print(char[l-i-2])
			}
		} else {
			for j := 1; j <= num; j++ {
				fmt.Print(j)
			}
		}
	}
}
