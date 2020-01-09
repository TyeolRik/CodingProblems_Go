// https://www.acmicpc.net/problem/2753
package main

import "fmt"

func main() {
	var year int
	fmt.Scanf("%d", &year)
	if year%4 == 0 {
		if (year%100 != 0) || (year%400 == 0) {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	} else {
		fmt.Print("0")
	}
}
