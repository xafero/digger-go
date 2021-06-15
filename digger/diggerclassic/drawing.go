package diggerclassic

type Drawing struct {
	dig        *digger
	field1     []int
	field2     []int
	field      []int
	diggerbuf  []int16
	bagbuf1    []int16
	bagbuf2    []int16
	bagbuf3    []int16
	bagbuf4    []int16
	bagbuf5    []int16
	bagbuf6    []int16
	bagbuf7    []int16
	monbuf1    []int16
	monbuf2    []int16
	monbuf3    []int16
	monbuf4    []int16
	monbuf5    []int16
	monbuf6    []int16
	bonusbuf   []int16
	firebuf    []int16
	bitmasks   []int
	monspr     []int
	monspd     []int
	digspr     int
	digspd     int
	firespr    int
	fireheight int
}

func NewDrawing(d *digger) *Drawing {
	rcvr := new(Drawing)

	rcvr.field1 = []int{ // [150]
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
	rcvr.field2 = []int{ // [150]
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
	rcvr.field = []int{ // [150]
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

	rcvr.bitmasks = []int{0xfffe, 0xfffd, 0xfffb, 0xfff7, 0xffef, 0xffdf,
		0xffbf, 0xff7f, 0xfeff, 0xfdff, 0xfbff, 0xf7ff} // [12]

	rcvr.monspr = []int{0, 0, 0, 0, 0, 0} // [6]
	rcvr.monspd = []int{0, 0, 0, 0, 0, 0} // [6]

	rcvr.diggerbuf = make([]int16, 480)
	rcvr.bagbuf1 = make([]int16, 480)
	rcvr.bagbuf2 = make([]int16, 480)
	rcvr.bagbuf3 = make([]int16, 480)
	rcvr.bagbuf4 = make([]int16, 480)
	rcvr.bagbuf5 = make([]int16, 480)
	rcvr.bagbuf6 = make([]int16, 480)
	rcvr.bagbuf7 = make([]int16, 480)
	rcvr.monbuf1 = make([]int16, 480)
	rcvr.monbuf2 = make([]int16, 480)
	rcvr.monbuf3 = make([]int16, 480)
	rcvr.monbuf4 = make([]int16, 480)
	rcvr.monbuf5 = make([]int16, 480)
	rcvr.monbuf6 = make([]int16, 480)
	rcvr.bonusbuf = make([]int16, 480)
	rcvr.firebuf = make([]int16, 128)
	rcvr.digspr = 0
	rcvr.digspd = 0
	rcvr.firespr = 0
	rcvr.fireheight = 8
	rcvr.dig = d
	return rcvr
}

func (rcvr *Drawing) createdbfspr() {
	rcvr.digspd = 1
	rcvr.digspr = 0
	rcvr.firespr = 0
	rcvr.dig.Sprite.createspr(0, 0, rcvr.diggerbuf, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(14, 81, rcvr.bonusbuf, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(15, 82, rcvr.firebuf, 2, rcvr.fireheight, 0, 0)
}

func (rcvr *Drawing) creatembspr() {
	var i int
	rcvr.dig.Sprite.createspr(1, 62, rcvr.bagbuf1, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(2, 62, rcvr.bagbuf2, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(3, 62, rcvr.bagbuf3, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(4, 62, rcvr.bagbuf4, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(5, 62, rcvr.bagbuf5, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(6, 62, rcvr.bagbuf6, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(7, 62, rcvr.bagbuf7, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(8, 71, rcvr.monbuf1, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(9, 71, rcvr.monbuf2, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(10, 71, rcvr.monbuf3, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(11, 71, rcvr.monbuf4, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(12, 71, rcvr.monbuf5, 4, 15, 0, 0)
	rcvr.dig.Sprite.createspr(13, 71, rcvr.monbuf6, 4, 15, 0, 0)
	rcvr.createdbfspr()
	for i = 0; i < 6; i++ {
		rcvr.monspr[i] = 0
		rcvr.monspd[i] = 1
	}
}

func (rcvr *Drawing) drawbackg(l int) {
	var x int
	var y int
	for y = 14; y < 200; y += 4 {
		for x = 0; x < 320; x += 20 {
			rcvr.dig.Sprite.drawmiscspr(x, y, 93+l, 5, 4)
		}
	}
}

func (rcvr *Drawing) drawbonus(x int, y int) {
	rcvr.dig.Sprite.initspr(14, 81, 4, 15, 0, 0)
	rcvr.dig.Sprite.movedrawspr(14, x, y)
}

func (rcvr *Drawing) drawbottomblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x-4, y+15, 6, 6)
	rcvr.dig.Sprite.drawmiscspr(x-4, y+15, 105, 6, 6)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) drawdigger(t int, x int, y int, f bool) int {
	rcvr.digspr += rcvr.digspd
	if rcvr.digspr == 2 || rcvr.digspr == 0 {
		rcvr.digspd = -rcvr.digspd
	}
	if rcvr.digspr > 2 {
		rcvr.digspr = 2
	}
	if rcvr.digspr < 0 {
		rcvr.digspr = 0
	}
	if t >= 0 && t <= 6 && !(t&1 != 0) {
		var fn int
		if f {
			fn = 0
		} else {
			fn = 1
		}
		rcvr.dig.Sprite.initspr(0, (t+(fn))*3+rcvr.digspr+1, 4, 15, 0, 0)
		return rcvr.dig.Sprite.drawspr(0, x, y)
	}
	if t >= 10 && t <= 15 {
		rcvr.dig.Sprite.initspr(0, 40-t, 4, 15, 0, 0)
		return rcvr.dig.Sprite.drawspr(0, x, y)
	}
	return 0
}

func (rcvr *Drawing) drawemerald(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x, y, 4, 10)
	rcvr.dig.Sprite.drawmiscspr(x, y, 108, 4, 10)
	rcvr.dig.Sprite.getis()
}

func (q *Drawing) drawfield() {
	var x int
	var y int
	var xp int
	var yp int
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if (q.field[y*15+x] & 0x2000) == 0 {
				xp = x*20 + 12
				yp = y*18 + 18
				if (q.field[y*15+x] & 0xfc0) != 0xfc0 {
					q.field[y*15+x] &= 0xd03f
					q.drawbottomblob(xp, yp-15)
					q.drawbottomblob(xp, yp-12)
					q.drawbottomblob(xp, yp-9)
					q.drawbottomblob(xp, yp-6)
					q.drawbottomblob(xp, yp-3)
					q.drawtopblob(xp, yp+3)
				}
				if (q.field[y*15+x] & 0x1f) != 0x1f {
					q.field[y*15+x] &= 0xdfe0
					q.drawrightblob(xp-16, yp)
					q.drawrightblob(xp-12, yp)
					q.drawrightblob(xp-8, yp)
					q.drawrightblob(xp-4, yp)
					q.drawleftblob(xp+4, yp)
				}
				if x < 14 {
					if (q.field[y*15+x+1] & 0xfdf) != 0xfdf {
						q.drawrightblob(xp, yp)
					}
				}
				if y < 9 {
					if (q.field[(y+1)*15+x] & 0xfdf) != 0xfdf {
						q.drawbottomblob(xp, yp)
					}
				}
			}
		}
	}
}

func (rcvr *Drawing) drawfire(x int, y int, t int) int {
	if t == 0 {
		rcvr.firespr++
		if rcvr.firespr > 2 {
			rcvr.firespr = 0
		}
		rcvr.dig.Sprite.initspr(15, 82+rcvr.firespr, 2, rcvr.fireheight, 0, 0)
	} else {
		rcvr.dig.Sprite.initspr(15, 84+t, 2, rcvr.fireheight, 0, 0)
	}
	return rcvr.dig.Sprite.drawspr(15, x, y)
}

func (rcvr *Drawing) drawfurryblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x-4, y+15, 6, 8)
	rcvr.dig.Sprite.drawmiscspr(x-4, y+15, 107, 6, 8)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) drawgold(n int, t int, x int, y int) int {
	rcvr.dig.Sprite.initspr(n, t+62, 4, 15, 0, 0)
	return rcvr.dig.Sprite.drawspr(n, x, y)
}

func (rcvr *Drawing) drawleftblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x-8, y-1, 2, 18)
	rcvr.dig.Sprite.drawmiscspr(x-8, y-1, 104, 2, 18)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) drawlife(t int, x int, y int) {
	rcvr.dig.Sprite.drawmiscspr(x, y, t+110, 4, 12)
}

func (q *Drawing) drawlives() {
	var l int
	var n int
	n = q.dig.Main.getlives(1) - 1
	for l = 1; l < 5; l++ {
		var nna int
		if n > 0 {
			nna = 0
		} else {
			nna = 2
		}
		q.drawlife(nna, l*20+60, 0)
		n--
	}
	if q.dig.Main.nplayers == 2 {
		n = q.dig.Main.getlives(2) - 1
		for l = 1; l < 5; l++ {
			var nnb int
			if n > 0 {
				nnb = 1
			} else {
				nnb = 2
			}
			q.drawlife(nnb, 244-l*20, 0)
			n--
		}
	}
}

func (q *Drawing) drawmon(n int, nobf bool, dir int, x int, y int) int {
	q.monspr[n] += q.monspd[n]
	if q.monspr[n] == 2 || q.monspr[n] == 0 {
		q.monspd[n] = -q.monspd[n]
	}
	if q.monspr[n] > 2 {
		q.monspr[n] = 2
	}
	if q.monspr[n] < 0 {
		q.monspr[n] = 0
	}
	if nobf {
		q.dig.Sprite.initspr(n+8, q.monspr[n]+69, 4, 15, 0, 0)
	} else {
		switch dir {
		case 0:
			q.dig.Sprite.initspr(n+8, q.monspr[n]+73, 4, 15, 0, 0)
			break
		case 4:
			q.dig.Sprite.initspr(n+8, q.monspr[n]+77, 4, 15, 0, 0)
			break
		}
	}
	return q.dig.Sprite.drawspr(n+8, x, y)
}

func (rcvr *Drawing) drawmondie(n int, nobf bool, dir int, x int, y int) int {
	if nobf {
		rcvr.dig.Sprite.initspr(n+8, 72, 4, 15, 0, 0)
	} else {
		switch dir {
		case 0:
			rcvr.dig.Sprite.initspr(n+8, 76, 4, 15, 0, 0)
			break
		case 4:
			rcvr.dig.Sprite.initspr(n+8, 80, 4, 14, 0, 0)
			break
		}
	}
	return rcvr.dig.Sprite.drawspr(n+8, x, y)
}

func (rcvr *Drawing) drawrightblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x+16, y-1, 2, 18)
	rcvr.dig.Sprite.drawmiscspr(x+16, y-1, 102, 2, 18)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) drawsquareblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x-4, y+17, 6, 6)
	rcvr.dig.Sprite.drawmiscspr(x-4, y+17, 106, 6, 6)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) drawstatics() {
	var x int
	var y int
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if rcvr.dig.Main.getcplayer() == 0 {
				rcvr.field[y*15+x] = rcvr.field1[y*15+x]
			} else {
				rcvr.field[y*15+x] = rcvr.field2[y*15+x]
			}
		}
	}
	rcvr.dig.Sprite.setretr(true)
	rcvr.dig.Pc.gpal(0)
	rcvr.dig.Pc.ginten(0)
	rcvr.drawbackg(rcvr.dig.Main.levplan())
	rcvr.drawfield()
	rcvr.dig.Pc.currentSource.NewPixels(0, 0, rcvr.dig.Pc.width, rcvr.dig.Pc.height)
}

func (rcvr *Drawing) drawtopblob(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x-4, y-6, 6, 6)
	rcvr.dig.Sprite.drawmiscspr(x-4, y-6, 103, 6, 6)
	rcvr.dig.Sprite.getis()
}

