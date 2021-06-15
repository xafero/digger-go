package diggerclassic

import "time"

type Main struct {
	dig         *digger
	digsprorder []int
	gamedat     []game
	pldispbuf   string
	curplayer   int
	nplayers    int
	penalty     int
	levnotdrawn bool
	flashplayer bool
	levfflag    bool
	biosflag    bool
	speedmul    int
	delaytime   int
	randv       int
	leveldat    [][]string
}

func NewMain(d *digger) *Main {
	rcvr := new(Main)

	rcvr.digsprorder = []int{14, 13, 7, 6, 5, 4, 3, 2, 1, 12, 11, 10, 9, 8, 15, 0} // [16]
	rcvr.gamedat = []game{NewGame(), NewGame()}

	rcvr.leveldat = [][]string{
		{"S   B     HHHHS",
			"V  CC  C  V B  ",
			"VB CC  C  V    ",
			"V  CCB CB V CCC",
			"V  CC  C  V CCC",
			"HH CC  C  V CCC",
			" V    B B V    ",
			" HHHH     V    ",
			"C   V     V   C",
			"CC  HHHHHHH  CC"},
		{"SHHHHH  B B  HS",
			" CC  V       V ",
			" CC  V CCCCC V ",
			"BCCB V CCCCC V ",
			"CCCC V       V ",
			"CCCC V B  HHHH ",
			" CC  V CC V    ",
			" BB  VCCCCV CC ",
			"C    V CC V CC ",
			"CC   HHHHHH    "},
		{"SHHHHB B BHHHHS",
			"CC  V C C V BB ",
			"C   V C C V CC ",
			" BB V C C VCCCC",
			"CCCCV C C VCCCC",
			"CCCCHHHHHHH CC ",
			" CC  C V C  CC ",
			" CC  C V C     ",
			"C    C V C    C",
			"CC   C H C   CC"},
		{"SHBCCCCBCCCCBHS",
			"CV  CCCCCCC  VC",
			"CHHH CCCCC HHHC",
			"C  V  CCC  V  C",
			"   HHH C HHH   ",
			"  B  V B V  B  ",
			"  C  VCCCV  C  ",
			" CCC HHHHH CCC ",
			"CCCCC CVC CCCCC",
			"CCCCC CHC CCCCC"},
		{"SHHHHHHHHHHHHHS",
			"VBCCCCBVCCCCCCV",
			"VCCCCCCV CCBC V",
			"V CCCC VCCBCCCV",
			"VCCCCCCV CCCC V",
			"V CCCC VBCCCCCV",
			"VCCBCCCV CCCC V",
			"V CCBC VCCCCCCV",
			"VCCCCCCVCCCCCCV",
			"HHHHHHHHHHHHHHH"},
		{"SHHHHHHHHHHHHHS",
			"VCBCCV V VCCBCV",
			"VCCC VBVBV CCCV",
			"VCCCHH V HHCCCV",
			"VCC V CVC V CCV",
			"VCCHH CVC HHCCV",
			"VC V CCVCC V CV",
			"VCHHBCCVCCBHHCV",
			"VCVCCCCVCCCCVCV",
			"HHHHHHHHHHHHHHH"},
		{"SHCCCCCVCCCCCHS",
			" VCBCBCVCBCBCV ",
			"BVCCCCCVCCCCCVB",
			"CHHCCCCVCCCCHHC",
			"CCV CCCVCCC VCC",
			"CCHHHCCVCCHHHCC",
			"CCCCV CVC VCCCC",
			"CCCCHH V HHCCCC",
			"CCCCCV V VCCCCC",
			"CCCCCHHHHHCCCCC"},
		{"HHHHHHHHHHHHHHS",
			"V CCBCCCCCBCC V",
			"HHHCCCCBCCCCHHH",
			"VBV CCCCCCC VBV",
			"VCHHHCCCCCHHHCV",
			"VCCBV CCC VBCCV",
			"VCCCHHHCHHHCCCV",
			"VCCCC V V CCCCV",
			"VCCCCCV VCCCCCV",
			"HHHHHHHHHHHHHHH"},
	}

	rcvr.pldispbuf = ""
	rcvr.curplayer = 0
	rcvr.nplayers = 0
	rcvr.penalty = 0
	rcvr.levnotdrawn = false
	rcvr.flashplayer = false
	rcvr.levfflag = false
	rcvr.biosflag = false
	rcvr.speedmul = 40
	rcvr.delaytime = 0
	rcvr.dig = d
	return rcvr
}

func (rcvr *Main) addlife(pl int) {
	rcvr.gamedat[pl-1].lives++
	rcvr.dig.Sound.sound1up()
}

func (rcvr *Main) calibrate() {
	rcvr.dig.Sound.volume = (rcvr.dig.Pc.getkips() / 291)
	if rcvr.dig.Sound.volume == 0 {
		rcvr.dig.Sound.volume = 1
	}
}

