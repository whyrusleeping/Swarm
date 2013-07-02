package main

import (
	"runtime/pprof"
	"os"
)

func main() {
	fi, _ := os.Create("pro.out")
	pprof.StartCPUProfile(fi)
	defer pprof.StopCPUProfile()
	s := NewSimulation(400,400)
	s.Run()
}
