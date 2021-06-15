package diggerclassic

type Input struct {
	dig          *digger
	leftpressed  bool
	rightpressed bool
	uppressed    bool
	downpressed  bool
	f1pressed    bool
	firepressed  bool
	minuspressed bool
	pluspressed  bool
	f10pressed   bool
	escape       bool
	keypressed   int
	akeypressed  int
	dynamicdir   int
	staticdir    int
	joyx         int
	joyy         int
	joybut1      bool
	joybut2      bool
	keydir       int
	jleftthresh  int
	jupthresh    int
	jrightthresh int
	jdownthresh  int
	joyanax      int
	joyanay      int
	firepflag    bool
	joyflag      bool
}

func (rcvr *Input) Key_downpressed() {
	rcvr.downpressed = true
	rcvr.dynamicdir = 6
	rcvr.staticdir = 6
}

func (rcvr *Input) Key_downreleased() {
	rcvr.downpressed = false
	if rcvr.dynamicdir == 6 {
		rcvr.setdirec()
	}
}
func (rcvr *Input) Key_f1pressed() {
	rcvr.firepressed = true
	rcvr.f1pressed = true
}
func (rcvr *Input) Key_f1released() {
	rcvr.f1pressed = false
}
func (rcvr *Input) Key_leftpressed() {
	rcvr.leftpressed = true
	rcvr.dynamicdir = 4
	rcvr.staticdir = 4
}
func (rcvr *Input) Key_leftreleased() {
	rcvr.leftpressed = false
	if rcvr.dynamicdir == 4 {
		rcvr.setdirec()
	}
}
func (rcvr *Input) Key_rightpressed() {
	rcvr.rightpressed = true
	rcvr.dynamicdir = 0
	rcvr.staticdir = 0
}
func (rcvr *Input) Key_rightreleased() {
	rcvr.rightpressed = false
	if rcvr.dynamicdir == 0 {
		rcvr.setdirec()
	}
}
func (rcvr *Input) Key_uppressed() {
	rcvr.uppressed = true
	rcvr.dynamicdir = 2
	rcvr.staticdir = 2
}
func (rcvr *Input) Key_upreleased() {
	rcvr.uppressed = false
	if rcvr.dynamicdir == 2 {
		rcvr.setdirec()
	}
}

func NewInput(d *digger) *Input {
	rcvr := new(Input)
	rcvr.leftpressed = false
	rcvr.rightpressed = false
	rcvr.uppressed = false
	rcvr.downpressed = false
	rcvr.f1pressed = false
	rcvr.firepressed = false
	rcvr.escape = false
	rcvr.keypressed = 0
	rcvr.dynamicdir = -1
	rcvr.staticdir = -1
	rcvr.joyx = 0
	rcvr.joyy = 0
	rcvr.joybut1 = false
	rcvr.joybut2 = false
	rcvr.keydir = 0
	rcvr.jleftthresh = 0
	rcvr.jupthresh = 0
	rcvr.jrightthresh = 0
	rcvr.jdownthresh = 0
	rcvr.joyanax = 0
	rcvr.joyanay = 0
	rcvr.firepflag = false
	rcvr.joyflag = false
	rcvr.dig = d
	return rcvr
}

func (rcvr *Input) checkkeyb() {
	if rcvr.pluspressed {
		if rcvr.dig.FrameTime > MIN_RATE {
			rcvr.dig.FrameTime -= 5
		}
	}
	if rcvr.minuspressed {
		if rcvr.dig.FrameTime < MAX_RATE {
			rcvr.dig.FrameTime += 5
		}
	}
	if rcvr.f10pressed {
		rcvr.escape = true
	}
}

func (rcvr *Input) detectjoy() {
	rcvr.joyflag = false
	rcvr.staticdir = -1
	rcvr.dynamicdir = -1
}

func (rcvr *Input) getasciikey(make int) int {
	//var k int
	if make == ' ' || make >= 'a' && make <= 'z' || make >= '0' && make <= '9' {
		return make
	} else {
		return 0
	}
}

func (rcvr *Input) getdir() int {
	bp2 := rcvr.keydir
	return bp2
}

func (rcvr *Input) initkeyb() {
}

func (rcvr *Input) processkey(key int) {
	rcvr.keypressed = key
	if key > 0x80 {
		rcvr.akeypressed = key & 0x7f
	}
	switch key {
	case 0x4b:
		rcvr.Key_leftpressed()
	case 0xcb:
		rcvr.Key_leftreleased()
	case 0x4d:
		rcvr.Key_rightpressed()
	case 0xcd:
		rcvr.Key_rightreleased()
	case 0x48:
		rcvr.Key_uppressed()
	case 0xc8:
		rcvr.Key_upreleased()
	case 0x50:
		rcvr.Key_downpressed()
	case 0xd0:
		rcvr.Key_downreleased()
	case 0x3b:
		rcvr.Key_f1pressed()
	case 0xbb:
		rcvr.Key_f1released()
	case 0x78:
		rcvr.f10pressed = true
	case 0xf8:
		rcvr.f10pressed = false
	case 0x2b:
		rcvr.pluspressed = true
	case 0xab:
		rcvr.pluspressed = false
	case 0x2d:
		rcvr.minuspressed = true
	case 0xad:
		rcvr.minuspressed = false
	}
}

func (rcvr *Input) readdir() {
	rcvr.keydir = rcvr.staticdir
	if rcvr.dynamicdir != -1 {
		rcvr.keydir = rcvr.dynamicdir
	}
	rcvr.staticdir = -1
	if rcvr.f1pressed || rcvr.firepressed {
		rcvr.firepflag = true
	} else {
		rcvr.firepflag = false
	}
	rcvr.firepressed = false
}

func (rcvr *Input) readjoy() {
}

func (rcvr *Input) setdirec() {
	rcvr.dynamicdir = -1
	if rcvr.uppressed {
		rcvr.dynamicdir = 2
		rcvr.staticdir = 2
	}
	if rcvr.downpressed {
		rcvr.dynamicdir = 6
		rcvr.staticdir = 6
	}
	if rcvr.leftpressed {
		rcvr.dynamicdir = 4
		rcvr.staticdir = 4
	}
	if rcvr.rightpressed {
		rcvr.dynamicdir = 0
		rcvr.staticdir = 0
	}
}

func (rcvr *Input) teststart() bool {
	startf := false
	if rcvr.keypressed != 0 && rcvr.keypressed&0x80 == 0 && rcvr.keypressed != 27 {
		startf = true
		rcvr.joyflag = false
		rcvr.keypressed = 0
	}
	if !startf {
		return false
	}
	return true
}
