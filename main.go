package main

import (
	"runtime/pprof"
	"os"
)

func main() {
	fi, _ := os.Create("pro.out")
	pprof.StartCPUProfile(fi)
	defer pprof.StopCPUProfile()
	// g := NewGrid(400,400)
	g,err := LoadGridFromBitmap("map.bmp")
	if err != nil {
		panic(err)
	}
	s := NewSimulation(g)
	s.Run()
}
