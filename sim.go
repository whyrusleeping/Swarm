package main

import (
	"runtime"
)

type Simulation struct {
	X,Y int
	g *Grid
	w *Window
}

func NewSimulation(scrx, scry int) *Simulation {
	s := new(Simulation)
	s.X = scrx
	s.Y = scry
	s.g = NewGrid(scrx,scry)
	s.w = NewWindow(scrx, scry)
	return s
}

func (s *Simulation) Run() {
	runtime.GOMAXPROCS(4)
	go func() {s.w.DrawChan <- s.g}()
	s.w.Run()
}
