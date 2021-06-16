package diggerclassic

import "math"

type Monster struct {
	dig           *digger
	mondat        []*monsterdata
	nextmonster   int
	totalmonsters int
	maxmononscr   int
	nextmontime   int
	mongaptime    int
	unbonusflag   bool
	mongotgold    bool
}

func NewMonster(d *digger) *Monster {
	rcvr := new(Monster)

	rcvr.mondat = []*monsterdata{NewMonsterData(), NewMonsterData(), NewMonsterData(),
		NewMonsterData(), NewMonsterData(), NewMonsterData(),
	}

	rcvr.nextmonster = 0
	rcvr.totalmonsters = 0
	rcvr.maxmononscr = 0
	rcvr.nextmontime = 0
	rcvr.mongaptime = 0
	rcvr.unbonusflag = false
	rcvr.mongotgold = false
	rcvr.dig = d
	return rcvr
}

func (q *Monster) checkcoincide(mon int, bits int) {
	var m int
	var b int
	b = 256
	for m = 0; m < 6; m++ {
		if ((bits & b) != 0) && (q.mondat[mon].dir == q.mondat[m].dir) && (q.mondat[m].stime == 0) && (q.mondat[mon].stime == 0) {
			q.mondat[m].dir = q.dig.reversedir(q.mondat[m].dir)
		}
		b <<= 1
	}
}

func (q *Monster) checkmonscared(h int) {
	var m int
	for m = 0; m < 6; m++ {
		if (h == q.mondat[m].h) && (q.mondat[m].dir == 2) {
			q.mondat[m].dir = 6
		}
	}
}

func (q *Monster) createmonster() {
	for i := 0; i < 6; i++ {
		if !q.mondat[i].flag {
			q.mondat[i].flag = true
			q.mondat[i].alive = true
			q.mondat[i].t = 0
			q.mondat[i].nob = true
			q.mondat[i].hnt = 0
			q.mondat[i].h = 14
			q.mondat[i].v = 0
			q.mondat[i].x = 292
			q.mondat[i].y = 18
			q.mondat[i].xr = 0
			q.mondat[i].yr = 0
			q.mondat[i].dir = 4
			q.mondat[i].hdir = 4
			q.nextmonster++
			q.nextmontime = q.mongaptime
			q.mondat[i].stime = 5
			q.dig.Sprite.movedrawspr(i+8, q.mondat[i].x, q.mondat[i].y)
			break
		}
	}
}

func (rcvr *Monster) domonsters() {
	var i int
	if rcvr.nextmontime > 0 {
		rcvr.nextmontime--
	} else {
		if rcvr.nextmonster < rcvr.totalmonsters && rcvr.nmononscr() < rcvr.maxmononscr && rcvr.dig.digonscr && !rcvr.dig.bonusmode {
			rcvr.createmonster()
		}
		if rcvr.unbonusflag && rcvr.nextmonster == rcvr.totalmonsters && rcvr.nextmontime == 0 {
			if rcvr.dig.digonscr {
				rcvr.unbonusflag = false
				rcvr.dig.createbonus()
			}
		}
	}
	for i = 0; i < 6; i++ {
		if rcvr.mondat[i].flag {
			if rcvr.mondat[i].hnt > 10-rcvr.dig.Main.levof10() {
				if rcvr.mondat[i].nob {
					rcvr.mondat[i].nob = false
					rcvr.mondat[i].hnt = 0
				}
			}
			if rcvr.mondat[i].alive {
				if rcvr.mondat[i].t == 0 {
					rcvr.monai(i)
					if rcvr.dig.Main.randno(15-rcvr.dig.Main.levof10()) == 0 && rcvr.mondat[i].nob {
						rcvr.monai(i)
					}
				} else {
					rcvr.mondat[i].t--
				}
			} else {
				rcvr.mondie(i)
			}
		}
	}
}

func (q *Monster) erasemonsters() {
	var i int
	for i = 0; i < 6; i++ {
		if q.mondat[i].flag {
			q.dig.Sprite.erasespr(i + 8)
		}
	}
}

