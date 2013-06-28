package main

type Grid struct {
	v [][]Entity
	x,y int

	//A channel to lock during renders
	lock chan bool
}

func NewGrid(x,y int) *Grid {
	g := new(Grid)
	g.lock = make(chan bool)
	g.x = x
	g.y = y
	g.v = make([][]Entity, x)
	for i := 0; i < x; i++ {
		g.v[i] = make([]Entity, y)
	}
	return g
}

func (g *Grid) At(x,y int) Entity {
	//Do a bounds check
	//return the value

	/*Temporary code*/
	if x > 40 {
		return struct{}{} //This is an anonymous struct
	}
	return nil
}

func (g *Grid) InBounds(x,y int) bool {
	return false
}

func (g *Grid) Width() int {
	return g.x
}

func (g *Grid) Height() int {
	return g.y
}
