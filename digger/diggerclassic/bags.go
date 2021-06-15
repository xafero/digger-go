package diggerclassic

type Bags struct {
	dig       *digger
	bagdat1   []bag
	bagdat2   []bag
	bagdat    []bag
	pushcount int
	goldtime  int
	wblanim   []int
}

func NewBags(d *digger) *Bags {
	rcvr := new(Bags)

	rcvr.wblanim = []int{2, 0, 1, 0} // [4]
	rcvr.bagdat1 = []bag{NewBag(), NewBag(), NewBag(), NewBag(),
		NewBag(), NewBag(), NewBag(), NewBag()}
	rcvr.bagdat2 = []bag{NewBag(), NewBag(), NewBag(), NewBag(),
		NewBag(), NewBag(), NewBag(), NewBag()}
	rcvr.bagdat = []bag{NewBag(), NewBag(), NewBag(), NewBag(),
		NewBag(), NewBag(), NewBag(), NewBag()}

	rcvr.pushcount = 0
	rcvr.goldtime = 0
	rcvr.dig = d
	return rcvr
}

func (rcvr *Bags) bagbits() int {
	var bag int
	var b int
	var bagx int
	b = 2
	bagx = 0
	for bag = 1; bag < 8; bag++ {
		if rcvr.bagdat[bag].exist {
			bagx |= b
		}
		b <<= 1
	}
	return bagx
}

func (q *Bags) baghitground(bag int) {
	var bn int
	var b int
	var clbits int
	if q.bagdat[bag].dir == 6 && q.bagdat[bag].fallh > 1 {
		q.bagdat[bag].gt = 1
	} else {
		q.bagdat[bag].fallh = 0
	}
	q.bagdat[bag].dir = -1
	q.bagdat[bag].wt = 15
	q.bagdat[bag].wobbling = false
	clbits = q.dig.Drawing.drawgold(bag, 0, q.bagdat[bag].x, q.bagdat[bag].y)
	q.dig.Main.incpenalty()
	b = 2
	for bn = 1; bn < 8; bn++ {
		if (b & clbits) != 0 {
			q.removebag(bn)
		}
		b <<= 1
	}
}

func (rcvr *Bags) bagy(bag int) int {
	return rcvr.bagdat[bag].y
}

func (q *Bags) cleanupbags() {
	var bpa int
	q.dig.Sound.soundfalloff()
	q.dig.newSound.fallEnd()

	for bpa = 1; bpa < 8; bpa++ {
		if q.bagdat[bpa].exist && ((q.bagdat[bpa].h == 7 && q.bagdat[bpa].v == 9) ||
			q.bagdat[bpa].xr != 0 || q.bagdat[bpa].yr != 0 || q.bagdat[bpa].gt != 0 ||
			q.bagdat[bpa].fallh != 0 || q.bagdat[bpa].wobbling) {
			q.bagdat[bpa].exist = false
			q.dig.Sprite.erasespr(bpa)
		}
		if q.dig.Main.getcplayer() == 0 {
			q.bagdat1[bpa] = q.bagdat[bpa].copyFrom()
		} else {
			q.bagdat2[bpa] = q.bagdat[bpa].copyFrom()
		}
	}
}

