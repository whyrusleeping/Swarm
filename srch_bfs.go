package main

func BreadthFirst(g *Grid, e Entity) *PathQueue {
	q := new(LPQueue)
	final := NewPathQueue()
	visited := NewPointSet(g.Width())

	q.Push(NewLinkPoint(Point{e.GetX(),e.GetY()},nil))

	valid := func(p Point) bool {
		if g.InBounds(p) && !visited.Find(p) && g.At(p) == nil {
			return true
		}
		return false
	}
	for q.Size() > 0 {
		n := q.Pop()
		visited.Add(n.Point)
		if n.Equals(e.GetGoal()) {
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
