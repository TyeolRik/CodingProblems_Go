// https://www.acmicpc.net/problem/2588
package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scanf("%d\n%d", &A, &B)
	fmt.Println(A * (B % 10))
	fmt.Println(A * (B / 10 % 10))
	fmt.Println(A * (B / 100))
	fmt.Println(A * B)
}
