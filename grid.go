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

//Return the 'Entity' at the given coordinates
func (g *Grid) At(x,y int) Entity {
	//Do a bounds check
		//If bounds check fails, return nil
	//return the value
}

func (g *Grid) InBounds(x,y int) bool {
	//Check if X and Y are within the bounds 
}

func (g *Grid) Width() int {
	return g.x
}

func (g *Grid) Height() int {
	return g.y
}
