package main

type Entity interface {
	Move(g *Grid) Movement
	AtGoal() bool
}


