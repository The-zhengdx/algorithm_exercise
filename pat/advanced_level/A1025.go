package main

import (
	"fmt"
	"sort"
)

type testee struct {
	id      string
	loc     int
	locRank int
	grade   int
}

type testeeSlice []*testee

func (s testeeSlice) Len() int {
	return len(s)
}

func (s testeeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s testeeSlice) Less(i, j int) bool {
	if s[i].grade != s[j].grade {
		return s[i].grade > s[j].grade
	}
	return s[i].id < s[j].id
}

func main() {
	testees := make(testeeSlice, 0, 30000)
	n := 0
	fmt.Scanf("%d\n", &n)
	for i := 1; i <= n; i++ {
		k := 0
		fmt.Scanf("%d\n", &k)
		tempSlice := make(testeeSlice, 0, k)
		// 开始读入k个考生信息
		for j := 0; j < k; j++ {
			temp := &testee{loc: i}
			fmt.Scanf("%s %d\n", &temp.id, &temp.grade)
			tempSlice = append(tempSlice, temp)
		}
		// 考场内排序
		sort.Sort(tempSlice)
		// 计算考场内排名
		if k > 0 {
			tempSlice[0].locRank = 1
		}
		for m := 1; m < k; m++ {
			if tempSlice[m].grade != tempSlice[m-1].grade {
				tempSlice[m].locRank = m + 1
			} else {
				tempSlice[m].locRank = tempSlice[m-1].locRank
			}
		}
		// 加入总数组
		testees = append(testees, tempSlice...)
	}
	// 总数组排名
	sort.Sort(testees)

	l, finalRank := len(testees), 1
	fmt.Println(len(testees))
	if l > 0 {
		fmt.Println(testees[0].id, 1, testees[0].loc, testees[0].locRank)
	}
	for i := 1; i < l; i++ {
		if testees[i].grade != testees[i-1].grade {
			finalRank = i + 1
		}
		fmt.Println(testees[i].id, finalRank, testees[i].loc, testees[i].locRank)
	}
}
