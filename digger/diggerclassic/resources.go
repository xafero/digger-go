package diggerclassic

import (
	"os"
	"path"
	"path/filepath"
)

func FindResource(name string) string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(exe)
	resource := path.Join(exePath, "Resources", name)
	return resource
}
