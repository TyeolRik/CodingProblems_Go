// https://www.acmicpc.net/problem/15552

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputLength, tempA, tempB int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &inputLength)
	for i := 0; i < inputLength; i++ {
		fmt.Fscanln(reader, &tempA, &tempB)
		fmt.Fprintln(writer, tempA+tempB)
	}
	writer.Flush()
}
