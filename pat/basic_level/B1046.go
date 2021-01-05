package main

import "fmt"

func main() {
	var n, aMouth, aHand, bMouth, bHand, aDreak, bDreak int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d\n", &aMouth, &aHand, &bMouth, &bHand)
		if aHand != bHand {
			sum := aMouth + bMouth
			if aHand == sum {
				bDreak++
			}
			if bHand == sum {
				aDreak++
			}
		}
	}
	fmt.Println(aDreak, bDreak)
}
