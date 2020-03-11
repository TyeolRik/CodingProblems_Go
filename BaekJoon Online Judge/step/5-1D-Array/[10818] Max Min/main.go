// https://www.acmicpc.net/problem/10818

package main

import (
	"bufio"
	"fmt"
	"os"
)

func minMax(arr []int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}

func main() {
	var arrayLength int
	var inputArray []int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader, "%d\n", &arrayLength)
	inputArray = make([]int, arrayLength)
	for i := 0; i < arrayLength; i++ {
		fmt.Fscanf(reader, "%d", &inputArray[i])
	}
	fmt.Println(minMax(inputArray))
	writer.Flush()
}