func (q *Drawing) eatfield(x int, y int, dir int) {
	h := (x - 12) / 20
	xr := (x - 12) % 20 / 4
	v := (y - 18) / 18
	yr := (y - 18) % 18 / 3
	q.dig.Main.incpenalty()
	switch dir {
	case 0:
		h++
		q.field[v*15+h] &= q.bitmasks[xr]
		if q.field[v*15+h]&0x1f != 0 {
			break
		}
		q.field[v*15+h] &= 0xdfff
	case 4:
		xr--
		if xr < 0 {
			xr += 5
			h--
		}
		q.field[v*15+h] &= q.bitmasks[xr]
		if q.field[v*15+h]&0x1f != 0 {
			break
		}
		q.field[v*15+h] &= 0xdfff
	case 2:
		yr--
		if yr < 0 {
			yr += 6
			v--
		}
		q.field[v*15+h] &= q.bitmasks[6+yr]
		if q.field[v*15+h]&0xfc0 != 0 {
			break
		}
		q.field[v*15+h] &= 0xdfff
	case 6:
		v++
		q.field[v*15+h] &= q.bitmasks[6+yr]
		if q.field[v*15+h]&0xfc0 != 0 {
			break
		}
		q.field[v*15+h] &= 0xdfff
	}
}

func (rcvr *Drawing) eraseemerald(x int, y int) {
	rcvr.dig.Sprite.initmiscspr(x, y, 4, 10)
	rcvr.dig.Sprite.drawmiscspr(x, y, 109, 4, 10)
	rcvr.dig.Sprite.getis()
}

