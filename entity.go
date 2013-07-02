package main

import (
	"image/color"
)

//the entity interface represents anything that can take up a space on the map
type Entity interface {
	Move(g *Grid) Point
	AtGoal() bool
	Color() color.RGBA
	GetPos() Point
	SetPos(p Point)
	GetGoal() Point
}
