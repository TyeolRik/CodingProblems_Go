// https://www.acmicpc.net/problem/10951

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var A, B int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for {
		_, err := fmt.Fscanf(reader, "%d %d\n", &A, &B)
		if err == io.EOF {
			break
		} else {
			fmt.Fprintln(writer, A+B)
		}
	}
	writer.Flush()
}
