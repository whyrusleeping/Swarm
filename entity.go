package main

import (
	"image/color"
)

type Entity interface {
	Move(g *Grid) Movement
	AtGoal() bool
	Color() color.RGBA
}


