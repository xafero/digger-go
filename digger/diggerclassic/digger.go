package diggerclassic

import (
	"log"
	"time"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	"golang.design/x/thread"
)

var MAX_RATE = 200
var MIN_RATE = 40

type digger struct {
	Width              int
	Height             int
	FrameTime          int
	gamethread         thread.Thread
	subaddr            string
	Bags               *Bags
	Main               *Main
	Sound              *Sound
	newSound           *NewSound
	Monster            *Monster
	Scores             *Scores
	Sprite             *Sprite
	Drawing            *Drawing
	Input              *Input
	Pc                 *Pc
	diggerx            int
	diggery            int
	diggerh            int
	diggerv            int
	diggerrx           int
	diggerry           int
	digmdir            int
	digdir             int
	digtime            int
	rechargetime       int
	firex              int
	firey              int
	firedir            int
	expsn              int
	deathstage         int
	deathbag           int
	deathani           int
	deathtime          int
	startbonustimeleft int
	bonustimeleft      int
	eatmsc             int
	emocttime          int
	emmask             int
	emfield            []byte
	digonscr           bool
	notfiring          bool
	bonusvisible       bool
	bonusmode          bool
	diggervisible      bool
	time               int64
	ftime              int64
	embox              []int
	deatharc           []int
	Control            *gtk.DrawingArea
}

func NewDigger() *digger {
	rcvr := new(digger)

	rcvr.Width = 320
	rcvr.Height = 200
	rcvr.FrameTime = 66

	rcvr.deatharc = []int{3, 5, 6, 6, 5, 3, 0}     // [7]
	rcvr.embox = []int{8, 12, 12, 9, 16, 12, 6, 9} // [8]
	rcvr.emfield = []byte{                         //[150]
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	rcvr.diggerx = 0
	rcvr.diggery = 0
	rcvr.diggerh = 0
	rcvr.diggerv = 0
	rcvr.diggerrx = 0
	rcvr.diggerry = 0
	rcvr.digmdir = 0
	rcvr.digdir = 0
	rcvr.digtime = 0
	rcvr.rechargetime = 0
	rcvr.firex = 0
	rcvr.firey = 0
	rcvr.firedir = 0
	rcvr.expsn = 0
	rcvr.deathstage = 0
	rcvr.deathbag = 0
	rcvr.deathani = 0
	rcvr.deathtime = 0
	rcvr.startbonustimeleft = 0
	rcvr.bonustimeleft = 0
	rcvr.eatmsc = 0
	rcvr.emocttime = 0
	rcvr.emmask = 0
	rcvr.digonscr = false
	rcvr.notfiring = false
	rcvr.bonusvisible = false
	rcvr.bonusmode = false
	rcvr.diggervisible = false
	rcvr.ftime = 50
	rcvr.time = 50
	rcvr.Bags = NewBags(rcvr)
	rcvr.Main = NewMain(rcvr)
	rcvr.Sound = NewSoundObj(rcvr)
	rcvr.newSound = NewNewSound(rcvr)
	rcvr.Monster = NewMonster(rcvr)
	rcvr.Scores = NewScores(rcvr)
	rcvr.Sprite = NewSprite(rcvr)
	rcvr.Drawing = NewDrawing(rcvr)
	rcvr.Input = NewInput(rcvr)
	rcvr.Pc = NewPc(rcvr)

	// Custom drawing area
	ctrl, cerr := gtk.DrawingAreaNew()
	if cerr != nil {
		log.Fatal("Unable to create area:", cerr)
	}

	rcvr.Control = ctrl
	rcvr.Control.Connect("draw", rcvr.OnDrawn)

	return rcvr
}

func (d *digger) GetScores() *Scores {
	return d.Scores
}

func (d *digger) GetMain() *Main {
	return d.Main
}

func (d *digger) SetFocusable(v bool) {

}

func (rcvr *digger) checkdiggerunderbag(h int, v int) bool {
	if rcvr.digmdir == 2 || rcvr.digmdir == 6 {
		if (rcvr.diggerx-12)/20 == h {
			if (rcvr.diggery-18)/18 == v || (rcvr.diggery-18)/18+1 == v {
				return true
			}
		}
	}
	return false
}

func (q *digger) countem() int {
	var x int
	var y int
	n := 0
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if (q.emfield[y*15+x] & byte(q.emmask)) != 0 {
				n++
			}
		}
	}
	return n
}

func (rcvr *digger) createbonus() {
	rcvr.bonusvisible = true
	rcvr.Drawing.drawbonus(292, 18)
}

