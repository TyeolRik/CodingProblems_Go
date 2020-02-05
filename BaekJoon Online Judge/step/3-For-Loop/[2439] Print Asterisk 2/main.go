// https://www.acmicpc.net/problem/2439

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
	for i := inputNumber; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			fmt.Fprintf(writer, " ")
		}
		for k := inputNumber - i; k >= 0; k-- {
			fmt.Fprintf(writer, "*")
		}
		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}
