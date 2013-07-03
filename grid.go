package main

import (
	"os"
	"code.google.com/p/go.image/bmp"
	"fmt"
)

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

func LoadGridFromBitmap(path string) (*Grid, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, err := bmp.Decode(fi)
	if err != nil {
		return nil, err
	}

	size := img.Bounds().Max
	gr := NewGrid(size.X, size.Y)

	for i := 0; i < gr.x; i++ {
		for j := 0; j < gr.y; j++ {
			r,g,b,_ := img.At(i,j).RGBA()
			if r == 0 && g == 0 && b == 0 {
				fmt.Println(i,j)
				gr.v[i][j] = Wall{}
			}
		}
	}
	return gr, nil
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
	pos := e.GetPos()
	g.v[pos.X][pos.Y] = nil
	g.v[p.X][p.Y] = e
	e.SetPos(p)
}