func (q *Bags) dobags() {
	var bag int
	soundfalloffflag := true
	soundwobbleoffflag := true
	for bag = 1; bag < 8; bag++ {
		if q.bagdat[bag].exist {
			if q.bagdat[bag].gt != 0 {
				if q.bagdat[bag].gt == 1 {
					q.dig.newSound.playBagBreak()
					q.dig.Sound.soundbreak()
					q.dig.Drawing.drawgold(bag, 4, q.bagdat[bag].x, q.bagdat[bag].y)
					q.dig.Main.incpenalty()
				}
				if q.bagdat[bag].gt == 3 {
					q.dig.Drawing.drawgold(bag, 5, q.bagdat[bag].x, q.bagdat[bag].y)
					q.dig.Main.incpenalty()
				}
				if q.bagdat[bag].gt == 5 {
					q.dig.Drawing.drawgold(bag, 6, q.bagdat[bag].x, q.bagdat[bag].y)
					q.dig.Main.incpenalty()
				}
				q.bagdat[bag].gt++
				if q.bagdat[bag].gt == q.goldtime {
					q.removebag(bag)
				} else if q.bagdat[bag].v < 9 && q.bagdat[bag].gt < q.goldtime-10 {
					if (q.dig.Monster.getfield(q.bagdat[bag].h, q.bagdat[bag].v+1) & 0x2000) == 0 {
						q.bagdat[bag].gt = q.goldtime - 10
					}
				}
			} else {
				q.updatebag(bag)
			}
		}
	}
	for bag = 1; bag < 8; bag++ {
		if q.bagdat[bag].dir == 6 && q.bagdat[bag].exist {
			soundfalloffflag = false
		}
		if q.bagdat[bag].dir != 6 && q.bagdat[bag].wobbling && q.bagdat[bag].exist {
			soundwobbleoffflag = false
		}
	}
	if soundfalloffflag {
		q.dig.Sound.soundfalloff()
		q.dig.newSound.fallEnd()
	}
	if soundwobbleoffflag {
		q.dig.Sound.soundwobbleoff()
	}
}

func (q *Bags) drawbags() {
	var bag int
	for bag = 1; bag < 8; bag++ {
		if q.dig.Main.getcplayer() == 0 {
			q.bagdat[bag] = q.bagdat1[bag].copyFrom()
		} else {
			q.bagdat[bag] = q.bagdat2[bag].copyFrom()
		}
		if q.bagdat[bag].exist {
			q.dig.Sprite.movedrawspr(bag, q.bagdat[bag].x, q.bagdat[bag].y)
		}
	}
}

func (q *Bags) getbagdir(bag int) int {
	if q.bagdat[bag].exist {
		return q.bagdat[bag].dir
	}
	return -1
}

func (q *Bags) getgold(bag int) {
	var clbits int
	clbits = q.dig.Drawing.drawgold(bag, 6, q.bagdat[bag].x, q.bagdat[bag].y)
	q.dig.Main.incpenalty()
	if (clbits & 1) != 0 {
		q.dig.Scores.scoregold()
		q.dig.Sound.soundgold()
		q.dig.digtime = 0
	} else {
		q.dig.Monster.mongold()
	}
	q.removebag(bag)
}

func (q *Bags) getnmovingbags() int {
	var bag int
	n := 0
	for bag = 1; bag < 8; bag++ {
		if q.bagdat[bag].exist && q.bagdat[bag].gt < 10 &&
			(q.bagdat[bag].gt != 0 || q.bagdat[bag].wobbling) {
			n++
		}
	}
	return n
}

func (q *Bags) initbags() {
	var bag int
	var x int
	var y int
	q.pushcount = 0
	q.goldtime = 150 - q.dig.Main.levof10()*10
	for bag = 1; bag < 8; bag++ {
		q.bagdat[bag].exist = false
	}
	bag = 1
	for x = 0; x < 15; x++ {
		for y = 0; y < 10; y++ {
			if q.dig.Main.getlevch(x, y, q.dig.Main.levplan()) == 'B' {
				if bag < 8 {
					q.bagdat[bag].exist = true
					q.bagdat[bag].gt = 0
					q.bagdat[bag].fallh = 0
					q.bagdat[bag].dir = -1
					q.bagdat[bag].wobbling = false
					q.bagdat[bag].wt = 15
					q.bagdat[bag].unfallen = true
					q.bagdat[bag].x = x*20 + 12
					q.bagdat[bag].y = y*18 + 18
					q.bagdat[bag].h = x
					q.bagdat[bag].v = y
					q.bagdat[bag].xr = 0
					q.bagdat[bag].yr = 0
					bag++
				}
			}
		}
	}
	var i int
	if q.dig.Main.getcplayer() == 0 {
		for i = 1; i < 8; i++ {
			q.bagdat1[i] = q.bagdat[i].copyFrom()
		}
	} else {
		for i = 1; i < 8; i++ {
			q.bagdat2[i] = q.bagdat[i].copyFrom()
		}
	}
}

