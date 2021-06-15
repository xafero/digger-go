package diggerclassic

import "github.com/gotk3/gotk3/gdk"

type AppletCompat struct {
}

func NewAppletCompat() AppletCompat {
	d := AppletCompat{}
	return d
}

func ConvertToLegacy(netCode gdk.EventKey) int {
	numeric := netCode.KeyVal()
	switch numeric {
	case gdk.KEY_Left:
	case gdk.KEY_leftarrow:
		return 1006
	case gdk.KEY_Right:
	case gdk.KEY_rightarrow:
		return 1007
	case gdk.KEY_Up:
	case gdk.KEY_uparrow:
		return 1004
	case gdk.KEY_Down:
	case gdk.KEY_downarrow:
		return 1005
	case gdk.KEY_F1:
		return 1008
	case gdk.KEY_F10:
		return 1021
	case gdk.KEY_plus:
		return 1031
	case gdk.KEY_minus:
		return 1032
	}
	ascii := int(numeric)
	return ascii
}

func GetParameter(name string) string {
	return " "
}
