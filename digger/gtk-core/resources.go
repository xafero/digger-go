package gtkcore

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/xafero/digger-go/diggerapi"
)

func LoadGtkImage(file string) *gdk.Pixbuf {
	bytes, err := diggerapi.Asset(file)
	if err != nil {
		panic(err)
	}
	img, err := gdk.PixbufNewFromBytesOnly(bytes)
	if err != nil {
		panic(err)
	}
	return img
}
