package leetcode

func checkStraightline(coordinates [][]int) bool {
	l := len(coordinates)
	// 两点必在一条直线上
	if l == 2 {
		return true
	}
	// 检查是否在一条竖线上
	x, flag := coordinates[0][0], true
	for i := 1; i < l; i++ {
		if x != coordinates[i][0] {
			flag = false
		}
	}
	if flag {
		return true
	}
	// 检查斜率
	slope := (float64(coordinates[1][1]) - float64(coordinates[0][1])) / (float64(coordinates[1][0]) - float64(coordinates[0][0]))
	for i := 2; i < l; i++ {
		if slope != (float64(coordinates[i][1])-float64(coordinates[i-1][1]))/(float64(coordinates[i][0])-float64(coordinates[i-1][0])) {
			return false
		}
	}
	return true
}
