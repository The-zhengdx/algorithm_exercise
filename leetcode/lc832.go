package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	for i, l := 0, len(A); i < l; i++ {
		change(A[i])
	}
	return A
}

func change(arr []int) {
	l := len(arr)
	for i, r := 0, l/2; i < r; i++ {
		if arr[i] == arr[l-1-i] {
			arr[i], arr[l-1-i] = arr[i]^1, arr[l-1-i]^1
		}
	}
	if l%2 == 1 {
		arr[l/2] = arr[l/2] ^ 1
	}
}
