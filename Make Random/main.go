package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//var max int64
	//max = 1000000000 // 10억
	p := rand.Perm(1000000000)
	_ = p
	for idx, r := range p[:30000000] {
		fmt.Printf("%d", r)
		if idx != 29999999 {
			fmt.Printf(" ")
		}
	}
}
