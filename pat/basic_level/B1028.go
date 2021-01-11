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
	n, _ := strconv.Atoi(sc.Text())
	const (
		floor = "1814/09/06"
		ceil  = "2014/09/06"
	)
	var (
		inputStrs                []string
		cnt                      int
		oldest                   = ceil
		youngest                 = floor
		oldestName, youngestName string
	)
	for i := 0; i < n; i++ {
		sc.Scan()
		inputStrs = strings.Split(sc.Text(), " ")
		if inputStrs[1] >= floor && inputStrs[1] <= ceil {
			cnt++
			if inputStrs[1] < oldest {
				oldest = inputStrs[1]
				oldestName = inputStrs[0]
			}
			if inputStrs[1] > youngest {
				youngest = inputStrs[1]
				youngestName = inputStrs[0]
			}
		}
	}
	if cnt > 0 {
		fmt.Println(cnt, oldestName, youngestName)
	} else {
		fmt.Println(0)
	}
}
