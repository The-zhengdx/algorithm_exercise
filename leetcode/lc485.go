package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	cnt, max := 0, 0
	for i, l := 0, len(nums); i < l; i++ {
		if nums[i] == 1 {
			cnt++
		} else {
			if cnt > max {
				max = cnt
			}
			cnt = 0
		}
	}
	if cnt > max {
		return cnt
	}
	return max
}
