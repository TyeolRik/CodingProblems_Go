// https://www.acmicpc.net/problem/2739

package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 1; i < 10; i++ {
		fmt.Println(N, "*", i, "=", N*i)
	}
}
