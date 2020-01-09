// https://www.acmicpc.net/problem/10430
package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	fmt.Println((A + B) % C)
	fmt.Println((A%C + B%C) % C)
	fmt.Println((A * B) % C)
	fmt.Println((A % C * B % C) % C)
}
