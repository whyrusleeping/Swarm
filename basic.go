package main

import (
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
	Pos Point
	Goal Point
	path *PathQueue
	visited *PointSet
	last_d int
	pred_at_goal bool
	SearchFunc func(g *Grid, e Entity) *PathQueue
}

func NewBasic (pos Point, goal Point) *Basic {
	b := new(Basic)
	b.Pos = pos
	b.Goal = goal
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
		b.path = b.SearchFunc(g,b)
	}

	if b.path.Size() > 0 {
		return b.path.PopBack()
	} else {
		return b.Pos
	}
}

/*
func PredictPath(g *Grid, e Entity) *PathQueue {
	var direc *PathQueue
	var temp Point
	visited := NewPointSet(g.Width())
	visited.Add(pos)

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
*/

func (b *Basic) AtGoal() bool {
	return b.Pos.Equals(b.Goal)
}

func (b *Basic) AtPredictGoal(pos Point) bool {
	return pos.Equals(b.Goal)
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

func (b *Basic) GetPos() Point {
	return b.Pos
}

func (b *Basic) SetPos(p Point) {
	b.Pos = p
}

func (b *Basic) GetGoal() Point {
	return b.Goal
}
