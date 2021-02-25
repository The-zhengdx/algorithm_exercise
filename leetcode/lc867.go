package leetcode

func transpose(matrix [][]int) [][]int {
	r, c := len(matrix), len(matrix[0])
	res := make([][]int, c)
	for i := 0; i < c; i++ {
		temp := make([]int, r)
		for j := 0; j < r; j++ {
			temp[j] = matrix[j][i]
		}
		res[i] = temp
	}
	return res
}