func (rcvr *Main) checklevdone() {
	if (rcvr.dig.countem() == 0 || rcvr.dig.Monster.monleft() == 0) && rcvr.dig.digonscr {
		rcvr.gamedat[rcvr.curplayer].levdone = true
	} else {
		rcvr.gamedat[rcvr.curplayer].levdone = false
	}
}

func (rcvr *Main) cleartopline() {
	rcvr.dig.Drawing.outtext2("                          ", 0, 0, 3)
	rcvr.dig.Drawing.outtext2(" ", 308, 0, 3)
}

func (rcvr *Main) drawscreen() {
	rcvr.dig.Drawing.creatembspr()
	rcvr.dig.Drawing.drawstatics()
	rcvr.dig.Bags.drawbags()
	rcvr.dig.drawemeralds()
	rcvr.dig.initdigger()
	rcvr.dig.Monster.initmonsters()
}

func (rcvr *Main) getcplayer() int {
	return rcvr.curplayer
}

func (rcvr *Main) getlevch(x int, y int, l int) int {
	if l == 0 {
		l++
	}
	return int(rcvr.leveldat[l-1][y][x])
}

func (rcvr *Main) getlives(pl int) int {
	return rcvr.gamedat[pl-1].lives
}

func (rcvr *Main) incpenalty() {
	rcvr.penalty++
}

func (rcvr *Main) initchars() {
	rcvr.dig.Drawing.initmbspr()
	rcvr.dig.initdigger()
	rcvr.dig.Monster.initmonsters()
}

func (rcvr *Main) initlevel() {
	rcvr.gamedat[rcvr.curplayer].levdone = false
	rcvr.dig.Drawing.makefield()
	rcvr.dig.makeemfield()
	rcvr.dig.Bags.initbags()
	rcvr.levnotdrawn = true
}

func (q *Main) levno() int {
	return q.gamedat[q.curplayer].level
}

func (q *Main) levof10() int {
	if q.gamedat[q.curplayer].level > 10 {
		return 10
	}
	return q.gamedat[q.curplayer].level
}

func (rcvr *Main) levplan() int {
	l := rcvr.levno()
	if l > 8 {
		l = l&3 + 5
	}
	return l
}

