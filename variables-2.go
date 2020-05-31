package main

import "fmt"

// Different type ones should be placed in it's own lines
// Unless you initialize them
// We are initializing these to respective values
var a, d, e int = 1, 2, 3
var b bool = false

// Below is not allowed in the package level,
// It is only allowed in the
// i := 12

func main() {
	// local `e` overrides `e` in global scope
	var c, f, g, e = "hello", true, 1, 2
	h := 10
	fmt.Println(a, b, c, d, e, f, g, h)
}
