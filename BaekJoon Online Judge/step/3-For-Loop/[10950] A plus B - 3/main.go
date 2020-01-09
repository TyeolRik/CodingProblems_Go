// https://www.acmicpc.net/problem/10950

package main

import "fmt"

func main() {
	var a int
	var tempA, tempB int
	fmt.Scanf("%d\n", &a)
	for i := 0; i < a; i++ {
		fmt.Scanf("%d %d\n", &tempA, &tempB)
		fmt.Println(tempA + tempB)
	}
}
