package main

import "fmt"

func main() {
	res := addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1)
	fmt.Println(res)
}

func addToArrayForm(A []int, K int) []int {
	kArr := make([]int, 0, 10)
	for K > 0 {
		kArr = append(kArr, K%10)
		K /= 10
	}
	carry, l1, l2, n1, n2, sum := 0, len(A), len(kArr), 0, 0, 0
	res := make([]int, 0, max(l1, l2)+1)
	for i, j := l1-1, 0; i >= 0 || j < l2; i, j = i-1, j+1 {
		if i >= 0 {
			n1 = A[i]
		} else {
			n1 = 0
		}
		if j < l2 {
			n2 = kArr[j]
		} else {
			n2 = 0
		}
		sum = n1 + n2 + carry
		carry = sum / 10
		res = append(res, sum%10)
	}
	if carry > 0 {
		res = append(res, carry)
	}
	reverse(res)
	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func reverse(a []int) {
	for i, l := 0, len(a); i < l/2; i++ {
		a[i], a[l-i-1] = a[l-i-1], a[i]
	}
}
