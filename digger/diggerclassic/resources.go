package diggerclassic

import (
	"os"
	"path"
	"path/filepath"

	"github.com/gotk3/gotk3/gdk"
)

type Resources struct {
}

func NewResources() Resources {
	d := Resources{}
	return d
}

func FindResource(name string) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(exe)
	resource := path.Join(exePath, "Resources", name)
	return resource
}

func LoadImage(file string) *gdk.Pixbuf {
	img, _ := gdk.PixbufNewFromFile(file)
	return img
}
