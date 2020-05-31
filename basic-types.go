package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const MyPi = 3.14

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("constant's type = %T val = %v\n", MyPi, MyPi)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Printf("Type %T value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T value %v\n", z, z)
	// ----------------------------------------------
	x, y := 1, 2
	myFloat := float64(x)
	myUint := uint(y)
	fmt.Printf("Type %T value %v\n", myFloat, myFloat)
	fmt.Printf("Type %T value %v\n", myUint, myFloat)
	fmt.Printf("Type %T value %v\n", x, x)
	fmt.Printf("Type %T value %v\n", y, y)

}
