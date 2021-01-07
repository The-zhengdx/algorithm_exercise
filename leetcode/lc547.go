package leetcode

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	provinceNum, flag := 0, make([]bool, n)

	for i := 0; i < n; i++ {
		if !flag[i] {
			provinceNum++
			dfsByRecursion(isConnected, flag, i, n)
			// bfsByQueue(isConnected, flag, i, n)
		}
	}

	return provinceNum
}

// 广度优先搜索
func bfsByQueue(isConnected [][]int, flag []bool, v int, n int) {
	queue := make([]int, 0, n)
	queue = append(queue, v)
	flag[v] = true
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for j := 0; j < n; j++ {
			if !flag[j] && isConnected[i][j] == 1 {
				queue = append(queue, j)
				flag[j] = true
			}
		}
	}
}

// 深度优先搜索递归写法
func dfsByRecursion(isConnected [][]int, flag []bool, v int, n int) {
	flag[v] = true
	for j := 0; j < n; j++ {
		if !flag[j] && isConnected[v][j] == 1 {
			dfsByRecursion(isConnected, flag, j, n)
		}
	}
}
