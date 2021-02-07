package leetcode

func checkPossibility(nums []int) bool {
	l := len(nums)
	if l < 3 {
		return true
	}
	bpCnt, firstEnd, secondHead, bp := 0, -1, -1, -1
	for i := 1; i < l; i++ {
		if nums[i] < nums[i-1] {
			if bpCnt == 0 {
				bp = i
				firstEnd = i - 1
				secondHead = i + 1
				bpCnt++
			} else {
				return false
			}
		}
	}
	if firstEnd == -1 || firstEnd == 0 || secondHead == l || nums[firstEnd] <= nums[secondHead] || nums[firstEnd-1] <= nums[bp] {
		return true
	}
	return false
}
