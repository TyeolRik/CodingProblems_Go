// https://www.acmicpc.net/problem/2884
package main

import "fmt"

func main() {
	var H, M int
	fmt.Scanf("%d %d", &H, &M)
	M = M - 45
	if M < 0 {
		H = H - 1
		M = M + 60
		if H < 0 {
			H = H + 24
		}
	}
	fmt.Printf("%d %d", H, M)
}
