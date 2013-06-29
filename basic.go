package main

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

//Basic implements Entity and will simply try to get from where it is to its goal
type Basic struct {
	X,Y int
	x,y int
	goalX, goalY int
	queue *point
	last_d int
	pred_at_goal bool
}

func (b *Basic) Move(g *Grid) Movement {
	//Decide on where to move based on information from the grid
	//for now, if g.At(x,y) != null, that space is occupied

	//Here is where we are going to do cool pathfinding stuff, or maybe in
	//another function that gets passed in? 
}

func (b *Basic) PredictPath(g *Grid) *Point {
	var direc *Point
	if b.AtPredictGoal() {
		pred_at_goal = true
	}
	if b.last_d != UP && g.At(b.x, b.y - 1) == nil && !pred_at_goal {
		b.y--
		b.last_d = UP
		direc = b.PredictPath(g)
	} if b.last_d != DOWN && g.At(b.x, b.y + 1) == nil && !pred_at_goal {
		b.y++
		b.last_d = DOWN
		direc = b.PredictPath(g)
	} if b.last_d != LEFT && g.At(b.x - 1, b.y) == nil && !pred_at_goal {
		b.x--
		b.last_d = LEFT
		direc = b.PredictPath(g)
	} if b.last_d != RIGHT && g.At(b.x + 1, b.y) == nil && !pred_at_goal {
		b.x++
		b.last_d = RIGHT
		direc = b.PredictPath(g)
	}
	direc.pushBack(x, y)

	return direc
}

func (b *Basic) AtGoal() bool {
	return b.X == b.goalX && b.Y == b.goalY
}

func (b *(Basic) AtPredictGoal() bool {
	return b.x == b.goalX && b.y == b.goalY
}

