package leetcode

func fairCandySwap(A, B []int) []int {
	sumA, sumB, mapB, avgSub := 0, 0, make(map[int]bool), 0
	for i, l := 0, len(A); i < l; i++ {
		sumA += A[i]
	}
	for i, l := 0, len(B); i < l; i++ {
		sumB += B[i]
		mapB[B[i]] = true
	}
	avgSub = (sumA - sumB) / 2
	for i, l := 0, len(A); i < l; i++ {
		if mapB[A[i]-avgSub] {
			return []int{A[i], A[i] - avgSub}
		}
	}
	return []int{0, 0}
}
