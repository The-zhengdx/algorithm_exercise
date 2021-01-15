package main

import "fmt"

func main() {
	colorInMars := [13]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C"}
	fmt.Print("#")
	var color int
	for n, _ := fmt.Scanf("%d", &color); n != 0; n, _ = fmt.Scanf("%d", &color) {
		fmt.Printf("%s%s", colorInMars[color/13], colorInMars[color%13])
	}
}
