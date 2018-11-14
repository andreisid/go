// https://www.youtube.com/watch?v=nSYFfWijl8U&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=1

package main

import (
	"fmt"
	_ "fmt"
	"math/rand"
	"time"
)

// add all elements of an array
func add(n []int) int {
	sum := 0
	for _, j := range n {
		sum += j
	}
	return sum
}

//
func test() int {
	var a, b int = 1, 2
	return a + b
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("bla ", rand.Intn(9))
	fmt.Println(add([]int{rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9)}))
}
