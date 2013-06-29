package main

import (
	"testing"
	"fmt"
)

func TestBasicPathfind(t *testing.T) {
	grid := NewGrid(20,20)
	basic := NewBasic(0,0,10,10)
	queue := basic.PredictPath(grid)

	for i, p := range queue.arr {
		fmt.Printf("%d : %d - %d\n", i,p.X,p.Y)
	}
}
