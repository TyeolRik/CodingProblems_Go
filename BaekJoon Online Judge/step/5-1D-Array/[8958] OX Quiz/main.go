// https://www.acmicpc.net/problem/8958

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var theNumberOfTestCase int
	var oxResult []string

	fmt.Fscanln(reader, &theNumberOfTestCase)
	oxResult = make([]string, theNumberOfTestCase)

	for i := 0; i < theNumberOfTestCase; i++ {
		fmt.Fscanln(reader, &oxResult[i])
	}

	writer.Flush()
}
