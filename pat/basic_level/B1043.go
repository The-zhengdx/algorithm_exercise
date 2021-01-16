package main

import "fmt"

func main() {
	str, m, cnt := "", [6]string{"P", "A", "T", "e", "s", "t"}, [6]int{0}
	fmt.Scanf("%s\n", &str)
	for i, l := 0, len(str); i < l; i++ {
		switch str[i] {
		case 'P':
			cnt[0]++
		case 'A':
			cnt[1]++
		case 'T':
			cnt[2]++
		case 'e':
			cnt[3]++
		case 's':
			cnt[4]++
		case 't':
			cnt[5]++
		}
	}
	flag := true // 标记是否还有"PATest"的字符
	for flag {
		flag = false
		for i := 0; i < 6; i++ {
			if cnt[i] > 0 {
				fmt.Print(m[i])
				cnt[i]--
				flag = true
			}
		}
	}
}
