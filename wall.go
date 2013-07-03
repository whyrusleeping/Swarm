package main

import (
	"image/color"
)

type Wall struct {
	pos Point
}

func (w Wall) Color() color.RGBA {
	return color.RGBA{0,0,0,0}
}

func (w Wall) Move(g *Grid) Point {
	return w.pos
}

func (w Wall) AtGoal() bool {
	return false
}

func (w Wall) GetPos() Point {
	return w.pos
}

func (w Wall) SetPos(p Point) {}

func (w Wall) GetGoal() Point {
	return w.pos
}