func (rcvr *digger) Destroy() {
	if rcvr.gamethread != nil {
		rcvr.gamethread.Terminate()
	}
}

func (rcvr *digger) diggerdie() {
	var clbits int
	switch rcvr.deathstage {
	case 1:
		if rcvr.Bags.bagy(rcvr.deathbag)+6 > rcvr.diggery {
			rcvr.diggery = rcvr.Bags.bagy(rcvr.deathbag) + 6
		}
		rcvr.Drawing.drawdigger(15, rcvr.diggerx, rcvr.diggery, false)
		rcvr.Main.incpenalty()
		if rcvr.Bags.getbagdir(rcvr.deathbag)+1 == 0 {
			rcvr.Sound.soundddie()
			rcvr.deathtime = 5
			rcvr.deathstage = 2
			rcvr.deathani = 0
			rcvr.diggery -= 6
		}
	case 2:
		if rcvr.deathtime != 0 {
			rcvr.deathtime--
			break
		}
		if rcvr.deathani == 0 {
			rcvr.Sound.music(2)
		}
		clbits = rcvr.Drawing.drawdigger(14-rcvr.deathani, rcvr.diggerx, rcvr.diggery, false)
		rcvr.Main.incpenalty()
		if rcvr.deathani == 0 && clbits&0x3f00 != 0 {
			rcvr.Monster.killmonsters(clbits)
		}
		if rcvr.deathani < 4 {
			rcvr.deathani++
			rcvr.deathtime = 2
		} else {
			rcvr.deathstage = 4
			if rcvr.Sound.musicflag {
				rcvr.deathtime = 60
			} else {
				rcvr.deathtime = 10
			}
		}
	case 3:
		rcvr.deathstage = 5
		rcvr.deathani = 0
		rcvr.deathtime = 0
	case 5:
		if rcvr.deathani >= 0 && rcvr.deathani <= 6 {
			rcvr.Drawing.drawdigger(15, rcvr.diggerx, rcvr.diggery-rcvr.deatharc[rcvr.deathani], false)
			if rcvr.deathani == 6 {
				rcvr.Sound.musicoff()
			}
			rcvr.Main.incpenalty()
			rcvr.deathani++
			if rcvr.deathani == 1 {
				rcvr.Sound.soundddie()
			}
			if rcvr.deathani == 7 {
				rcvr.deathtime = 5
				rcvr.deathani = 0
				rcvr.deathstage = 2
			}
		}
	case 4:
		if rcvr.deathtime != 0 {
			rcvr.deathtime--
		} else {
			rcvr.Main.setdead(true)
		}
	}
}

func (rcvr *digger) dodigger() {
	rcvr.newframe()
	if rcvr.expsn != 0 {
		rcvr.drawexplosion()
	} else {
		rcvr.updatefire()
	}
	if rcvr.diggervisible {
		if rcvr.digonscr {
			if rcvr.digtime != 0 {
				rcvr.Drawing.drawdigger(rcvr.digmdir, rcvr.diggerx, rcvr.diggery, rcvr.notfiring && rcvr.rechargetime == 0)
				rcvr.Main.incpenalty()
				rcvr.digtime--
			} else {
				rcvr.updatedigger()
			}
		} else {
			rcvr.diggerdie()
		}
	}
	if rcvr.bonusmode && rcvr.digonscr {
		rcvr.newSound.stopNormalBackgroundMusic()
		if rcvr.bonustimeleft != 0 {
			rcvr.bonustimeleft--
			if rcvr.startbonustimeleft != 0 || rcvr.bonustimeleft < 20 {
				rcvr.startbonustimeleft--
				if rcvr.bonustimeleft&1 != 0 {
					rcvr.Pc.ginten(0)
					rcvr.Sound.soundbonus()
					rcvr.newSound.stopBonusBackgroundMusic()
					rcvr.newSound.startBonusPulse()
				} else {
					rcvr.newSound.stopBonusBackgroundMusic()
					rcvr.newSound.startBonusPulse()
					rcvr.Pc.ginten(1)
					rcvr.Sound.soundbonus()
				}
				if rcvr.startbonustimeleft == 0 {
					rcvr.newSound.endBonusPulse()
					rcvr.Sound.music(0)
					rcvr.Sound.soundbonusoff()
					rcvr.Pc.ginten(1)
					rcvr.newSound.startBonusBackgroundMusic()
				}
			}
		} else {
			rcvr.endbonusmode()
			rcvr.Sound.soundbonusoff()
			rcvr.Sound.music(1)
			rcvr.newSound.endBonusPulse()
			rcvr.newSound.stopBonusBackgroundMusic()
			rcvr.newSound.startNormalBackgroundMusic()
		}
	}
	if rcvr.bonusmode && !rcvr.digonscr {
		rcvr.endbonusmode()
		rcvr.newSound.endBonusPulse()
		rcvr.newSound.stopBonusBackgroundMusic()
		rcvr.newSound.startNormalBackgroundMusic()
		rcvr.Sound.soundbonusoff()
		rcvr.Sound.music(1)
	}
	if rcvr.emocttime > 0 {
		rcvr.emocttime--
	}
}

