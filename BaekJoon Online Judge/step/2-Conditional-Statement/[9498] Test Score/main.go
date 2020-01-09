// https://www.acmicpc.net/problem/9498
package main

import "fmt"

func main() {
	var score int
	fmt.Scanf("%d", &score)
	if score >= 90 {
		fmt.Print("A")
	} else if score >= 80 {
		fmt.Print("B")
	} else if score >= 70 {
		fmt.Print("C")
	} else if score >= 60 {
		fmt.Print("D")
	} else {
		fmt.Print("F")
	}
}
