package main

import (
)

//Basic implements Entity and will simply try to get from where it is to its goal
type Basic struct {
	X,Y int
	goalX, goalY int
}

func (b *Basic) Move(g *Grid) Movement {
	//Decide on where to move based on information from the grid
	//for now, if g.At(x,y) != null, that space is occupied
}

func (b *Basic) AtGoal() bool {
	return b.X == b.goalX && b.Y == b.goalY
}
