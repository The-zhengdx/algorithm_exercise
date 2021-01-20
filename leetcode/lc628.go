package leetcode

import (
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	p1 := nums[l-1] * nums[l-2] * nums[l-3]
	p2 := nums[0] * nums[1] * nums[l-1]
	max := p1
	if p1 < p2 {
		max = p2
	}
	return max
}
