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

func NewSimulation(g *Grid) *Simulation {
	s := new(Simulation)
	s.X = g.x
	s.Y = g.y
	s.grid = g
	s.w = NewWindow(s.X,s.Y)
	return s
}

func (s *Simulation) Run() {
	runtime.GOMAXPROCS(4)
	go s.w.Run()
	s.entities = append(s.entities, NewBasic(Point{0,0},Point{5,5}))
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
