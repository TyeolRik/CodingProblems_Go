// https://www.acmicpc.net/problem/10817
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	var ints = []int{a, b, c}
	sort.Sort(sort.IntSlice(ints))
	fmt.Println(ints[1])
}
