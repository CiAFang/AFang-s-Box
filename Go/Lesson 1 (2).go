package main

import (
	"fmt"
	"math"
)

func main() {
	var rf int
	var rt float64
	fmt.Scanln(&rf)
	rt = float64(rf)
	var area float64
	area = rt * rt * math.Pi
	fmt.Println(area)
}
