// https://www.acmicpc.net/problem/10871

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arrayLength, standard int
	var inputArray []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader, "%d %d\n", &arrayLength, &standard)
	writer.Flush()
	inputArray = make([]int, arrayLength)
	for i := 0; i < arrayLength; i++ {
		fmt.Fscan(reader, &inputArray[i])
		if inputArray[i] < standard {
			fmt.Fprintf(writer, "%d ", inputArray[i])
		}
	}

	writer.Flush()
}
