package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	var a, cnt [5]int
	sign, num := 1, 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num)
		switch {
		case num%5 == 0 && num%2 == 0:
			cnt[0]++
			a[0] += num
		case num%5 == 1:
			cnt[1]++
			a[1] += num * sign
			sign *= -1
		case num%5 == 2:
			cnt[2]++
			a[2]++
		case num%5 == 3:
			cnt[3]++
			a[3] += num
		case num%5 == 4:
			cnt[4]++
			if num > a[4] {
				a[4] = num
			}
		}
	}

	for i := 0; i < 5; i++ {
		if cnt[i] > 0 {
			if i == 3 {
				fmt.Printf("%.1f", float64(a[i])/float64(cnt[i]))
			} else {
				fmt.Print(a[i])
			}
		} else {
			fmt.Print("N")
		}
		if i != 4 {
			fmt.Print(" ")
		}
	}
}
