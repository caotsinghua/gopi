package main

import (
	"fmt"
	"math"
)

func main() {
	// t2()
	t3()
}

func t1() {
	var f float32 = 1 << 24
	fmt.Println(f == f+1) // true 精度问题
	math.IsNaN(1)
}
func t2() {
	var f float64 = 1 << 24
	fmt.Println(f == f+1)
}

func t3() {
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	fmt.Printf("d:%d ,g:%[1]g ,e:%[1]e,f:%[1]f\n", Avogadro)
	const f = 1.23e1
	fmt.Printf("%g %[1]e\n", f)
}

func t4() {
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

func t5(){

}
