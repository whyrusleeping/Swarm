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
	if g.InBounds(x,y) {
		return g.v[x][y]
	} else {
		return nil
	}
}

func (g *Grid) InBounds(x,y int) bool {
	return x > 0 && x < g.x && y > 0 && y < g.y
}

func (g *Grid) Width() int {
	return g.x
}

func (g *Grid) Height() int {
	return g.y
}

func (g *Grid) Update(e Entity, m Movement) {
	//Move e as specified by m
}
