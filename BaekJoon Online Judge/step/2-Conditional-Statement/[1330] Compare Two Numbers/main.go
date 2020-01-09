// https://www.acmicpc.net/problem/1330
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	if a > b {
		fmt.Print(">")
	} else if a < b {
		fmt.Print("<")
	} else {
		fmt.Print("==")
	}
}
