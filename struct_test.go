package main

import (
	"fmt"
	"testing"
)

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
