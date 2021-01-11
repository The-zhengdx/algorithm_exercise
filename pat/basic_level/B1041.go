package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	id    string
	sitNo string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	stuMap := make([]*student, n)
	var inputStrs []string
	for i := 0; i < n; i++ {
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		stu := new(student)
		stu.id = inputStrs[0]
		stu.sitNo = inputStrs[2]
		tryNo, _ := strconv.Atoi(inputStrs[1])
		stuMap[tryNo-1] = stu
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	inputNums := make([]int, m)
	sc.Scan()
	inputStrs = strings.Split(sc.Text(), " ")
	for i := 0; i < m; i++ {
		inputNums[i], _ = strconv.Atoi(inputStrs[i])
	}
	for i := 0; i < m; i++ {
		fmt.Println(stuMap[inputNums[i]-1].id, stuMap[inputNums[i]-1].sitNo)
	}
}