func (rcvr *Bags) pushbag(bag int, dir int) bool {
	var x int
	var y int
	var h int
	var v int
	var ox int
	var oy int
	var clbits int
	push := true
	ox = rcvr.bagdat[bag].x
	x = rcvr.bagdat[bag].x
	oy = rcvr.bagdat[bag].y
	y = rcvr.bagdat[bag].y
	h = rcvr.bagdat[bag].h
	v = rcvr.bagdat[bag].v
	if rcvr.bagdat[bag].gt != 0 {
		rcvr.getgold(bag)
		return true
	}
	if rcvr.bagdat[bag].dir == 6 && (dir == 4 || dir == 0) {
		clbits = rcvr.dig.Drawing.drawgold(bag, 3, x, y)
		rcvr.dig.Main.incpenalty()
		if clbits&1 != 0 && rcvr.dig.diggery >= y {
			rcvr.dig.killdigger(1, bag)
		}
		if clbits&0x3f00 != 0 {
			rcvr.dig.Monster.squashmonsters(bag, clbits)
		}
		return true
	}
	if x == 292 && dir == 0 || x == 12 && dir == 4 || y == 180 && dir == 6 || y == 18 && dir == 2 {
		push = false
	}
	if push {
		switch dir {
		case 0:
			x += 4
		case 4:
			x -= 4
		case 6:
			if rcvr.bagdat[bag].unfallen {
				rcvr.bagdat[bag].unfallen = false
				rcvr.dig.Drawing.drawsquareblob(x, y)
				rcvr.dig.Drawing.drawtopblob(x, y+21)
			} else {
				rcvr.dig.Drawing.drawfurryblob(x, y)
			}
			rcvr.dig.Drawing.eatfield(x, y, dir)
			rcvr.dig.killemerald(h, v)
			y += 6
		}
		switch dir {
		case 6:
			clbits = rcvr.dig.Drawing.drawgold(bag, 3, x, y)
			rcvr.dig.Main.incpenalty()
			if clbits&1 != 0 && rcvr.dig.diggery >= y {
				rcvr.dig.killdigger(1, bag)
			}
			if clbits&0x3f00 != 0 {
				rcvr.dig.Monster.squashmonsters(bag, clbits)
			}
		case 0:
			fallthrough
		case 4:
			rcvr.bagdat[bag].wt = 15
			rcvr.bagdat[bag].wobbling = false
			clbits = rcvr.dig.Drawing.drawgold(bag, 0, x, y)
			rcvr.dig.Main.incpenalty()
			rcvr.pushcount = 1
			if clbits&0xfe != 0 {
				if !rcvr.pushbags(dir, clbits) {
					x = ox
					y = oy
					rcvr.dig.Drawing.drawgold(bag, 0, ox, oy)
					rcvr.dig.Main.incpenalty()
					push = false
				}
			}
			if clbits&1 != 0 || clbits&0x3f00 != 0 {
				x = ox
				y = oy
				rcvr.dig.Drawing.drawgold(bag, 0, ox, oy)
				rcvr.dig.Main.incpenalty()
				push = false
			}
		}
		if push {
			rcvr.bagdat[bag].dir = dir
		} else {
			rcvr.bagdat[bag].dir = rcvr.dig.reversedir(dir)
		}
		rcvr.bagdat[bag].x = x
		rcvr.bagdat[bag].y = y
		rcvr.bagdat[bag].h = (x - 12) / 20
		rcvr.bagdat[bag].v = (y - 18) / 18
		rcvr.bagdat[bag].xr = (x - 12) % 20
		rcvr.bagdat[bag].yr = (y - 18) % 18
	}
	return push
}