func (rcvr *Monster) fieldclear(dir int, x int, y int) bool {
	switch dir {
	case 0:
		if x < 14 {
			if rcvr.getfield(x+1, y)&0x2000 == 0 {
				if rcvr.getfield(x+1, y)&1 == 0 || rcvr.getfield(x, y)&0x10 == 0 {
					return true
				}
			}
		}
	case 4:
		if x > 0 {
			if rcvr.getfield(x-1, y)&0x2000 == 0 {
				if rcvr.getfield(x-1, y)&0x10 == 0 || rcvr.getfield(x, y)&1 == 0 {
					return true
				}
			}
		}
	case 2:
		if y > 0 {
			if rcvr.getfield(x, y-1)&0x2000 == 0 {
				if rcvr.getfield(x, y-1)&0x800 == 0 || rcvr.getfield(x, y)&0x40 == 0 {
					return true
				}
			}
		}
	case 6:
		if y < 9 {
			if rcvr.getfield(x, y+1)&0x2000 == 0 {
				if rcvr.getfield(x, y+1)&0x40 == 0 || rcvr.getfield(x, y)&0x800 == 0 {
					return true
				}
			}
		}
	}
	return false
}

func (rcvr *Monster) getfield(x int, y int) int {
	return rcvr.dig.Drawing.field[y*15+x]
}

func (rcvr *Monster) incmont(n int) {
	var m int
	if n > 6 {
		n = 6
	}
	for m = 1; m < n; m++ {
		rcvr.mondat[m].t++
	}
}

func (rcvr *Monster) incpenalties(bits int) {
	var m int
	var b int
	b = 256
	for m = 0; m < 6; m++ {
		if (bits & b) != 0 {
			rcvr.dig.Main.incpenalty()
		}
		b <<= 1
		b <<= 1
	}
}

func (rcvr *Monster) initmonsters() {
	var i int
	for i = 0; i < 6; i++ {
		rcvr.mondat[i].flag = false
	}
	rcvr.nextmonster = 0
	rcvr.mongaptime = 45 - rcvr.dig.Main.levof10()<<1
	rcvr.totalmonsters = rcvr.dig.Main.levof10() + 5
	switch rcvr.dig.Main.levof10() {
	case 1:
		rcvr.maxmononscr = 3
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fallthrough
	case 6:
		fallthrough
	case 7:
		rcvr.maxmononscr = 4
	case 8:
		fallthrough
	case 9:
		fallthrough
	case 10:
		rcvr.maxmononscr = 5
	}
	rcvr.nextmontime = 10
	rcvr.unbonusflag = true
}

func (q *Monster) killmon(mon int) {
	if q.mondat[mon].flag {
		q.mondat[mon].flag = false
		q.mondat[mon].alive = false
		q.dig.Sprite.erasespr(mon + 8)
		if q.dig.bonusmode {
			q.totalmonsters++
		}
	}
}

func (rcvr *Monster) killmonsters(bits int) int {
	var m int
	var b int
	n := 0
	b = 256
	for m = 0; m < 6; m++ {
		if bits&b != 0 {
			rcvr.killmon(m)
			n++
		}
		b <<= 1
	}
	return n
}

