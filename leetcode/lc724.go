package leetcode

func pivotIndex(nums []int) int {
	left, right, l := 0, 0, len(nums)
	for i = 0; i < l; i++ {
		right += num[i]
	}
	for i = 0; i < l; i++ {
		right -= nums[i]
		if i > 0 {
			left += nums[i-1]
		}
		if right == left {
			return i
		}
	}
	return -1
}
