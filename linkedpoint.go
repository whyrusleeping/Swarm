package main

type LinkPoint struct {
	Point
	parent *LinkPoint
	hval int
	cost int
}

func NewLinkPoint(pos Point, parent *LinkPoint) *LinkPoint {
	l := new(LinkPoint)
	l.X = pos.X
	l.Y = pos.Y
	l.parent = parent
	return l
}

type LPQueue struct {
	arr []*LinkPoint
}

func (l *LPQueue) Push(p *LinkPoint) {
	l.arr = append(l.arr, p)
}

func (l *LPQueue) Pop() (p *LinkPoint) {
	p = l.arr[0]
	l.arr = l.arr[1:]
	return
}

func (l *LPQueue) Size() int {
	return len(l.arr)
}
