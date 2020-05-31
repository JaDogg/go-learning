package main

import "fmt"

func calc(x int) int {
	fmt.Println("Calc is called")
	return x + x
}

func main() {
	// calc is called immeditely
	defer fmt.Println(calc(1))
	defer fmt.Println("3")
	defer fmt.Println("4")

	// hellow happens next
	fmt.Println("hello")
	// defered statements run in reverse order (as a stack)
}

/* Result
Calc is called
hello
4
3
2

*/
