// https://www.acmicpc.net/problem/8393

package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	temp := 0
	for i := 0; i <= a; i++ {
		temp = temp + i
	}
	fmt.Print(temp)
}
