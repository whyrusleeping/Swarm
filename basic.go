package main

import "fmt"

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
	x,y int
	goalX, goalY int
	//queue *Point
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
	b.x = b.X
	b.y = b.Y

	return b
}

func (b *Basic) Move(g *Grid) Movement {
	//Decide on where to move based on information from the grid
	//for now, if g.At(x,y) != null, that space is occupied

	//Here is where we are going to do cool pathfinding stuff, or maybe in
	//another function that gets passed in? 
	return Movement{}
}

func (b *Basic) PredictPath(g *Grid) *PathQueue {
	var direc *PathQueue

	if b.AtPredictGoal() {
		b.pred_at_goal = true
		direc = NewPathQueue()
	}
	if b.last_d != UP && g.At(b.x, b.y - 1) == nil && !b.pred_at_goal {
		b.y--
		b.last_d = UP
		direc = b.PredictPath(g)
	}
	if b.last_d != DOWN && g.At(b.x, b.y + 1) == nil && !b.pred_at_goal {
		b.y++
		b.last_d = DOWN
		direc = b.PredictPath(g)
	}
	if b.last_d != LEFT && g.At(b.x - 1, b.y) == nil && !b.pred_at_goal {
		b.x--
		b.last_d = LEFT
		direc = b.PredictPath(g)
	}
	if b.last_d != RIGHT && g.At(b.x + 1, b.y) == nil && !b.pred_at_goal {
		b.x++
		b.last_d = RIGHT
		direc = b.PredictPath(g)
	}
	if !b.AtDeadEnd(g) {
		direc.PushBack(Point{b.x,b.y})
	}
	return direc
}

func (b *Basic) AtGoal() bool {
	fmt.Println("GOOOOOOOOOOOOOOOOOOOOAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAL!")
	return b.X == b.goalX && b.Y == b.goalY
}

func (b *Basic) AtPredictGoal() bool {
	return b.x == b.goalX && b.y == b.goalY
}

func (b *Basic) AtDeadEnd(g *Grid) bool {
	 count := 0

	if b.last_d != UP && g.At(b.x, b.y-1) != nil {
		count++
	}
	if b.last_d != DOWN && g.At(b.x, b.y+1) != nil {
		count++
	}
	if b.last_d != LEFT && g.At(b.x-1, b.y) != nil {
		count++
	}
	if b.last_d != RIGHT && g.At(b.x+1, b.y) != nil {
		count++
	}

	return count == 3
}

