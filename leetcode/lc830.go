package leetcode

func largeGroupPositions(s string) [][]int {
	l := len(s)
	res := make([][]int, 0, l/3)
	cnt, start := 1, 0
	for i := 1; i < l; i++ {
		if s[i] == s[start] {
			cnt++
		} else {
			if cnt >= 3 {
				res = append(res, []int{start, i - 1})
			}
			cnt, start = 1, i
		}
	}
	return res
}
