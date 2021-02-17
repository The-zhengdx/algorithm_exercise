package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	originR, originC := len(nums), len(nums[0])
	if originR*originC != r*c {
		return nums
	}
	reshapeNums, m, n := make([][]int, r), 0, 0
	for i := 0; i < r; i++ {
		reshapeNums[i] = make([]int, c)
	}
	for i := 0; i < originR; i++ {
		for j := 0; j < originC; j++ {
			reshapeNums[m][n] = nums[i][j]
			n++
			if n >= c {
				m++
				n = 0
			}
		}
	}
	return reshapeNums
}
