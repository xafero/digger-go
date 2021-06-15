package diggerclassic

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Refresher struct {
	area  *gtk.DrawingArea
	model *ColorModel
}

func NewRefresher(area *gtk.DrawingArea, model *ColorModel) *Refresher {
	d := new(Refresher)
	d.area = area
	d.model = model
	return d
}

func (q Refresher) NewPixels(x int, y int, w int, h int) {
	glib.TimeoutAdd(0, func() {
		q.area.QueueDrawArea(x, y, w, h)
	})
}

func (q Refresher) NewPixelsAll() {
	glib.TimeoutAdd(0, func() {
		q.area.QueueDraw()
	})
}
