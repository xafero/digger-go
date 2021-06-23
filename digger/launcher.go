package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"

	"github.com/xafero/digger-go/diggerclassic"
	gtkcore "github.com/xafero/digger-go/gtk-core"
)

func main() {
	gtk.Init(nil)

	ctx := gtkcore.NewGtkWrap()
	game := diggerclassic.NewDigger(ctx)
	game.SetFocusable(true)
	game.Init()
	game.Start()

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Digger Remastered")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(int(float64(game.Width)*4.03), int(float64(game.Height)*4.15))
	win.SetPosition(gtk.WIN_POS_CENTER)

	icon := gtkcore.LoadGtkImage("icons/digger.png")
	win.SetIcon(icon)

	gtkctrl := game.Control.(*gtkcore.GtkDrawing)
	gtkwid := gtkctrl.Control

	win.Add(gtkwid)
	win.ShowAll()

	win.Connect("key-press-event", gtkcore.CreateOnKeyPress(game))
	win.Connect("key-release-event", gtkcore.CreateOnKeyRelease(game))

	gtk.Main()
}
