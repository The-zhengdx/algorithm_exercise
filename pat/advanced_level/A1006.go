package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		lockPerson, unlockPerson string
		earliestTime             = "23:59:59"
		latestTime               = "00:00:00"
		inputStrs                []string
	)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	for i := 0; i < m; i++ {
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		if inputStrs[1] < earliestTime {
			earliestTime = inputStrs[1]
			unlockPerson = inputStrs[0]
		}
		if inputStrs[2] > latestTime {
			latestTime = inputStrs[2]
			lockPerson = inputStrs[0]
		}
	}
	fmt.Println(unlockPerson, lockPerson)
}