func (q *digger) drawemeralds() {
	var x int
	var y int
	q.emmask = 1 << q.Main.getcplayer()
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if (q.emfield[y*15+x] & byte(q.emmask)) != 0 {
				q.Drawing.drawemerald(x*20+12, y*18+21)
			}
		}
	}
}

func (rcvr *digger) drawexplosion() {
	switch rcvr.expsn {
	case 1:
		rcvr.newSound.fireEnd()
		rcvr.Sound.soundexplode()
		rcvr.newSound.playExplode()
		fallthrough
	case 2:
		fallthrough
	case 3:
		rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, rcvr.expsn)
		rcvr.Main.incpenalty()
		rcvr.expsn++
	default:
		rcvr.killfire()
		rcvr.expsn = 0
	}
}

func (rcvr *digger) endbonusmode() {
	rcvr.bonusmode = false
	rcvr.Pc.ginten(0)
}

func (rcvr *digger) erasebonus() {
	if rcvr.bonusvisible {
		rcvr.bonusvisible = false
		rcvr.Sprite.erasespr(14)
	}
	rcvr.Pc.ginten(0)
}

func (rcvr *digger) erasedigger() {
	rcvr.Sprite.erasespr(0)
	rcvr.diggervisible = false
}

func (rcvr *digger) GetAppletInfo() string {
	return "The Digger Remastered -- http://www.digger.org, Copyright (c) Andrew Jenner & Marek Futrega / MAF"
}

func (rcvr *digger) getfirepflag() bool {
	return rcvr.Input.firepflag
}

func (q *digger) hitemerald(x int, y int, rx int, ry int, dir int) bool {
	hit := false
	var r int
	if dir < 0 || dir > 6 || dir&1 != 0 {
		return hit
	}
	if dir == 0 && rx != 0 {
		x++
	}
	if dir == 6 && ry != 0 {
		y++
	}
	if dir == 0 || dir == 4 {
		r = rx
	} else {
		r = ry
	}
	if (q.emfield[y*15+x] & byte(q.emmask)) != 0 {
		if r == q.embox[dir] {
			q.Drawing.drawemerald(x*20+12, y*18+21)
			q.Main.incpenalty()
		}
		if r == q.embox[dir+1] {
			q.Drawing.eraseemerald(x*20+12, y*18+21)
			q.Main.incpenalty()
			hit = true
			q.emfield[y*15+x] = q.emfield[y*15+x] & byte(^q.emmask)
		}
	}
	return hit
}

func (rcvr digger) Init() {
	if rcvr.gamethread != nil {
		rcvr.gamethread.Terminate()
	}
	rcvr.subaddr = GetParameter("submit")

	rcvr.FrameTime, _ = strconv.Atoi(GetParameter("speed"))
	if rcvr.FrameTime > MAX_RATE {
		rcvr.FrameTime = MAX_RATE
	} else if rcvr.FrameTime < MIN_RATE {
		rcvr.FrameTime = MIN_RATE
	}

	rcvr.Pc.pixels = make([]int, 65536)

	for i := 0; i < 2; i++ {
		model := NewColorModel(8, 4, rcvr.Pc.pal[i][0], rcvr.Pc.pal[i][1], rcvr.Pc.pal[i][2])
		rcvr.Pc.source[i] = NewRefresher(rcvr.Control, &model)
		rcvr.Pc.source[i].NewPixelsAll()
	}

	rcvr.Pc.currentSource = rcvr.Pc.source[0]

	rcvr.gamethread = thread.New()
	rcvr.gamethread.CallNonBlock(rcvr.Run)
}

