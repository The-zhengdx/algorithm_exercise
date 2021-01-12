package main

import "fmt"

func main() {
	res := []string{"W", "T", "L"}
	maxRes := make([]int, 3)
	var (
		w, t, l, odds float64
		product       = 1.0
	)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%f %f %f\n", &w, &t, &l)
		maxRes[i], odds = getMaxOdds(w, t, l)
		product *= odds
	}
	fmt.Printf("%s %s %s %.2f", res[maxRes[0]], res[maxRes[1]], res[maxRes[2]], (product*0.65-1)*2)
}

func getMaxOdds(w, t, l float64) (int, float64) {
	if w > t && w > l {
		return 0, w
	}
	if t > w && t > l {
		return 1, t
	}
	return 2, l
}