func (q *Monster) monai(mon int) {
	var clbits int
	var monox int
	var monoy int
	var dir int
	var mdirp1 int
	var mdirp2 int
	var mdirp3 int
	var mdirp4 int
	var t int
	var push bool

	monox = q.mondat[mon].x
	monoy = q.mondat[mon].y

	if q.mondat[mon].xr == 0 && q.mondat[mon].yr == 0 {

		/* If we are here the monster needs to know which way to turn next. */

		/* Turn hobbin back into nobbin if it's had its time */

		if q.mondat[mon].hnt > 30+(q.dig.Main.levof10()<<1) {
			if !q.mondat[mon].nob {
				q.mondat[mon].hnt = 0
				q.mondat[mon].nob = true
			}
		}

		/* Set up monster direction properties to chase dig */

		if math.Abs(float64(q.dig.diggery-q.mondat[mon].y)) > math.Abs(float64(q.dig.diggerx-q.mondat[mon].x)) {
			if q.dig.diggery < q.mondat[mon].y {
				mdirp1 = 2
				mdirp4 = 6
			} else {
				mdirp1 = 6
				mdirp4 = 2
			}
			if q.dig.diggerx < q.mondat[mon].x {
				mdirp2 = 4
				mdirp3 = 0
			} else {
				mdirp2 = 0
				mdirp3 = 4
			}
		} else {
			if q.dig.diggerx < q.mondat[mon].x {
				mdirp1 = 4
				mdirp4 = 0
			} else {
				mdirp1 = 0
				mdirp4 = 4
			}
			if q.dig.diggery < q.mondat[mon].y {
				mdirp2 = 2
				mdirp3 = 6
			} else {
				mdirp2 = 6
				mdirp3 = 2
			}
		}

		/* In bonus mode, run away from digger */

		if q.dig.bonusmode {
			t = mdirp1
			mdirp1 = mdirp4
			mdirp4 = t
			t = mdirp2
			mdirp2 = mdirp3
			mdirp3 = t
		}

		/* Adjust priorities so that monsters don't reverse direction unless they
		   really have to */

		dir = q.dig.reversedir(q.mondat[mon].dir)
		if dir == mdirp1 {
			mdirp1 = mdirp2
			mdirp2 = mdirp3
			mdirp3 = mdirp4
			mdirp4 = dir
		}
		if dir == mdirp2 {
			mdirp2 = mdirp3
			mdirp3 = mdirp4
			mdirp4 = dir
		}
		if dir == mdirp3 {
			mdirp3 = mdirp4
			mdirp4 = dir
		}

		/* Introduce a randno element on levels <6 : occasionally swap p1 and p3 */

		if q.dig.Main.randno(q.dig.Main.levof10()+5) == 1 && q.dig.Main.levof10() < 6 {
			t = mdirp1
			mdirp1 = mdirp3
			mdirp3 = t
		}

		/* Check field and find direction */

		if q.fieldclear(mdirp1, q.mondat[mon].h, q.mondat[mon].v) {
			dir = mdirp1
		} else {
			if q.fieldclear(mdirp2, q.mondat[mon].h, q.mondat[mon].v) {
				dir = mdirp2
			} else {
				if q.fieldclear(mdirp3, q.mondat[mon].h, q.mondat[mon].v) {
					dir = mdirp3
				} else {
					if q.fieldclear(mdirp4, q.mondat[mon].h, q.mondat[mon].v) {
						dir = mdirp4
					}
				}
			}
		}

		/* Hobbins don't care about the field: they go where they want. */

		if !q.mondat[mon].nob {
			dir = mdirp1
		}

		/* Monsters take a time penalty for changing direction */

		if q.mondat[mon].dir != dir {
			q.mondat[mon].t++
		}

		/* Save the new direction */

		q.mondat[mon].dir = dir
	}

	/* If monster is about to go off edge of screen, stop it. */

	if (q.mondat[mon].x == 292 && q.mondat[mon].dir == 0) ||
		(q.mondat[mon].x == 12 && q.mondat[mon].dir == 4) ||
		(q.mondat[mon].y == 180 && q.mondat[mon].dir == 6) ||
		(q.mondat[mon].y == 18 && q.mondat[mon].dir == 2) {
		q.mondat[mon].dir = -1
	}

	/* Change hdir for hobbin */

	if q.mondat[mon].dir == 4 || q.mondat[mon].dir == 0 {
		q.mondat[mon].hdir = q.mondat[mon].dir
	}

	/* Hobbins digger */

	if !q.mondat[mon].nob {
		q.dig.Drawing.eatfield(q.mondat[mon].x, q.mondat[mon].y, q.mondat[mon].dir)
	}

	/* (Draw new tunnels) and move monster */

	switch q.mondat[mon].dir {
	case 0:
		if !q.mondat[mon].nob {
			q.dig.Drawing.drawrightblob(q.mondat[mon].x, q.mondat[mon].y)
		}
		q.mondat[mon].x += 4
		break
	case 4:
		if !q.mondat[mon].nob {
			q.dig.Drawing.drawleftblob(q.mondat[mon].x, q.mondat[mon].y)
		}
		q.mondat[mon].x -= 4
		break
	case 2:
		if !q.mondat[mon].nob {
			q.dig.Drawing.drawtopblob(q.mondat[mon].x, q.mondat[mon].y)
		}
		q.mondat[mon].y -= 3
		break
	case 6:
		if !q.mondat[mon].nob {
			q.dig.Drawing.drawbottomblob(q.mondat[mon].x, q.mondat[mon].y)
		}
		q.mondat[mon].y += 3
		break
	}

	/* Hobbins can eat emeralds */

	if !q.mondat[mon].nob {
		q.dig.hitemerald((q.mondat[mon].x-12)/20, (q.mondat[mon].y-18)/18, (q.mondat[mon].x-12)%20, (q.mondat[mon].y-18)%18, q.mondat[mon].dir)
	}

	/* If digger's gone, don't bother */

	if !q.dig.digonscr {
		q.mondat[mon].x = monox
		q.mondat[mon].y = monoy
	}

	/* If monster's just started, don't move yet */

	if q.mondat[mon].stime != 0 {
		q.mondat[mon].stime--
		q.mondat[mon].x = monox
		q.mondat[mon].y = monoy
	}

	/* Increase time counter for hobbin */

	if !q.mondat[mon].nob && q.mondat[mon].hnt < 100 {
		q.mondat[mon].hnt++
	}

	/* Draw monster */

	push = true
	clbits = q.dig.Drawing.drawmon(mon, q.mondat[mon].nob, q.mondat[mon].hdir, q.mondat[mon].x, q.mondat[mon].y)
	q.dig.Main.incpenalty()

	/* Collision with another monster */

	if (clbits & 0x3f00) != 0 {
		q.mondat[mon].t++            /* Time penalty */
		q.checkcoincide(mon, clbits) /* Ensure both aren't moving in the same dir. */
		q.incpenalties(clbits)
	}

	/* Check for collision with bag */

	if (clbits & q.dig.Bags.bagbits()) != 0 {
		q.mondat[mon].t++ /* Time penalty */
		q.mongotgold = false
		if q.mondat[mon].dir == 4 || q.mondat[mon].dir == 0 { /* Horizontal push */
			push = q.dig.Bags.pushbags(q.mondat[mon].dir, clbits)
			q.mondat[mon].t++ /* Time penalty */
		} else {
			if !q.dig.Bags.pushudbags(clbits) { /* Vertical push */
				push = false
			}
		}
		if q.mongotgold { /* No time penalty if monster eats gold */
			q.mondat[mon].t = 0
		}
		if !q.mondat[mon].nob && q.mondat[mon].hnt > 1 {
			q.dig.Bags.removebags(clbits) /* Hobbins eat bags */
		}
	}

	/* Increase hobbin cross counter */

	if q.mondat[mon].nob && ((clbits & 0x3f00) != 0) && q.dig.digonscr {
		q.mondat[mon].hnt++
	}

	/* See if bags push monster back */

	if !push {
		q.mondat[mon].x = monox
		q.mondat[mon].y = monoy
		q.dig.Drawing.drawmon(mon, q.mondat[mon].nob, q.mondat[mon].hdir, q.mondat[mon].x, q.mondat[mon].y)
		q.dig.Main.incpenalty()
		if q.mondat[mon].nob {
			q.mondat[mon].hnt++
		}
		if (q.mondat[mon].dir == 2 || q.mondat[mon].dir == 6) && q.mondat[mon].nob {
			q.mondat[mon].dir = q.dig.reversedir(q.mondat[mon].dir) /* If vertical, give up */
		}
	}

	/* Collision with digger */

	if ((clbits & 1) != 0) && q.dig.digonscr {
		if q.dig.bonusmode {
			q.killmon(mon)
			q.dig.Scores.scoreeatm()
			q.dig.Sound.soundeatm() /* Collision in bonus mode */
		} else {
			q.dig.killdigger(3, 0) /* Kill digger */
		}
	}

	/* Update co-ordinates */

	q.mondat[mon].h = (q.mondat[mon].x - 12) / 20
	q.mondat[mon].v = (q.mondat[mon].y - 18) / 18
	q.mondat[mon].xr = (q.mondat[mon].x - 12) % 20
	q.mondat[mon].yr = (q.mondat[mon].y - 18) % 18
}

