package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const numOfCards = 54
	const numOfShapes = 5

	// 初始化卡片数组
	var cards [2][numOfCards]string
	shapes := [numOfShapes]string{"S", "H", "C", "D", "J"}
	for i := 0; i < numOfCards; i++ {
		cards[0][i] = shapes[i/13] + strconv.Itoa(i%13+1)
	}

	// 读取洗牌次数及指令数组
	k, order := 0, [numOfCards]int{}
	fmt.Scanf("%d\n", &k)
	for i := 0; i < numOfCards; i++ {
		fmt.Scanf("%d", &order[i])
	}

	// 洗牌
	for i := 0; i < k; i++ {
		from, to := i%k, (i+1)%k
		for j := 0; j < numOfCards; j++ {
			cards[to][order[j]-1] = cards[from][j]
		}
	}

	fmt.Println(strings.Join(cards[k%2][:], " "))
}
