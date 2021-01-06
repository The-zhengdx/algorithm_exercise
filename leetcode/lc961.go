package leetcode

func repeatedNTimes(A []int) int {
	flag := make(map[int]bool)
	for a := range A {
		if flag[a] {
			return a
		}
		flag[a] = true
	}
	return 0
}
