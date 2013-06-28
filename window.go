package main

import (
	"github.com/whyrusleeping/sdl"
	"log"
	"runtime"
	"image/color"
)

type Window struct {
	DrawChan chan *Grid
	events chan sdl.Event
	X,Y int
}

func NewWindow(x,y int) *Window {
	w := new(Window)
	w.X = x
	w.Y = y
	w.DrawChan = make(chan *Grid)
	w.events = make(chan sdl.Event)
	return w
}

//sdl.PollEvent does not block normally, and that would make it difficult
//to integrate it into our flow, so here we loop and collect events into
//a channel
func (w *Window) poolEvents() {
	for {
		ev := sdl.PollEvent()
		if ev != nil {
			w.events <- ev
		}
	}
}

//Render the grid on the screen
func drawGrid(scr *sdl.Display, gr *Grid) {
	for x := 0; x < gr.Width(); x++ {
		for y := 0; y < gr.Height(); y++ {
			at := gr.At(x,y)
			if at != nil {
				scr.SetDrawColor(color.RGBA{255,255,255,128})
			} else {
				//Eventually use the entity.Color option
				scr.SetDrawColor(color.RGBA{0,0,0,128})
			}
			scr.DrawPoint(x,y)
		}
	}
	scr.Present()
}

func (w *Window) Run() {
	//Lock the thread because SDL uses openGL and openGL cant be
	//called across threads
	runtime.LockOSThread()
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		//If this fails, we kinda want to die for now
		panic(err)
	}

	//If this function quits unexpectedly, shut down SDL
	defer sdl.Quit()

	//Make a new screen object to render to
	scr, err := sdl.NewDisplay(w.X, w.Y, sdl.WINDOW_OPENGL)

	scr.SetTitle(title)

	events := make(chan sdl.Event)
	go w.poolEvents()
	for {
		select {
		case ev := <-events:
			switch ev.(type) {
			case *sdl.QuitEvent:
				return
			}
		case gr := <-w.DrawChan:
			drawGrid(scr, gr)
		}
	}
}

