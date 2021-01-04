package main

import "fmt"

func main() {
	words, word := make([]string, 0, 80), ""
	for {
		n, _ := fmt.Scanf("%s", &word)
		if n > 0 {
			words = append(words, word)
		} else {
			break
		}
	}

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i == 0 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}
