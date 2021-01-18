package main

import (
	"fmt"
	"sort"
)

func main() {
	var k int
	fmt.Scanf("%d\n", &k)
	nums, isCovered, m := make([]int, k), make([]bool, k), make(map[int]int)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Ints(nums)
	for i := 0; i < k; i++ {
		m[nums[i]] = i
	}
	for i := 0; i < k; i++ {
		if !isCovered[i] {
			num := nums[i]
			for num != 1 {
				if num%2 == 1 {
					num = 3*num + 1
				}
				num /= 2
				if index, ok := m[num]; ok {
					isCovered[index] = true
				}
			}
		}
	}
	cnt := 0
	for i := k - 1; i >= 0; i-- {
		if !isCovered[i] {
			if cnt > 0 {
				fmt.Print(" ")
			}
			cnt++
			fmt.Print(nums[i])
		}
	}
}
