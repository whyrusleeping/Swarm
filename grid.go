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
func (g *Grid) At(p Point) Entity {
	if g.InBounds(p) {
		return g.v[p.X][p.Y]
	} else {
		return nil
	}
}

func (g *Grid) InBounds(p Point) bool {
	return p.X >= 0 && p.X < g.x && p.Y >= 0 && p.Y < g.y
}

func (g *Grid) Width() int {
	return g.x
}

func (g *Grid) Height() int {
	return g.y
}

func (g *Grid) Update(e Entity, p Point) {
	g.v[e.GetX()][e.GetY()] = nil
	g.v[p.X][p.Y] = e
	e.SetPos(p)
}
