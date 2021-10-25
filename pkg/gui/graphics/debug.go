package graphics

/*
func highlightElement(drawer Canvas, layer Layer) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 0)
	// ↑
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()))
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()+drawer.Height()))
	// →
	imd.Push(pixel.V(drawer.Position().X()+drawer.Width(), drawer.Position().Y()+drawer.Height()))
	// ↓
	imd.Push(pixel.V(drawer.Position().X()+drawer.Width(), drawer.Position().Y()))
	// ←
	imd.Push(pixel.V(drawer.Position().X(), drawer.Position().Y()))

	imd.Rectangle(1)
	imd.Draw(layer.target())
}

var memstats runtime.MemStats

func PrintMemStats() {
	mbytes := uint64(1024 * 1024)
	time.AfterFunc(time.Second * 1, func() {
		runtime.ReadMemStats(&memstats)
		fmt.Println("mem.Alloc: ", memstats.Alloc/mbytes, "MB")
		fmt.Println("mem.TotalAlloc:", memstats.TotalAlloc/mbytes, "MB")
		fmt.Println("mem.HeapAlloc:", memstats.HeapAlloc/mbytes, "MB")
		fmt.Println("mem.Garbage collected:", memstats.NumGC, "times")
		fmt.Println("-----")
		PrintMemStats()
	})
}

*/
