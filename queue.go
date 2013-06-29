package main

type Point struct {
	X,Y int
}
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

func (pq *PathQueue) Peek() Point {
	return pq.arr[len(pq.arr)-1]
}

func (pq *PathQueue) Size() int {
	return len(pq.arr)
}
