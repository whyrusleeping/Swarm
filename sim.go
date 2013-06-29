package main

import (
	"runtime"
)

type Simulation struct {
	X,Y int
	grid *Grid
	w *Window
	entities []Entity
}

func NewSimulation(scrx, scry int) *Simulation {
	s := new(Simulation)
	s.X = scrx
	s.Y = scry
	s.grid = NewGrid(scrx,scry)
	s.w = NewWindow(scrx, scry)
	return s
}

func (s *Simulation) Run() {
	runtime.GOMAXPROCS(4)
	s.w.Run()
	for {
		for _,e := range s.entities {
			move := e.Move(s.grid)
			s.grid.Update(e,move)
		}

		//Send grid to be drawn, and wait for confirmation that it was
		s.w.DrawChan <- s.grid
		<-s.grid.lock
	}
}