func (rcvr *Drawing) initdbfspr() {
	rcvr.digspd = 1
	rcvr.digspr = 0
	rcvr.firespr = 0
	rcvr.dig.Sprite.initspr(0, 0, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(14, 81, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(15, 82, 2, rcvr.fireheight, 0, 0)
}

func (rcvr *Drawing) initmbspr() {
	rcvr.dig.Sprite.initspr(1, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(2, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(3, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(4, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(5, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(6, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(7, 62, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(8, 71, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(9, 71, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(10, 71, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(11, 71, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(12, 71, 4, 15, 0, 0)
	rcvr.dig.Sprite.initspr(13, 71, 4, 15, 0, 0)
	rcvr.initdbfspr()
}

func (q *Drawing) makefield() {
	var c int
	var x int
	var y int
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			q.field[y*15+x] = -1
			c = q.dig.Main.getlevch(x, y, q.dig.Main.levplan())
			if c == 'S' || c == 'V' {
				q.field[y*15+x] &= 0xd03f
			}
			if c == 'S' || c == 'H' {
				q.field[y*15+x] &= 0xdfe0
			}
			if q.dig.Main.getcplayer() == 0 {
				q.field1[y*15+x] = q.field[y*15+x]
			} else {
				q.field2[y*15+x] = q.field[y*15+x]
			}
		}
	}
}

func (rcvr *Drawing) outtext(p string, x int, y int, c int, b bool) {
	var i int
	rx := x
	for i = 0; i < len(p); i++ {
		rcvr.dig.Pc.gwrite2(x, y, int(p[i]), c)
		x += 12
	}
	if b {
		rcvr.dig.Pc.currentSource.NewPixels(rx, y, len(p)*12, 12)
	}
}

func (rcvr *Drawing) outtext2(p string, x int, y int, c int) {
	rcvr.outtext(p, x, y, c, false)
}

func (q *Drawing) savefield() {
	var x int
	var y int
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if q.dig.Main.getcplayer() == 0 {
				q.field1[y*15+x] = q.field[y*15+x]
			} else {
				q.field2[y*15+x] = q.field[y*15+x]
			}
		}
	}
}
