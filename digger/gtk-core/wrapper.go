package gtkcore

import (
	"github.com/xafero/digger-go/diggerapi"
)

type GtkWrap struct {
}

func NewGtkWrap() *GtkWrap {
	d := new(GtkWrap)
	return d
}

func (q *GtkWrap) CreateControl(dig diggerapi.DiggerRender) diggerapi.DrawingCore {
	return NewGtkDrawing(dig)
}

func (q *GtkWrap) CreateSource(dig diggerapi.DrawingCore, model *diggerapi.ColorModel) diggerapi.Refresher {
	gtkctrl := dig.(*GtkDrawing)
	gtkwid := gtkctrl.Control
	return NewGtkRefresher(gtkwid, model)
}
