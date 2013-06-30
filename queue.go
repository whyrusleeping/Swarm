package main

type PathQueue struct {
	arr []Point
	top int
}

func NewPathQueue() *PathQueue {
	pq := new(PathQueue)
	pq.arr = make([]Point,0,64)
	pq.top = 0
	return pq
}

func (pq *PathQueue) PushBack(p Point) {
	pq.arr = append(pq.arr, p)
}

func (pq *PathQueue) PopBack() (p Point) {
	p = pq.arr[len(pq.arr) - 1]
	pq.arr = pq.arr[:len(pq.arr)-1]
	return
}

func (pq *PathQueue) PopFront() (p Point) {
	p = pq.arr[0]
	pq.arr = pq.arr[1:]
	return
}

func (pq *PathQueue) PushFront(p Point) {
	pq.arr = append([]Point{p}, pq.arr...)
}

func (pq *PathQueue) Peek() Point {
	return pq.arr[len(pq.arr)-1]
}

func (pq *PathQueue) Size() int {
	return len(pq.arr)
}

func (pq *PathQueue) Contains(p Point) bool {
	for _,v := range pq.arr {
		if p.X == v.X && p.Y == v.Y {
			return true
		}
	}
	return false
}
