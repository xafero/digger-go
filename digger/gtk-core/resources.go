package gtkcore

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/xafero/digger-go/diggerclassic"
)

func LoadGtkImage(file string) *gdk.Pixbuf {
	bytes, err := diggerclassic.Asset(file)
	if err != nil {
		panic(err)
	}
	img, err := gdk.PixbufNewFromBytesOnly(bytes)
	if err != nil {
		panic(err)
	}
	return img
}
