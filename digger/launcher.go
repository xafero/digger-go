package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"

	"github.com/xafero/digger-go/diggerclassic"
)

func main() {
	gtk.Init(nil)

	game := diggerclassic.NewDigger()
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

	err_icon := win.SetIconFromFile("./icons/digger.png")
	if err_icon != nil {
		log.Fatal("Unable to find icon:", err_icon)
	}

	win.Add(game.Control)
	win.ShowAll()

	win.Connect("key-press-event", game.OnKeyPress)
	win.Connect("key-release-event", game.OnKeyRelease)

	gtk.Main()
}