func (q *Monster) mondie(mon int) {
	switch q.mondat[mon].death {
	case 1:
		if q.dig.Bags.bagy(q.mondat[mon].bag)+6 > q.mondat[mon].y {
			q.mondat[mon].y = q.dig.Bags.bagy(q.mondat[mon].bag)
		}
		q.dig.Drawing.drawmondie(mon, q.mondat[mon].nob, q.mondat[mon].hdir, q.mondat[mon].x, q.mondat[mon].y)
		q.dig.Main.incpenalty()
		if q.dig.Bags.getbagdir(q.mondat[mon].bag) == -1 {
			q.mondat[mon].dtime = 1
			q.mondat[mon].death = 4
		}
		break
	case 4:
		if q.mondat[mon].dtime != 0 {
			q.mondat[mon].dtime--
		} else {
			q.killmon(mon)
			q.dig.Scores.scorekill()
		}
		break
	}
}

func (rcvr *Monster) mongold() {
	rcvr.mongotgold = true
}

func (rcvr *Monster) monleft() int {
	return rcvr.nmononscr() + rcvr.totalmonsters - rcvr.nextmonster
}

func (q *Monster) nmononscr() int {
	var i int
	n := 0
	for i = 0; i < 6; i++ {
		if q.mondat[i].flag {
			n++
		}
	}
	return n
}

func (q *Monster) squashmonster(mon int, death int, bag int) {
	q.mondat[mon].alive = false
	q.mondat[mon].death = death
	q.mondat[mon].bag = bag
}

func (q *Monster) squashmonsters(bag int, bits int) {
	var m int
	var b int
	b = 256
	for m = 0; m < 6; m++ {
		if bits&b != 0 {
			if q.mondat[m].y >= q.dig.Bags.bagy(bag) {
				q.squashmonster(m, 1, bag)
			}
		}
		b <<= 1
	}
}
