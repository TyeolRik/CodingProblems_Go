// https://www.acmicpc.net/problem/1546

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMax(arr []float64) float64 {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var N int
	var originalScore []float64
	fmt.Fscanln(reader, &N)
	originalScore = make([]float64, N)
	newScore := make([]float64, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%f", &originalScore[i])
	}
	max := getMax(originalScore)
	var average float64 = 0
	for i := 0; i < N; i++ {
		newScore[i] = originalScore[i] / max * 100
		average = average + newScore[i]
	}
	fmt.Fprintln(writer, average/float64(N))
	writer.Flush()
}
