package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	dist, btw := make([]int, n+1), 0
	for i := 1; i < n+1; i++ {
		if i < n {
			fmt.Scanf("%d", &btw)
		} else {
			fmt.Scanf("%d\n", &btw)
		}
		dist[i] = dist[i-1] + btw
	}

	var m, from, to, clockWiseDist int
	fmt.Scanf("%d\n", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &from, &to)
		if from > to {
			from, to = to, from
		}
		clockWiseDist = dist[to-1] - dist[from-1]
		if clockWiseDist < dist[n]-clockWiseDist {
			fmt.Println(clockWiseDist)
		} else {
			fmt.Println(dist[n] - clockWiseDist)
		}
	}
}
