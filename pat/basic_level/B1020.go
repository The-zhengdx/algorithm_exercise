package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mooncake struct {
	cnt    float64
	amount float64
	price  float64
}

type cakeSlice []*mooncake

func (s cakeSlice) Len() int {
	return len(s)
}

func (s cakeSlice) Less(i, j int) bool {
	return s[i].price > s[j].price
}

func (s cakeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	n, d := 0, 0.0
	fmt.Scanf("%d %f\n", &n, &d)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	cntStr := sc.Text()
	sc.Scan()
	amountStr := sc.Text()
	cnts, amounts := strings.Split(cntStr, " "), strings.Split(amountStr, " ")
	s := make([]*mooncake, 0, n)
	for i := 0; i < n; i++ {
		cnt, _ := strconv.ParseFloat(cnts[i], 64)
		amount, _ := strconv.ParseFloat(amounts[i], 64)
		price := float64(amount) / float64(cnt)
		mc := new(mooncake)
		mc.cnt = cnt
		mc.amount = amount
		mc.price = price
		s = append(s, mc)
	}
	sort.Sort(cakeSlice(s))
	var max float64
	for i := 0; d > 0 && i < n; i++ {
		if s[i].cnt < d {
			max += float64(s[i].amount)
			d -= s[i].cnt
		} else {
			max += s[i].price * float64(d)
			d = 0
		}
	}
	fmt.Printf("%.2f", max)
}
