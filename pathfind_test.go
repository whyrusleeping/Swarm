package main

import (
	"testing"
	"fmt"
)

func TestBasicPathfind(t *testing.T) {
	grid := NewGrid(8,8)
	basic := NewBasic(0,0,7,7)
	basic.visited = NewPointSet(10)
	queue := basic.PredictPath(grid,Point{0,0})
	for i, p := range queue.arr {
		fmt.Printf("%d : %d - %d\n", i,p.X,p.Y)
	}
}
