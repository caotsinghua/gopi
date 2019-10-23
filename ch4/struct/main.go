package main

import (
	"fmt"
	"structmod/mianzhi"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) Employee {
	return Employee{ID: id}
}
func main() {
	// var dilbert Employee
	// dilbert.Salary -= 100
	// fmt.Println(dilbert)
	// position := &dilbert.Position
	// *position = "Senior " + *position
	// fmt.Println(dilbert.Position)
	// var epleOfMon *Employee = &dilbert
	// epleOfMon.Position += " (proactive team player)"
	// fmt.Println(dilbert)
	// em := EmployeeByID(12)
	// // EmployeeByID(12).ID=12;
	// fmt.Println(em)
	// em.ID = 120
	// fmt.Println(em.ID)
	// arr := []int{5, 6, 4, 5, 3, 1, 2, 3, 66, 1}
	// arr =Sort1(arr)
	// fmt.Println(arr)
	// p:= mianzhi.Point{1,2}
	// fmt.Println(mianzhi.Scale(p))
	// testEq()
	// tmap()
	insertAndHideName()
}

type tree struct {
	value       int
	left, right *tree
}

func Sort1(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	r := appendValues(values[:0], root)
	return r
}
func appendValues(values []int, t *tree) []int {
	// fmt.Println(t)
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {

	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// func Bonus(e *Employee, percent int)int {
// 	return e.Salary*percent/100
// }

func testEq() {
	p1 := mianzhi.Point{X: 1, Y: 2}
	p2 := mianzhi.Point{X: 1, Y: 2}
	fmt.Println(p1 == p2)
}

// 可比较的结构体也可以做key
func tmap() {
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}

func insertAndHideName() {
	// // 圆
	// type Circle struct{
	// 	X,Y,Radius int
	// }
	// // 椭圆
	// type Wheel struct{
	// 	X,Y,Radius,Spokes int
	// }
	// type Point struct {
	// 	X, Y int
	// }
	// type Circle struct {
	// 	Center Point
	// 	Radius int
	// }
	// type Wheel struct {
	// 	Circle Circle
	// 	Spokes int
	// }
	// var w Wheel
	// w.Circle.Center.X = 8
	// w.Circle.Center.Y = 8

	type Point struct {
		X, Y int
	}
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	var w Wheel
	w.Circle.X = 12
	w.X = 13
	w.Y = 9
	w.Radius = 5
	w.Spokes = 12
	fmt.Println(w)
	w = Wheel{Circle: Circle{Point: Point{X: 1, Y: 2}}}
	fmt.Printf("%v %#v \n", w, w)

}
