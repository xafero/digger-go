package gtkcore

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/xafero/digger-go/diggerapi"
)

type GtkRefresher struct {
	area  *gtk.DrawingArea
	model *diggerapi.ColorModel
}

func NewGtkRefresher(area *gtk.DrawingArea, model *diggerapi.ColorModel) *GtkRefresher {
	d := new(GtkRefresher)
	d.area = area
	d.model = model
	return d
}

func (q *GtkRefresher) NewPixels(x int, y int, w int, h int) {
	glib.TimeoutAdd(0, func() {
		q.area.QueueDrawArea(x, y, w, h)
	})
}

func (q *GtkRefresher) NewPixelsAll() {
	glib.TimeoutAdd(0, func() {
		q.area.QueueDraw()
	})
}

func (q *GtkRefresher) GetColor(index int) (float64, float64, float64) {
	return q.model.GetColor(index)
}
