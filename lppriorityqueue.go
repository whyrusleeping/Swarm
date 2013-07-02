package main

type LPPriorityQueue struct {
	arr []*LinkPoint
}

func NewPriorityQueue() *LPPriorityQueue {
	lp := new(LPPriorityQueue)
	lp.arr = []*LinkPoint{nil}
	return lp
}

func (lp *LPPriorityQueue) Push(p *LinkPoint) {
	lp.arr = append(lp.arr, p)
	done := false
	index := len(lp.arr)-1
	for !done && index != 1 {
		if lp.arr[index/2].val() > lp.arr[index].val() {
			temp := lp.arr[index/2]
			lp.arr[index/2] = lp.arr[index]
			lp.arr[index] = temp
			index/=2
		} else {
			done = true
		}
	}
}

func (lp *LPPriorityQueue) PopMin() *LinkPoint {
	min := lp.arr[1]
	var temp *LinkPoint
	lp.arr[1] = lp.arr[len(lp.arr)-1]
	lp.arr = lp.arr[:len(lp.arr)-1]
	done:= false
	index := 1
	for !done && index != (len(lp.arr) - 1) {
		//fmt.Println(index)
		//fmt.Println(lp.arr)
		if (index*2) <= (len(lp.arr) - 1) && lp.arr[index].val() > lp.arr[index*2].val() {
			if (index*2 + 1) <= (len(lp.arr) - 1) && lp.arr[index*2].val() > lp.arr[index*2 + 1].val() {
				//fmt.Println("A")
				temp = lp.arr[index*2 + 1]
				lp.arr[index*2 + 1] = lp.arr[index]
				lp.arr[index] = temp
				index = index*2 + 1
			} else {
				//fmt.Println("B")
				temp = lp.arr[index*2]
				lp.arr[index*2] = lp.arr[index]
				lp.arr[index] = temp
				index = index*2
			}
		} else if (index*2 + 1) <= (len(lp.arr) - 1) && lp.arr[index].val() > lp.arr[index*2 + 1].val() {
				//fmt.Println("C")
				temp = lp.arr[index*2 + 1]
				lp.arr[index*2 + 1] = lp.arr[index]
				lp.arr[index] = temp
				index = index*2 + 1
		} else {
			done = true
		}
	}
	return min
}
