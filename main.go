package main

func main() {
	//This whole function is temporary code, we want to make another struct that
	//is the simulation struct to contain the run loop and other fun variables
	runtime.GOMAXPROCS(4)
	w := NewWindow(400,400)
	g := NewGrid(400,400)
	go func() {w.DrawChan <- g}()
	w.Run()
}
