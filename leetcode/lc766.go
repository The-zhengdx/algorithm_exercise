package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, 0; i < m-1; i++ {
		for r, c := i+1, j+1; r < m && c < n; r, c = r+1, c+1 {
			if matrix[r][c] != matrix[i][j] {
				return false
			}
		}
	}
	for i, j := 0, 1; j < n-1; j++ {
		for r, c := i+1, j+1; r < m && c < n; r, c = r+1, c+1 {
			if matrix[r][c] != matrix[i][j] {
				return false
			}
		}
	}
	return true
}
