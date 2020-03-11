// https://www.acmicpc.net/problem/2562

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMax(arr []int) (max int, idx int) {
	max = arr[0]
	idx = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			idx = i
		}
	}
	return max, idx
}

func main() {
	var inputArray []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	inputArray = make([]int, 9)
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &inputArray[i])
	}
	max, idx := getMax(inputArray)
	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, idx+1)
	writer.Flush()
}
