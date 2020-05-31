package main

import "fmt"

func calc(x int, y int) (a int, b int) {
    // in below line `temp` is not defined, so we need `:=`
    temp := x + y // 1 + 2 -> 3
    temp += 1     // 4
    // in below line `a` is defined since it's named return
    a = temp + y  // 6
    b = temp * y  // 8 
    // sum += 1 -> this is also allowed... 
    return
}

func main() {
    fmt.Println(calc(1, 2))
}