func (rcvr digger) initbonusmode() {
	rcvr.bonusmode = true
	rcvr.erasebonus()
	rcvr.Pc.ginten(1)
	rcvr.bonustimeleft = 250 - rcvr.Main.levof10()*20
	rcvr.startbonustimeleft = 20
	rcvr.eatmsc = 1
}

func (rcvr digger) initdigger() {
	rcvr.diggerv = 9
	rcvr.digmdir = 4
	rcvr.diggerh = 7
	rcvr.diggerx = rcvr.diggerh*20 + 12
	rcvr.digdir = 0
	rcvr.diggerrx = 0
	rcvr.diggerry = 0
	rcvr.digtime = 0
	rcvr.digonscr = true
	rcvr.deathstage = 1
	rcvr.diggervisible = true
	rcvr.diggery = rcvr.diggerv*18 + 18
	rcvr.Sprite.movedrawspr(0, rcvr.diggerx, rcvr.diggery)
	rcvr.notfiring = true
	rcvr.emocttime = 0
	rcvr.bonusvisible = false
	rcvr.bonusmode = false
	rcvr.Input.firepressed = false
	rcvr.expsn = 0
	rcvr.rechargetime = 0
}

func (rcvr digger) KeyDown(key int) bool {
	switch key {
	case 1006:
		rcvr.Input.processkey(0x4b)
	case 1007:
		rcvr.Input.processkey(0x4d)
	case 1004:
		rcvr.Input.processkey(0x48)
	case 1005:
		rcvr.Input.processkey(0x50)
	case 1008:
		rcvr.Input.processkey(0x3b)
	case 1021:
		rcvr.Input.processkey(0x78)
	case 1031:
		rcvr.Input.processkey(0x2b)
	case 1032:
		rcvr.Input.processkey(0x2d)
	default:
		key &= 0x7f
		if key >= 65 && key <= 90 {
			key += 97 - 65
		}
		rcvr.Input.processkey(key)
	}
	return true
}

func (rcvr digger) KeyUp(key int) bool {
	switch key {
	case 1006:
		rcvr.Input.processkey(0xcb)
	case 1007:
		rcvr.Input.processkey(0xcd)
	case 1004:
		rcvr.Input.processkey(0xc8)
	case 1005:
		rcvr.Input.processkey(0xd0)
	case 1008:
		rcvr.Input.processkey(0xbb)
	case 1021:
		rcvr.Input.processkey(0xf8)
	case 1031:
		rcvr.Input.processkey(0xab)
	case 1032:
		rcvr.Input.processkey(0xad)
	default:
		key &= 0x7f
		if key >= 65 && key <= 90 {
			key += 97 - 65
		}
		rcvr.Input.processkey(0x80 | key)
	}
	return true
}

func (rcvr digger) killdigger(stage int, bag int) {
	if rcvr.deathstage < 2 || rcvr.deathstage > 4 {
		rcvr.newSound.stopNormalBackgroundMusic()
		rcvr.newSound.stopBonusBackgroundMusic()
		rcvr.newSound.playDeath()
		rcvr.digonscr = false
		rcvr.deathstage = stage
		rcvr.deathbag = bag
	}
}

func (q digger) killemerald(x int, y int) {
	if (q.emfield[y*15+x+15] & byte(q.emmask)) != 0 {
		q.emfield[y*15+x+15] = q.emfield[y*15+x+15] & byte(^q.emmask)
		q.Drawing.eraseemerald(x*20+12, (y+1)*18+21)
	}
}

func (rcvr digger) killfire() {
	if !rcvr.notfiring {
		rcvr.notfiring = true
		rcvr.Sprite.erasespr(15)
		rcvr.Sound.soundfireoff()
		rcvr.newSound.fireEnd()
	}
}

func (q digger) makeemfield() {
	var x int
	var y int
	q.emmask = 1 << q.Main.getcplayer()
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if q.Main.getlevch(x, y, q.Main.levplan()) == 'C' {
				q.emfield[y*15+x] = q.emfield[y*15+x] | byte(q.emmask)
			} else {
				q.emfield[y*15+x] = q.emfield[y*15+x] & byte(^q.emmask)
			}
		}
	}
}

func (rcvr digger) newframe() {
	rcvr.Input.checkkeyb()
	rcvr.time += int64(rcvr.FrameTime)
	l := rcvr.time - rcvr.Pc.gethrt()
	if l > 0 {
		time.Sleep(time.Duration(l) * time.Millisecond)
	}
	rcvr.Pc.currentSource.NewPixelsAll()
}

