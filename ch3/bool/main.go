package main

import (
	"fmt"
)

func main() {
	t1()
}
func t1() {

	fmt.Println(btoi(true))
	fmt.Println(itob(7))
}
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func itob(i int) bool {
	return i != 0
}
