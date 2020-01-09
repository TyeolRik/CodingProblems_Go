// https://www.acmicpc.net/problem/2438

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputNumber int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &inputNumber)
	for i := 1; i <= inputNumber; i++ {
		for j := 0; j < i; j++ {
			fmt.Fprintf(writer, "*")
		}
		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}
