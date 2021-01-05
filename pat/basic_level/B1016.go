package main

import "fmt"

func main() {
	a, da, b, db := "", 0, "", 0
	fmt.Scanf("%s %d %s %d\n", &a, &da, &b, &db)

	pa, pb := calcP(a, da), calcP(b, db)
	fmt.Println(pa + pb)
}

func calcP(s string, n int) int {
	p, factor := 0, 1
	for i, l := 0, len(s); i < l; i++ {
		tmp := int(s[i] - '0')
		if n == tmp {
			p += tmp * factor
			factor *= 10
		}
	}
	return p
}
