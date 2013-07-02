package main

//Simple Breadth First Search, accepts a grid to find a path through and
//the entity that is going to travel the path
func BreadthFirst(g *Grid, e Entity) *PathQueue {
	//Initialize a queue to hold unexplored nodes
	q := new(LPQueue)

	//Initialize a new set to hold explored nodes
	visited := NewPointSet(g.Width())

	q.Push(NewLinkPoint(e.GetPos(),nil))

	valid := func(p Point) bool {
		if g.InBounds(p) && g.At(p) == nil && !visited.Find(p) {
			return true
		}
		return false
	}
	for q.Size() > 0 {
		n := q.Pop()
		visited.Add(n.Point)
		if n.Equals(e.GetGoal()) {
			final := NewPathQueue()
			final.PushBack(n.Point)
			for {
				n = n.parent
				if n == nil {
					break
				}
				final.PushBack(n.Point)
			}
			return final
		}

		if valid(n.Down()) {
			q.Push(NewLinkPoint(n.Down(), n))
		}
		if valid(n.Up()) {
			q.Push(NewLinkPoint(n.Up(), n))
		}
		if valid(n.Left()) {
			q.Push(NewLinkPoint(n.Left(),n))
		}
		if valid(n.Right()) {
			q.Push(NewLinkPoint(n.Right(),n))
		}
	}
	return nil
}
