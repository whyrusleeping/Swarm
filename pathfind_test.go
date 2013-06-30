package main

import (
	"os"
	"testing"
	"fmt"
	"runtime/pprof"
)

func TestBasicPathfind(t *testing.T) {
	fi, err := os.Create("pro.out")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(fi)
	grid := NewGrid(20,20)
	basic := NewBasic(0,0,7,8)
	basic.visited = NewPointSet(100)
	//queue := basic.PredictPath(grid,Point{0,0})
	queue := BreadthFirst(grid, basic)
	pprof.StopCPUProfile()
	fi.Close()
	for i, p := range queue.arr {
		fmt.Printf("%d : %d - %d\n", i,p.X,p.Y)
	}
}