func (rcvr digger) reversedir(dir int) int {
	switch dir {
	case 0:
		return 4
		fallthrough
	case 4:
		return 0
		fallthrough
	case 2:
		return 6
		fallthrough
	case 6:
		return 2
	}
	return dir
}

func (rcvr digger) Run() {
	rcvr.Main.main()
}

func (rcvr digger) Start() {
	rcvr.Control.SetCanFocus(true)
	rcvr.Control.GrabFocus()
}

func (q digger) OnDrawn(da *gtk.DrawingArea, g *cairo.Context) bool {
	g.Scale(4, 4)

	var w = q.Pc.width
	var h = q.Pc.height
	var data = q.Pc.pixels

	shift := 1

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			arrayIndex := y*w + x
			c := q.Pc.currentSource.model.GetColor(data[arrayIndex])
			g.SetSourceRGB(c[0], c[1], c[2])
			g.Rectangle(float64(x+shift), float64(y+shift), 1, 1)
			g.Fill()
		}
	}

	return false
}

func (rcvr *digger) updatedigger() {
	var dir int
	var ddir int
	var clbits int
	var diggerox int
	var diggeroy int
	var nmon int
	push := false
	rcvr.Input.readdir()
	dir = rcvr.Input.getdir()
	if dir == 0 || dir == 2 || dir == 4 || dir == 6 {
		ddir = dir
	} else {
		ddir = -1
	}
	if rcvr.diggerrx == 0 && (ddir == 2 || ddir == 6) {
		rcvr.digdir = ddir
		rcvr.digmdir = ddir
	}
	if rcvr.diggerry == 0 && (ddir == 4 || ddir == 0) {
		rcvr.digdir = ddir
		rcvr.digmdir = ddir
	}
	if dir == -1 {
		rcvr.digmdir = -1
	} else {
		rcvr.digmdir = rcvr.digdir
	}
	if rcvr.diggerx == 292 && rcvr.digmdir == 0 || rcvr.diggerx == 12 && rcvr.digmdir == 4 || rcvr.diggery == 180 && rcvr.digmdir == 6 || rcvr.diggery == 18 && rcvr.digmdir == 2 {
		rcvr.digmdir = -1
	}
	diggerox = rcvr.diggerx
	diggeroy = rcvr.diggery
	if rcvr.digmdir != -1 {
		rcvr.Drawing.eatfield(diggerox, diggeroy, rcvr.digmdir)
	}
	switch rcvr.digmdir {
	case 0:
		rcvr.Drawing.drawrightblob(rcvr.diggerx, rcvr.diggery)
		rcvr.diggerx += 4
	case 4:
		rcvr.Drawing.drawleftblob(rcvr.diggerx, rcvr.diggery)
		rcvr.diggerx -= 4
	case 2:
		rcvr.Drawing.drawtopblob(rcvr.diggerx, rcvr.diggery)
		rcvr.diggery -= 3
	case 6:
		rcvr.Drawing.drawbottomblob(rcvr.diggerx, rcvr.diggery)
		rcvr.diggery += 3
	}
	if rcvr.hitemerald((rcvr.diggerx-12)/20, (rcvr.diggery-18)/18, (rcvr.diggerx-12)%20, (rcvr.diggery-18)%18, rcvr.digmdir) {
		rcvr.Scores.scoreemerald()
		rcvr.Sound.soundem()
		rcvr.Sound.soundemerald(rcvr.emocttime)
		rcvr.newSound.playEatEmerald(rcvr.emocttime <= 0)
		rcvr.emocttime = 9
	}
	clbits = rcvr.Drawing.drawdigger(rcvr.digdir, rcvr.diggerx, rcvr.diggery, rcvr.notfiring && rcvr.rechargetime == 0)
	rcvr.Main.incpenalty()
	if rcvr.Bags.bagbits()&clbits != 0 {
		if rcvr.digmdir == 0 || rcvr.digmdir == 4 {
			push = rcvr.Bags.pushbags(rcvr.digmdir, clbits)
			rcvr.digtime++
		} else if !rcvr.Bags.pushudbags(clbits) {
			push = false
		}
		if !push {
			rcvr.diggerx = diggerox
			rcvr.diggery = diggeroy
			rcvr.Drawing.drawdigger(rcvr.digmdir, rcvr.diggerx, rcvr.diggery, rcvr.notfiring && rcvr.rechargetime == 0)
			rcvr.Main.incpenalty()
			rcvr.digdir = rcvr.reversedir(rcvr.digmdir)
		}
	}
	if clbits&0x3f00 != 0 && rcvr.bonusmode {
		for nmon = rcvr.Monster.killmonsters(clbits); nmon != 0; nmon-- {
			rcvr.Sound.soundeatm()
			rcvr.Scores.scoreeatm()
		}
	}
	if clbits&0x4000 != 0 {
		rcvr.Scores.scorebonus()
		rcvr.initbonusmode()
	}
	rcvr.diggerh = (rcvr.diggerx - 12) / 20
	rcvr.diggerrx = (rcvr.diggerx - 12) % 20
	rcvr.diggerv = (rcvr.diggery - 18) / 18
	rcvr.diggerry = (rcvr.diggery - 18) % 18
}

