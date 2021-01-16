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
	betNums := strings.Split(sc.Text(), " ")[1:]
	var m [10001][2]int // 第一维记录出现次数，第二位记录第一次出现的下标
	for i, l := 0, len(betNums); i < l; i++ {
		num, _ := strconv.Atoi(betNums[i])
		m[num][0]++
		m[num][1] = i
	}
	res, flag := 0, true
	for i := 1; i < 10001; i++ {
		if m[i][0] == 1 {
			if flag {
				// 第一次遇到次数为1的数
				res = i
				flag = false
			} else {
				if m[i][1] < m[res][1] {
					// 遇到次数为1，且出现的位置比之前的靠前，更新res
					res = i
				}
			}
		}
	}
	if res == 0 {
		fmt.Print("None")
	} else {
		fmt.Print(res)
	}
}
