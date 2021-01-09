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
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, len(buf))
	sc.Scan()

	inputStrs := strings.Split(sc.Text(), " ")
	l := len(inputStrs)
	inputNums := make([]int, l)
	for i := 0; i < l; i++ {
		inputNums[i], _ = strconv.Atoi(inputStrs[i])
	}

	dist := make([]int, l)
	for i := 1; i < l; i++ {
		dist[i] = dist[i-1] + inputNums[i]
	}

	var m, from, to, clockWiseDist int
	sc.Scan()
	m, _ = strconv.Atoi(sc.Text())
	for i := 0; i < m; i++ {
		sc.Scan()
		from, _ = strconv.Atoi(strings.Split(sc.Text(), " ")[0])
		to, _ = strconv.Atoi(strings.Split(sc.Text(), " ")[1])
		if from > to {
			from, to = to, from
		}
		clockWiseDist = dist[to-1] - dist[from-1]
		if clockWiseDist < dist[l-1]-clockWiseDist {
			fmt.Println(clockWiseDist)
		} else {
			fmt.Println(dist[l-1] - clockWiseDist)
		}

	}
}