func (rcvr *digger) updatefire() {
	var clbits int
	var b int
	var mon int
	pix := 0
	if rcvr.notfiring {
		if rcvr.rechargetime != 0 {
			rcvr.rechargetime--
		} else if rcvr.getfirepflag() {
			if rcvr.digonscr {
				rcvr.rechargetime = rcvr.Main.levof10()*3 + 60
				rcvr.notfiring = false
				switch rcvr.digdir {
				case 0:
					rcvr.firex = rcvr.diggerx + 8
					rcvr.firey = rcvr.diggery + 4
				case 4:
					rcvr.firex = rcvr.diggerx
					rcvr.firey = rcvr.diggery + 4
				case 2:
					rcvr.firex = rcvr.diggerx + 4
					rcvr.firey = rcvr.diggery
				case 6:
					rcvr.firex = rcvr.diggerx + 4
					rcvr.firey = rcvr.diggery + 8
				}
				rcvr.firedir = rcvr.digdir
				rcvr.Sprite.movedrawspr(15, rcvr.firex, rcvr.firey)
				rcvr.Sound.soundfire()
				rcvr.newSound.fireStart()
			}
		}
	} else {
		switch rcvr.firedir {
		case 0:
			rcvr.firex += 8
			pix = rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+4) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+4)
		case 4:
			rcvr.firex -= 8
			pix = rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+4) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+4)
		case 2:
			rcvr.firey -= 7
			pix = (rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+1) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+2) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+3) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+4) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+5) | rcvr.Pc.ggetpix(rcvr.firex+4, rcvr.firey+6)) & 0xc0
		case 6:
			rcvr.firey += 7
			pix = (rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+1) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+2) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+3) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+4) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+5) | rcvr.Pc.ggetpix(rcvr.firex, rcvr.firey+6)) & 3
		}
		clbits = rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, 0)
		rcvr.Main.incpenalty()
		if clbits&0x3f00 != 0 {
			b = 256
			for mon = 0; mon < 6; mon++ {
				if (clbits & b) != 0 {
					rcvr.Monster.killmon(mon)
					rcvr.Scores.scorekill()
					rcvr.expsn = 1
				}
				b <<= 1
			}
		}
		if clbits&0x40fe != 0 {
			rcvr.expsn = 1
		}
		switch rcvr.firedir {
		case 0:
			if rcvr.firex > 296 {
				rcvr.expsn = 1
			} else if pix != 0 && clbits == 0 {
				rcvr.expsn = 1
				rcvr.firex -= 8
				rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, 0)
			}
		case 4:
			if rcvr.firex < 16 {
				rcvr.expsn = 1
			} else if pix != 0 && clbits == 0 {
				rcvr.expsn = 1
				rcvr.firex += 8
				rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, 0)
			}
		case 2:
			if rcvr.firey < 15 {
				rcvr.expsn = 1
			} else if pix != 0 && clbits == 0 {
				rcvr.expsn = 1
				rcvr.firey += 7
				rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, 0)
			}
		case 6:
			if rcvr.firey > 183 {
				rcvr.expsn = 1
			} else if pix != 0 && clbits == 0 {
				rcvr.expsn = 1
				rcvr.firey -= 7
				rcvr.Drawing.drawfire(rcvr.firex, rcvr.firey, 0)
			}
		}
	}
}

func (d *digger) OnKeyPress(win *gtk.Window, ev *gdk.Event) {
	keyEvent := gdk.EventKey{ev}
	num := ConvertToLegacy(keyEvent)
	if num >= 0 {
		d.KeyDown(num)
	}
}

func (d *digger) OnKeyRelease(win *gtk.Window, ev *gdk.Event) {
	keyEvent := gdk.EventKey{ev}
	num := ConvertToLegacy(keyEvent)
	if num >= 0 {
		d.KeyUp(num)
	}
}
