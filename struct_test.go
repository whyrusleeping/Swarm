package main

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	//Make a new priority queue
	pq := NewPriorityQueue()
	//Insert a bunch of things into the queue
	lp := &LinkPoint{}
	lp.cost = 100
	pq.Push(lp)
	lp = &LinkPoint{}
	lp.cost = 50
	pq.Push(lp)
	lp = &LinkPoint{}
	lp.cost = 200
	pq.Push(lp)
	lp = &LinkPoint{}
	lp.cost = 1
	pq.Push(lp)
	lp = &LinkPoint{}
	lp.cost = 4
	pq.Push(lp)
	//Make sure they come out in the right order
	fmt.Println(pq.arr)
	lp = pq.PopMin()
	if lp.cost != 1 {
		t.Fail()
		fmt.Println(lp)
	}
	fmt.Println(pq.arr)
	lp = pq.PopMin()
	if lp.cost != 4 {
		fmt.Println(lp)
		t.Fail()
	}
	fmt.Println(pq.arr)
	lp = pq.PopMin()
	if lp.cost != 50 {
		fmt.Println(lp)
		t.Fail()
	}
	fmt.Println(pq.arr)
	lp = pq.PopMin()
	if lp.cost != 100 {
		fmt.Println(lp)
		t.Fail()
	}
	fmt.Println(pq.arr)
	lp = pq.PopMin()
	if lp.cost != 200 {
		fmt.Println(lp)
		t.Fail()
	}
	//If they dont. call t.Fail()
}

func TestPointSet(t *testing.T) {
	p := NewPointSet(20)
	p.Add(Point{6,7})
	p.Add(Point{5,7})
	p.Add(Point{6,6})
	p.Add(Point{6,4})
	p.Add(Point{3,7})
	p.Add(Point{6,17})
	p.Add(Point{6,1})

	fmt.Println("Starting Finds...")
	if !p.Find(Point{3,7}) {
		t.Fail()
	}

	if p.Find(Point{15,2}) {
		t.Fail()
	}
}

func TestAvlTree(t *testing.T) {
	tr := makeAvlNode(1)
	for i := 2; i < 5000; i++ {
		tr.insert(i)
		checkAVL(&tr)
	}
	fmt.Printf("height: %d - %d\n", tr.height(), recurHeight(tr))
}

func BenchmarkAvlTreeInsert(b *testing.B) {
	tr := makeAvlNode(1)
	for i := 2; i < b.N; i++ {
		tr.insert(i)
		checkAVL(&tr)
	}
}

func BenchmarkAvlTreeFind(b *testing.B) {
	b.StopTimer()
	tr := makeAvlNode(1)
	for i := 2; i < 200000; i++ {
		tr.insert(i)
		checkAVL(&tr)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.find(i)
	}
}
