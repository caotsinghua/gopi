package mianzhi

type Point struct {
	X, Y int
}

// p:=Point{
// 	1,2
// }
func Scale(p Point, factor int) Point {
	return Point{X: p.X * factor, Y: p.Y * factor}
}