func (rcvr *Main) main() {
	var frame int
	var t int
	x := 0
	var start bool
	rcvr.randv = int(rcvr.dig.Pc.gethrt())
	rcvr.calibrate()
	rcvr.dig.ftime = int64(rcvr.speedmul * 2000)
	rcvr.dig.Sprite.setretr(false)
	rcvr.dig.Pc.ginit()
	rcvr.dig.Sprite.setretr(true)
	rcvr.dig.Pc.gpal(0)
	rcvr.dig.Input.initkeyb()
	rcvr.dig.Input.detectjoy()
	rcvr.dig.Scores.loadscores()
	rcvr.dig.Sound.initsound()
	rcvr.dig.Scores.Init()
	rcvr.dig.Scores._updatescores(rcvr.dig.Scores.scores)
	rcvr.nplayers = 1
	for {
		rcvr.dig.Sound.soundstop()
		rcvr.dig.newSound.killAll()
		rcvr.dig.Sprite.setsprorder(rcvr.digsprorder)
		rcvr.dig.Drawing.creatembspr()
		rcvr.dig.Input.detectjoy()
		rcvr.dig.Pc.gclear()
		rcvr.dig.Pc.gtitle()
		rcvr.dig.Drawing.outtext2("D I G G E R", 100, 0, 3)
		rcvr.shownplayers()
		rcvr.dig.Scores.showtable()
		start = false
		frame = 0
		rcvr.dig.time = rcvr.dig.Pc.gethrt()
		for !start {
			start = rcvr.dig.Input.teststart()
			if rcvr.dig.Input.akeypressed == 27 {
				rcvr.switchnplayers()
				rcvr.shownplayers()
				rcvr.dig.Input.akeypressed = 0
				rcvr.dig.Input.keypressed = 0
			}
			if frame == 0 {
				for t = 54; t < 174; t += 12 {
					rcvr.dig.Drawing.outtext2("            ", 164, t, 0)
				}
			}
			if frame == 50 {
				rcvr.dig.Sprite.movedrawspr(8, 292, 63)
				x = 292
			}
			if frame > 50 && frame <= 77 {
				x -= 4
				rcvr.dig.Drawing.drawmon(0, true, 4, x, 63)
			}
			if frame > 77 {
				rcvr.dig.Drawing.drawmon(0, true, 0, 184, 63)
			}
			if frame == 83 {
				rcvr.dig.Drawing.outtext2("NOBBIN", 216, 64, 2)
			}
			if frame == 90 {
				rcvr.dig.Sprite.movedrawspr(9, 292, 82)
				rcvr.dig.Drawing.drawmon(1, false, 4, 292, 82)
				x = 292
			}
			if frame > 90 && frame <= 117 {
				x -= 4
				rcvr.dig.Drawing.drawmon(1, false, 4, x, 82)
			}
			if frame > 117 {
				rcvr.dig.Drawing.drawmon(1, false, 0, 184, 82)
			}
			if frame == 123 {
				rcvr.dig.Drawing.outtext2("HOBBIN", 216, 83, 2)
			}
			if frame == 130 {
				rcvr.dig.Sprite.movedrawspr(0, 292, 101)
				rcvr.dig.Drawing.drawdigger(4, 292, 101, true)
				x = 292
			}
			if frame > 130 && frame <= 157 {
				x -= 4
				rcvr.dig.Drawing.drawdigger(4, x, 101, true)
			}
			if frame > 157 {
				rcvr.dig.Drawing.drawdigger(0, 184, 101, true)
			}
			if frame == 163 {
				rcvr.dig.Drawing.outtext2("DIGGER", 216, 102, 2)
			}
			if frame == 178 {
				rcvr.dig.Sprite.movedrawspr(1, 184, 120)
				rcvr.dig.Drawing.drawgold(1, 0, 184, 120)
			}
			if frame == 183 {
				rcvr.dig.Drawing.outtext2("GOLD", 216, 121, 2)
			}
			if frame == 198 {
				rcvr.dig.Drawing.drawemerald(184, 141)
			}
			if frame == 203 {
				rcvr.dig.Drawing.outtext2("EMERALD", 216, 140, 2)
			}
			if frame == 218 {
				rcvr.dig.Drawing.drawbonus(184, 158)
			}
			if frame == 223 {
				rcvr.dig.Drawing.outtext2("BONUS", 216, 159, 2)
			}
			rcvr.dig.newframe()
			frame++
			if frame > 250 {
				frame = 0
			}
		}
		rcvr.gamedat[0].level = 1
		rcvr.gamedat[0].lives = 3
		if rcvr.nplayers == 2 {
			rcvr.gamedat[1].level = 1
			rcvr.gamedat[1].lives = 3
		} else {
			rcvr.gamedat[1].lives = 0
		}
		rcvr.dig.Pc.gclear()
		rcvr.curplayer = 0
		rcvr.initlevel()
		rcvr.curplayer = 1
		rcvr.initlevel()
		rcvr.dig.Scores.zeroscores()
		rcvr.dig.bonusvisible = true
		if rcvr.nplayers == 2 {
			rcvr.flashplayer = true
		}
		rcvr.curplayer = 0
		for (rcvr.gamedat[0].lives != 0 || rcvr.gamedat[1].lives != 0) && !rcvr.dig.Input.escape {
			rcvr.gamedat[rcvr.curplayer].dead = false
			for !rcvr.gamedat[rcvr.curplayer].dead && rcvr.gamedat[rcvr.curplayer].lives != 0 && !rcvr.dig.Input.escape {
				rcvr.dig.Drawing.initmbspr()
				rcvr.play()
			}
			if rcvr.gamedat[1-rcvr.curplayer].lives != 0 {
				rcvr.curplayer = 1 - rcvr.curplayer
				rcvr.flashplayer = true
				rcvr.levnotdrawn = true
			}
		}
		rcvr.dig.Input.escape = false
		rcvr.dig.newSound.killAll()
		if !(!false) {
			break
		}
	}
}

