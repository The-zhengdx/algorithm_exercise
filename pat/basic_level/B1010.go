package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputStrs := strings.Split(sc.Text(), " ")
	l := len(inputStrs)
	inputNums := make([]int, l)
	for i := 0; i < l; i++ {
		inputNums[i], _ = strconv.Atoi(inputStrs[i])
	}

	for i := 0; i < l; i += 2 {
		inputNums[i] = inputNums[i] * inputNums[i+1]
		inputNums[i+1]--
	}

	flag := false
	for i := 0; i < l; i += 2 {
		if inputNums[i] != 0 {
			if flag {
				fmt.Print("_")
			}
			flag = true
			fmt.Printf("%d %d", inputNums[i], inputNums[i+1])
		}
	}
	if !flag {
		fmt.Print("0 0")
	}
}
