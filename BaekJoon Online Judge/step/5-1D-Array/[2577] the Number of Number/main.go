// https://www.acmicpc.net/problem/2577

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numberArray [10]int
	var A, B, C int
	var result int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	fmt.Fscanln(reader, &C)

	result = A * B * C
	resultString := strconv.Itoa(result)

	for _, v := range resultString {
		numberArray[v-48]++
	}

	for _, i := range numberArray {
		fmt.Fprintln(writer, i)
	}
	writer.Flush()
}