func (rcvr *Main) play() {
	var t int
	var c int
	if rcvr.levnotdrawn {
		rcvr.levnotdrawn = false
		rcvr.drawscreen()
		rcvr.dig.time = rcvr.dig.Pc.gethrt()
		if rcvr.flashplayer {
			rcvr.flashplayer = false
			rcvr.pldispbuf = "PLAYER "
			if rcvr.curplayer == 0 {
				rcvr.pldispbuf += "1"
			} else {
				rcvr.pldispbuf += "2"
			}
			rcvr.cleartopline()
			for t = 0; t < 15; t++ {
				for c = 1; c <= 3; c++ {
					rcvr.dig.Drawing.outtext2(rcvr.pldispbuf, 108, 0, c)
					rcvr.dig.Scores.writecurscore(c)
					/* olddelay(20); */
					rcvr.dig.newframe()
					if rcvr.dig.Input.escape {
						return
					}
				}
			}
			rcvr.dig.Scores.drawscores()
			rcvr.dig.Scores.addscore(0)
		}
	} else {
		rcvr.initchars()
	}
	rcvr.dig.Input.keypressed = 0
	rcvr.dig.Drawing.outtext2("        ", 108, 0, 3)
	rcvr.dig.Scores.initscores()
	rcvr.dig.Drawing.drawlives()
	rcvr.dig.Sound.music(1)
	rcvr.dig.newSound.startNormalBackgroundMusic()
	rcvr.dig.Input.readdir()
	rcvr.dig.time = rcvr.dig.Pc.gethrt()
	for !rcvr.gamedat[rcvr.curplayer].dead && !rcvr.gamedat[rcvr.curplayer].levdone && !rcvr.dig.Input.escape {
		rcvr.penalty = 0
		rcvr.dig.dodigger()
		rcvr.dig.Monster.domonsters()
		rcvr.dig.Bags.dobags()
		if rcvr.penalty > 8 {
			rcvr.dig.Monster.incmont(rcvr.penalty - 8)
		}
		rcvr.testpause()
		rcvr.checklevdone()
	}
	rcvr.dig.erasedigger()
	rcvr.dig.Sound.musicoff()
	t = 20
	for (rcvr.dig.Bags.getnmovingbags() != 0 || t != 0) && !rcvr.dig.Input.escape {
		if t != 0 {
			t--
		}
		rcvr.penalty = 0
		rcvr.dig.Bags.dobags()
		rcvr.dig.dodigger()
		rcvr.dig.Monster.domonsters()
		if rcvr.penalty < 8 {
			t = 0
		}
	}
	rcvr.dig.Sound.soundstop()
	rcvr.dig.newSound.killAll()
	rcvr.dig.killfire()
	rcvr.dig.erasebonus()
	rcvr.dig.Bags.cleanupbags()
	rcvr.dig.Drawing.savefield()
	rcvr.dig.Monster.erasemonsters()
	rcvr.dig.newframe()
	if rcvr.gamedat[rcvr.curplayer].levdone {
		rcvr.dig.Sound.soundlevdone()
	}
	if rcvr.dig.countem() == 0 {
		rcvr.gamedat[rcvr.curplayer].level++
		if rcvr.gamedat[rcvr.curplayer].level > 1000 {
			rcvr.gamedat[rcvr.curplayer].level = 1000
		}
		rcvr.initlevel()
	}
	if rcvr.gamedat[rcvr.curplayer].dead {
		rcvr.gamedat[rcvr.curplayer].lives--
		rcvr.dig.Drawing.drawlives()
		if rcvr.gamedat[rcvr.curplayer].lives == 0 && !rcvr.dig.Input.escape {
			rcvr.dig.Scores.endofgame()
		}
	}
	if rcvr.gamedat[rcvr.curplayer].levdone {
		rcvr.gamedat[rcvr.curplayer].level++
		if rcvr.gamedat[rcvr.curplayer].level > 1000 {
			rcvr.gamedat[rcvr.curplayer].level = 1000
		}
		rcvr.initlevel()
	}
}

func (rcvr *Main) randno(n int) int {
	rcvr.randv = rcvr.randv*0x15a4e35 + 1
	return rcvr.randv & 0x7fffffff % n
}

func (rcvr *Main) setdead(bp6 bool) {
	rcvr.gamedat[rcvr.curplayer].dead = bp6
}

func (rcvr *Main) shownplayers() {
	if rcvr.nplayers == 1 {
		rcvr.dig.Drawing.outtext2("ONE", 220, 25, 3)
		rcvr.dig.Drawing.outtext2(" PLAYER ", 192, 39, 3)
	} else {
		rcvr.dig.Drawing.outtext2("TWO", 220, 25, 3)
		rcvr.dig.Drawing.outtext2(" PLAYERS", 184, 39, 3)
	}
}

func (rcvr *Main) switchnplayers() {
	rcvr.nplayers = 3 - rcvr.nplayers
}

func (rcvr *Main) testpause() {
	if rcvr.dig.Input.akeypressed == 32 { /* Space bar */
		rcvr.dig.Input.akeypressed = 0
		rcvr.dig.Sound.soundpause()
		rcvr.dig.Sound.sett2val(40)
		rcvr.dig.Sound.setsoundt2()
		rcvr.cleartopline()
		rcvr.dig.Drawing.outtext2("PRESS ANY KEY", 80, 0, 1)
		rcvr.dig.newframe()
		rcvr.dig.Input.keypressed = 0
		for true {
			time.Sleep(time.Duration(50) * time.Millisecond)
			if rcvr.dig.Input.keypressed != 0 {
				break
			}
		}
		rcvr.cleartopline()
		rcvr.dig.Scores.drawscores()
		rcvr.dig.Scores.addscore(0)
		rcvr.dig.Drawing.drawlives()
		rcvr.dig.newframe()
		rcvr.dig.time = rcvr.dig.Pc.gethrt() - int64(rcvr.dig.FrameTime)
		rcvr.dig.Input.keypressed = 0
	} else {
		rcvr.dig.Sound.soundpauseoff()
	}
}
