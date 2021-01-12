package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var (
		n, needLine int
		char        string
		inputStrs   []string
	)
	inputStrs = strings.Split(sc.Text(), " ")
	n, _ = strconv.Atoi(inputStrs[0])
	char = inputStrs[1]
	needLine = int(math.Sqrt(float64((n + 1) / 2)))

	// 打印上半截
	for i := 0; i < needLine; i++ {
		// 打印空格
		printBlank(i)
		// 打印字符
		printChar(2*(needLine-i)-1, char)
	}
	// 打印下半截
	for i := needLine - 2; i >= 0; i-- {
		// 打印空格
		printBlank(i)
		// 打印字符
		printChar(2*(needLine-i)-1, char)
	}
	fmt.Println(n - 2*needLine*needLine + 1)
}

func printBlank(num int) {
	for i := 0; i < num; i++ {
		fmt.Print(" ")
	}
}

func printChar(num int, char string) {
	for i := 0; i < num; i++ {
		fmt.Print(char)
	}
	fmt.Println()
}
