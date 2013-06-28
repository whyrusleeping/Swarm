package main

func main() {
	runtime.GOMAXPROCS(4)
	w := NewWindow(400,400)
	g := NewGrid(400,400)
	go func() {w.DrawChan <- g}()
	w.Run()

}
