// https://www.acmicpc.net/problem/3052

package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkInArray(arr []int, data int) bool {
	for _, value := range arr {
		if value == data {
			return true
		}
	}
	return false
}

func makeUnique(arr []int) []int {
	newArray := make([]int, 0)
	for _, value := range arr {
		if !checkInArray(newArray, value) {
			newArray = append(newArray, value)
		}
	}
	return newArray
}

func main() {
	inputArray := make([]int, 10)
	remainders := make([]int, 10)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for idx := range inputArray {
		fmt.Fscanln(reader, &inputArray[idx])
		remainders[idx] = inputArray[idx] % 42
	}
	uniqueRemainders := makeUnique(remainders)
	fmt.Fprintln(writer, len(uniqueRemainders))

	writer.Flush()
}
