package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, len(buf))
	sc.Scan()
	brokenKeys := sc.Text()
	sc.Scan()
	inputStr := sc.Text()
	brokenKeysMap := make(map[byte]bool)
	for i, l := 0, len(brokenKeys); i < l; i++ {
		brokenKeysMap[brokenKeys[i]] = true
	}
	for i, l := 0, len(inputStr); i < l; i++ {
		key := inputStr[i]
		if key >= 'a' && key <= 'z' {
			key -= 'a' - 'A'
		}
		if brokenKeysMap[key] || brokenKeysMap['+'] && inputStr[i] >= 'A' && inputStr[i] <= 'Z' {
			continue
		}
		fmt.Printf("%c", inputStr[i])
	}
	fmt.Println()
}
