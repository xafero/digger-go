package diggerapi

type DiggerCore interface {
	GetScores() ScoresCore
	GetMain() MainCore
}

type ScoresCore interface {
	DoScoreOctave()
}

type MainCore interface {
	DoRandNo(n int) int
}

type DiggerRender interface {
	GetPc() PcRender
	KeyDown(num int) bool
	KeyUp(num int) bool
}

type PcRender interface {
	GetWidth() int
	GetHeight() int
	GetPixels() []int
	GetCurrentSource() Refresher
}

type Refresher interface {
	NewPixels(x int, y int, w int, h int)
	NewPixelsAll()
	GetColor(index int) (float64, float64, float64)
}

type DrawingCore interface {
	GrabFocus()
	SetCanFocus(focusable bool)
}

type SourceCreator interface {
	CreateSource(dig DrawingCore, model *ColorModel) Refresher
	CreateControl(dig DiggerRender) DrawingCore
}