func (q *Bags) pushbags(dir int, bits int) bool {
	var bag int
	var bit int
	push := true
	bit = 2
	for bag = 1; bag < 8; bag++ {
		if (bits & bit) != 0 {
			if !q.pushbag(bag, dir) {
				push = false
			}
		}
		bit <<= 1
	}
	return push
}

func (q *Bags) pushudbags(bits int) bool {
	var bag int
	var b int
	push := true
	b = 2
	for bag = 1; bag < 8; bag++ {
		if (bits & b) != 0 {
			if q.bagdat[bag].gt != 0 {
				q.getgold(bag)
			} else {
				push = false
			}
		}
		b <<= 1
	}
	return push
}

func (q *Bags) removebag(bag int) {
	if q.bagdat[bag].exist {
		q.bagdat[bag].exist = false
		q.dig.Sprite.erasespr(bag)
	}
}

func (q *Bags) removebags(bits int) {
	var bag int
	var b int
	b = 2
	for bag = 1; bag < 8; bag++ {
		if (q.bagdat[bag].exist) && ((bits & b) != 0) {
			q.removebag(bag)
		}
		b <<= 1
	}
}

func (rcvr *Bags) updatebag(bag int) {
	var x int
	var h int
	var xr int
	var y int
	var v int
	var yr int
	var wbl int
	x = rcvr.bagdat[bag].x
	h = rcvr.bagdat[bag].h
	xr = rcvr.bagdat[bag].xr
	y = rcvr.bagdat[bag].y
	v = rcvr.bagdat[bag].v
	yr = rcvr.bagdat[bag].yr
	switch rcvr.bagdat[bag].dir {
	case -1:
		if y < 180 && xr == 0 {
			if rcvr.bagdat[bag].wobbling {
				if rcvr.bagdat[bag].wt == 0 {
					rcvr.bagdat[bag].dir = 6
					rcvr.dig.Sound.soundfall()
					rcvr.dig.newSound.fallStart()
					break
				}
				rcvr.dig.newSound.playLooseLevel(rcvr.bagdat[bag].wt)
				rcvr.bagdat[bag].wt--
				wbl = rcvr.bagdat[bag].wt % 8
				if !(wbl&1 != 0) {
					rcvr.dig.Drawing.drawgold(bag, rcvr.wblanim[uint32(wbl)>>1], x, y)
					rcvr.dig.Main.incpenalty()
					rcvr.dig.Sound.soundwobble()
				}
			} else if rcvr.dig.Monster.getfield(h, v+1)&0xfdf != 0xfdf {
				if !rcvr.dig.checkdiggerunderbag(h, v+1) {
					rcvr.bagdat[bag].wobbling = true
				}
			}
		} else {
			rcvr.bagdat[bag].wt = 15
			rcvr.bagdat[bag].wobbling = false
		}
	case 0:
		fallthrough
	case 4:
		if xr == 0 {
			if y < 180 && rcvr.dig.Monster.getfield(h, v+1)&0xfdf != 0xfdf {
				rcvr.bagdat[bag].dir = 6
				rcvr.bagdat[bag].wt = 0
				rcvr.dig.Sound.soundfall()
				rcvr.dig.newSound.fallStart()
			} else {
				rcvr.baghitground(bag)
			}
		}
	case 6:
		if yr == 0 {
			rcvr.bagdat[bag].fallh++
		}
		if y >= 180 {
			rcvr.baghitground(bag)
		} else if rcvr.dig.Monster.getfield(h, v+1)&0xfdf == 0xfdf {
			if yr == 0 {
				rcvr.baghitground(bag)
			}
		}
		rcvr.dig.Monster.checkmonscared(rcvr.bagdat[bag].h)
		break
	}
	if rcvr.bagdat[bag].dir != -1 {
		if rcvr.bagdat[bag].dir != 6 && rcvr.pushcount != 0 {
			rcvr.pushcount--
		} else {
			rcvr.pushbag(bag, rcvr.bagdat[bag].dir)
		}
	}
}
