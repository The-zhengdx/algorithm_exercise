package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		py  = []string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}
		sum int
		c   byte
	)
	for fmt.Scanf("%c", &c); c != '\n'; fmt.Scanf("%c", &c) {
		sum += int(c - '0')
	}
	sumStr := strconv.Itoa(sum)
	for i, l := 0, len(sumStr); i < l; i++ {
		fmt.Print(py[sumStr[i]-'0'])
		if i != l-1 {
			fmt.Print(" ")
		}
	}
}
