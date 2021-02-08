package leetcode

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	cnt, last := 0, nums[0]
	for i, insert := 1, 1; i < l; i++ {
		if nums[i] != last {
			nums[insert] = nums[i]
			last = nums[i]
			insert++
		} else {
			cnt++
		}
	}
	return l - cnt
}
