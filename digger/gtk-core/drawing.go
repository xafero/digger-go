package gtkcore

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/xafero/digger-go/diggerapi"
)

type GtkDrawing struct {
	Control  *gtk.DrawingArea
	Renderer *GtkRender
}

func NewGtkDrawing(dig diggerapi.DiggerRender) *GtkDrawing {
	d := new(GtkDrawing)

	// Custom drawing area
	ctrl, cerr := gtk.DrawingAreaNew()
	if cerr != nil {
		log.Fatal("Unable to create area:", cerr)
	}

	d.Control = ctrl
	d.Renderer = NewGtkRender(dig)
	d.Control.Connect("draw", d.Renderer.OnDrawn)

	return d
}

func (q *GtkDrawing) GrabFocus() {
	q.Control.GrabFocus()
}

func (q *GtkDrawing) SetCanFocus(focusable bool) {
	q.Control.SetCanFocus(focusable)
}
