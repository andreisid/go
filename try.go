// https://www.youtube.com/watch?v=nSYFfWijl8U&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=1

package main

import (
	"fmt"
	_ "fmt"
	"runtime"
)

// add all elements of an array
func addTry(n []int) int {
	sum := 0
	for _, j := range n {
		sum += j
	}
	return sum
}

//variable assign test
func variableTry() int {
	a, b := 1, 2
	return a + b
}

func pointerTry() {
	x := 15
	fmt.Println(&x)
	fmt.Println(x)
	a := &x //a is a memory address
	fmt.Println(a)
	fmt.Println(*a) //value at that memory addres
	*a = 5          //value at the address changes
	fmt.Println(a)
	fmt.Println(x) //x changed
	*a = *a * *a   //like x=x*x
	fmt.Println(x)
}

func runtimeTry() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Compiler)
}

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	//fmt.Println("bla ", rand.Intn(9))
	//fmt.Println(addTry([]int{rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9), rand.Intn(9)}))
	//pointerTry()
	runtimeTry()
}