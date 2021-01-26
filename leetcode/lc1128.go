package leetcode

func numEquivDominoPairs(dominoes [][]int) int {
	m, num := [100]int{}, 0
	for i, l := 0, len(dominoes); i < l; i++ {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}
		idx := dominoes[i][0]*10 + dominoes[i][1]
		num += m[idx]
		m[idx]++
	}
	return num
}
