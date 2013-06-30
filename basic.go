package main

import (
	"fmt"
	"image/color"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	NONE
)

//Basic implements Entity and will simply try to get from where it is to its goal
type Basic struct {
	X,Y int
	goalX, goalY int
	path *PathQueue
	visited *PointSet
	last_d int
	pred_at_goal bool
}

func NewBasic (X,Y int, goalX, goalY int) *Basic {
	b := new(Basic)
	b.X = X
	b.Y = Y
	b.goalX = goalX
	b.goalY = goalY
	b.last_d = NONE
	b.pred_at_goal = false

	return b
}

//Decide on where to move based on information from the grid
//for now, if g.At(x,y) != null, that space is occupied

//Here is where we are going to do cool pathfinding stuff, or maybe in
//another function that gets passed in? 
func (b *Basic) Move(g *Grid) Point {
	if b.path == nil {
		b.visited = NewPointSet(g.Width())
		b.path = b.PredictPath(g,Point{b.X,b.Y})
	}

	return b.path.PopBack()
}

func (b *Basic) PredictPath(g *Grid, pos Point) *PathQueue {
	var direc *PathQueue
	var temp Point
	b.visited.Add(pos)

	if b.AtPredictGoal(pos) {
		b.pred_at_goal = true
		direc = NewPathQueue()
	}
	temp = pos.Up()
	if g.InBounds(temp) && b.last_d != UP && g.At(temp) == nil && !b.pred_at_goal && !b.visited.Find(temp) {
		b.last_d = UP
		direc = b.PredictPath(g,temp)
	}
	temp = pos.Down()
	if g.InBounds(temp) && b.last_d != DOWN && g.At(temp) == nil && !b.pred_at_goal && !b.visited.Find(temp) {
		b.last_d = DOWN
		direc = b.PredictPath(g,temp)
	}
	temp = pos.Left()
	if g.InBounds(temp) && b.last_d != LEFT && g.At(temp) == nil && !b.pred_at_goal && !b.visited.Find(temp) {
		b.last_d = LEFT
		direc = b.PredictPath(g,temp)
	}
	temp = pos.Right()
	if g.InBounds(temp) && b.last_d != RIGHT && g.At(temp) == nil && !b.pred_at_goal && !b.visited.Find(temp) {
		b.last_d = RIGHT
		direc = b.PredictPath(g,temp)
	}
	if !b.AtDeadEnd(g,pos) {
		if direc != nil {
			direc.PushBack(pos)
		}
	}
	return direc
}

func (b *Basic) AtGoal() bool {
	fmt.Println("GOOOOOOOOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAL!")
	return b.X == b.goalX && b.Y == b.goalY
}

func (b *Basic) AtPredictGoal(pos Point) bool {
	return pos.X == b.goalX && pos.Y == b.goalY
}

func (b *Basic) AtDeadEnd(g *Grid, pos Point) bool {
	count := 0

	if b.last_d != UP && g.At(pos.Up()) != nil {
		count++
	}
	if b.last_d != DOWN && g.At(pos.Down()) != nil {
		count++
	}
	if b.last_d != LEFT && g.At(pos.Left()) != nil {
		count++
	}
	if b.last_d != RIGHT && g.At(pos.Right()) != nil {
		count++
	}

	return count == 3
}

func (b *Basic) Color() color.RGBA {
	return color.RGBA{255,255,255,128}
}

func (b *Basic) GetX() int {
	return b.X
}

func (b *Basic) GetY() int {
	return b.Y
}

func (b *Basic) SetPos(p Point) {
	b.X = p.X
	b.Y = p.Y
}

func (b *Basic) GetGoal() Point {
	return Point{b.goalX,b.goalY}
}
