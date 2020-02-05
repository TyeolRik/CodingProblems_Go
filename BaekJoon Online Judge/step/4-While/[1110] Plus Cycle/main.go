// https://www.acmicpc.net/problem/10951

package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNewNumber(num int) int {
	var ten int = num / 10
	var one int = num % 10
	var plus int = ten + one
	var newNumber int = one*10 + plus%10
	return newNumber
}

func main() {
	var inputNumber, time, temp int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader, "%d", &inputNumber)
	time = 0
	temp = inputNumber
	for {
		time++
		temp = getNewNumber(temp)
		if inputNumber == temp {
			fmt.Fprintf(writer, "%d", time)
			break
		}
	}
	writer.Flush()
}
