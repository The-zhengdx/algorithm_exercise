package leetcode

func isMonotonic(A []int) bool {
	l := len(A)
	if l == 1 {
		return true
	}
	last, curr := 0, 0
	for i := 1; i < l; i++ {
		curr = A[i] - A[i-1]
		if curr != 0 {
			if (curr * last) < 0 {
				return false
			}
			last = curr
		}
	}
	return true
}
