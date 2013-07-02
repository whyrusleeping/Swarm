package main

var ZeroPoint = Point{0,0}

type Point struct {
	X,Y int
}

func (p Point) Up() Point {
	p.Y--
	return p
}
func (p Point) Down() Point {
	p.Y++
	return p
}
func (p Point) Left() Point {
	p.X--
	return p
}
func (p Point) Right() Point {
	p.X++
	return p
}
func (p Point) Equals(op Point) bool {
	return p.X == op.X && p.Y == op.Y
}
