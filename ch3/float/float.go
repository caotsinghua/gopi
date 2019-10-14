package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 1 << 24
	fmt.Println(f == f+1)
	math.IsNaN(1)
}
