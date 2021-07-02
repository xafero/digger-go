package gtkcore

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"github.com/xafero/digger-go/diggerapi"
)

type GtkRender struct {
	digger diggerapi.DiggerRender
}

func NewGtkRender(dig diggerapi.DiggerRender) *GtkRender {
	d := new(GtkRender)
	d.digger = dig
	return d
}

func (q *GtkRender) OnDrawn(da *gtk.DrawingArea, g *cairo.Context) bool {

	g.SetSourceRGB(0, 0, 0)
	g.Rectangle(0, 0, 3840, 2160)
	g.Fill()

	g.Scale(4, 4)

	var pc = q.digger.GetPc()

	var w = pc.GetWidth()
	var h = pc.GetHeight()
	var data = pc.GetPixels()

	shift := 1

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			arrayIndex := y*w + x
			mr, mg, mb := pc.GetCurrentSource().GetColor(data[arrayIndex])
			g.SetSourceRGB(mr, mg, mb)
			g.Rectangle(float64(x+shift), float64(y+shift), 1, 1)
			g.Fill()
		}
	}

	return false
}
