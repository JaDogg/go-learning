package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	// Make the matrix
	y := make([][]uint8, dy)
	for row := range y {
		y[row] = make([]uint8, dx)
	}

	// Fill up individual cell's
	// This will print a gradient with white lines
	for row := range y {
		for column := range y[row] {
			if column%10 == 0 {
				y[row][column] = 255
			} else {
				row_level := float64(row) / float64(dy)
				col_level := float64(column) / float64(dx)
				color := math.Abs(255.0 * (col_level - row_level))
				y[row][column] = uint8(color)
			}
		}
	}

	return y
}

func main() {
	pic.Show(Pic)
}
