package main

import "fmt"

func main() {
	const CLK_TCK = 100
	var c1, c2 int
	fmt.Scanf("%d %d\n", &c1, &c2)
	time := int(float64(c2-c1)/CLK_TCK + 0.5)
	fmt.Printf("%02d:%02d:%02d", time/3600, time/60%60, time%60)
}
