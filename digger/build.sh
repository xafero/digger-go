#!/bin/sh

# REQ: sudo apt-get install libgtk-3-dev

go get github.com/gotk3/gotk3/gtk
go get github.com/gotk3/gotk3/gdk
go get github.com/gotk3/gotk3/glib
go get github.com/gotk3/gotk3/cairo
go get golang.design/x/thread
go get log

go build launcher.go
mv launcher goDigger

