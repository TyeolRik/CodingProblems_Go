// https://www.acmicpc.net/problem/10952

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		fmt.Fscanf(reader, "%d %d\n", &A, &B)
		if A == 0 && B == 0 {
			break
		} else {
			fmt.Fprintln(writer, A+B)
		}
	}
	writer.Flush()
}
