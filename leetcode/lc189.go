package leetcode

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k > 0 {
		reverse(nums)
		reverse(nums[:k])
		reverse(nums[k:])
	}
}

func reverse(s []int) {
	for i, l := 0, len(s); i < l/2; i++ {
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
}
