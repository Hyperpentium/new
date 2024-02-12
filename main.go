package main

import (
	"fmt"
)

var c, d int = 1, 2

func main() {
	fmt.Println(abs(float64(c), float64(d)))
}

func abs(x, y float64) (float64, float64) {
	var a, b float64
	a = x + y
	b = x - y
	return a, b

}
