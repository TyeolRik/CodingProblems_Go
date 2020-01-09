// https://www.acmicpc.net/problem/2741

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &N)
	for i := 1; i <= N; i++ {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
