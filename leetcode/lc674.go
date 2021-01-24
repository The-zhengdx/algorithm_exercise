package leetcode

func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return 1
	}
	max, currLen := 1, 1
	for i, j := 0, 1; j < l; j++ {
		if nums[j] <= nums[j-1] || j == l-1 {
			if nums[j] <= nums[j-1] {
				currLen = j - i
			} else {
				currLen = j - i + 1
			}
			if currLen > max {
				max = currLen
			}
			i = j
		}
	}
	return max
}
