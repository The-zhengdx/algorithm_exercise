package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := strings.ToLower(sc.Text())
	var m [128]int
	for i, l := 0, len(str); i < l; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			m[str[i]]++
		}
	}
	max := int('a')
	for i := int('a'); i < int('z'); i++ {
		if m[i] > m[max] {
			max = i
		}
	}
	fmt.Printf("%c %d", byte(max), m[max])
}
