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
	p := inputStrs[0]
	a := inputStrs[1]
	pStrs := strings.Split(p, ".")
	aStrs := strings.Split(a, ".")
	pGalleon, _ := strconv.Atoi(pStrs[0])
	pSickle, _ := strconv.Atoi(pStrs[1])
	pKnut, _ := strconv.Atoi(pStrs[2])
	pKnut += pGalleon*17*29 + pSickle*29
	aGalleon, _ := strconv.Atoi(aStrs[0])
	aSickle, _ := strconv.Atoi(aStrs[1])
	aKnut, _ := strconv.Atoi(aStrs[2])
	aKnut += aGalleon*17*29 + aSickle*29
	remainKnut := aKnut - pKnut
	flag := false // 是否为负
	if remainKnut < 0 {
		remainKnut = -remainKnut
		flag = true
	}
	res := make([]string, 3)
	res[2] = strconv.Itoa(remainKnut % 29)
	res[1] = strconv.Itoa(remainKnut / 29 % 17)
	res[0] = strconv.Itoa(remainKnut / 29 / 17)
	if flag {
		fmt.Print("-")
	}
	fmt.Println(strings.Join(res, "."))
}
