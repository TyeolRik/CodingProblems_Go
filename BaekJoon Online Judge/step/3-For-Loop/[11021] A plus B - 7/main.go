// https://www.acmicpc.net/problem/11021

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputNumber, tempA, tempB int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &inputNumber)
	for i := 1; i <= inputNumber; i++ {
		fmt.Fscanln(reader, &tempA, &tempB)
		fmt.Fprintf(writer, "Case #%d: %d\n", i, tempA+tempB)
	}
	writer.Flush()
}
