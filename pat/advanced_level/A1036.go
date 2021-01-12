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
	var (
		inputStrs                                    []string
		highestName, highestID, lowestName, lowestID string
		highest, lowest, grade                       = -1, 101, 0
	)
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		grade, _ = strconv.Atoi(inputStrs[3])
		if inputStrs[1] == "M" && grade < lowest {
			lowest = grade
			lowestName = inputStrs[0]
			lowestID = inputStrs[2]
		}
		if inputStrs[1] == "F" && grade > highest {
			highest = grade
			highestName = inputStrs[0]
			highestID = inputStrs[2]
		}
	}
	if highestName == "" {
		fmt.Println("Absent")
	} else {
		fmt.Println(highestName, highestID)
	}
	if lowestName == "" {
		fmt.Println("Absent")
	} else {
		fmt.Println(lowestName, lowestID)
	}
	if highestName == "" || lowestName == "" {
		fmt.Println("NA")
	} else {
		fmt.Println(highest - lowest)
	}
}
