package diggerclassic

import (
	"fmt"
	"time"
)

type Scores struct {
	dig         *digger
	scores      []*ScoreTuple
	substr      string
	highbuf     []rune
	scorehigh   []int64
	scoreinit   []string
	scoreinit2  []string
	scoret      int64
	score1      int64
	score2      int64
	nextbs1     int64
	nextbs2     int64
	hsbuf       string
	scorebuf    []rune
	bonusscore  int
	gotinitflag bool
}

func NewScores(d *digger) *Scores {
	rcvr := new(Scores)

	rcvr.scorehigh = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} // [12]

	rcvr.highbuf = make([]rune, 10)
	rcvr.scoreinit = make([]string, 11)
	rcvr.scoret = 0
	rcvr.score1 = 0
	rcvr.score2 = 0
	rcvr.nextbs1 = 0
	rcvr.nextbs2 = 0
	rcvr.scorebuf = make([]rune, 512)
	rcvr.bonusscore = 20000
	rcvr.gotinitflag = false
	rcvr.dig = d
	return rcvr
}

func (rcvr *Scores) _submit(n string, s int) []*ScoreTuple {
	if len(rcvr.dig.subaddr) < 1 {
		now := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
		ms := 16 + (now % (65536 - 16))
		rcvr.substr = fmt.Sprintf("%v%v%v%v%v%v%v", n, '+', s, '+', ms, '+',
			((ms+int64(32768))*int64(s))%65536)
	}
	return rcvr.scores
}

func (q *Scores) _updatescores(o []*ScoreTuple) {
	if o == nil {
		return
	}
	in := make([]string, 10)
	sc := make([]int, 10)
	for i := 0; i < 10; i++ {
		in[i] = o[i].Key
		sc[i] = o[i].Value
	}
	for i := 0; i < 10; i++ {
		q.scoreinit[i+1] = in[i]
		q.scorehigh[i+2] = int64(sc[i])
	}
}

func (rcvr *Scores) addscore(score int) {
	if rcvr.dig.Main.getcplayer() == 0 {
		rcvr.score1 += int64(score)
		if rcvr.score1 > 999999 {
			rcvr.score1 = 0
		}
		rcvr.writenum(rcvr.score1, 0, 0, 6, 1)
		if rcvr.score1 >= rcvr.nextbs1 {
			if rcvr.dig.Main.getlives(1) < 5 {
				rcvr.dig.Main.addlife(1)
				rcvr.dig.Drawing.drawlives()
			}
			rcvr.nextbs1 += int64(rcvr.bonusscore)
		}
	} else {
		rcvr.score2 += int64(score)
		if rcvr.score2 > 999999 {
			rcvr.score2 = 0
		}
		if rcvr.score2 < 100000 {
			rcvr.writenum(rcvr.score2, 236, 0, 6, 1)
		} else {
			rcvr.writenum(rcvr.score2, 248, 0, 6, 1)
		}
		if rcvr.score2 > rcvr.nextbs2 {
			if rcvr.dig.Main.getlives(2) < 5 {
				rcvr.dig.Main.addlife(2)
				rcvr.dig.Drawing.drawlives()
			}
			rcvr.nextbs2 += int64(rcvr.bonusscore)
		}
	}
	rcvr.dig.Main.incpenalty()
	rcvr.dig.Main.incpenalty()
	rcvr.dig.Main.incpenalty()
}

func (rcvr *Scores) drawscores() {
	rcvr.writenum(rcvr.score1, 0, 0, 6, 3)
	if rcvr.dig.Main.nplayers == 2 {
		if rcvr.score2 < 100000 {
			rcvr.writenum(rcvr.score2, 236, 0, 6, 3)
		} else {
			rcvr.writenum(rcvr.score2, 248, 0, 6, 3)
		}
	}
}

func (rcvr *Scores) endofgame() {
	var i int
	var j int
	var z int
	rcvr.addscore(0)
	if rcvr.dig.Main.getcplayer() == 0 {
		rcvr.scoret = rcvr.score1
	} else {
		rcvr.scoret = rcvr.score2
	}
	if rcvr.scoret > rcvr.scorehigh[11] {
		rcvr.dig.Pc.gclear()
		rcvr.drawscores()
		rcvr.dig.Main.pldispbuf = "PLAYER "
		if rcvr.dig.Main.getcplayer() == 0 {
			rcvr.dig.Main.pldispbuf += "1"
		} else {
			rcvr.dig.Main.pldispbuf += "2"
		}
		rcvr.dig.Drawing.outtext(rcvr.dig.Main.pldispbuf, 108, 0, 2, true)
		rcvr.dig.Drawing.outtext(" NEW HIGH SCORE ", 64, 40, 2, true)
		rcvr.getinitials()
		rcvr._updatescores(rcvr._submit(rcvr.scoreinit[0], int(rcvr.scoret)))
		rcvr.shufflehigh()
		WriteToStorage(rcvr)
	} else {
		rcvr.dig.Main.cleartopline()
		rcvr.dig.Drawing.outtext("GAME OVER", 104, 0, 3, true)
		rcvr._updatescores(rcvr._submit("...", int(rcvr.scoret)))
		rcvr.dig.Sound.killsound()
		for j = 0; j < 20; j++ {
			for i = 0; i < 2; i++ {
				rcvr.dig.Sprite.setretr(true)
				//		dig.Pc.ginten(1);
				rcvr.dig.Pc.gpal(1 - (j & 1))
				rcvr.dig.Sprite.setretr(false)
				for z = 0; z < 111; z++ {
					/* A delay loop */
				}
				rcvr.dig.Pc.gpal(0)
				//		dig.Pc.ginten(0);
				rcvr.dig.Pc.ginten(1 - i&1)
				rcvr.dig.newframe()
			}
		}
		rcvr.dig.Sound.setupsound()
		rcvr.dig.Drawing.outtext("         ", 104, 0, 3, true)
		rcvr.dig.Sprite.setretr(true)
	}
}

func (rcvr *Scores) flashywait(n int) {
	time.Sleep(time.Duration(n*2) * time.Millisecond)
}

func (rcvr *Scores) getinitial(x int, y int) int {
	var i int
	var j int
	rcvr.dig.Input.keypressed = 0
	rcvr.dig.Pc.gwrite(x, y, '_', 3, true)
	for j = 0; j < 5; j++ {
		for i = 0; i < 40; i++ {
			if rcvr.dig.Input.keypressed&0x80 == 0 && rcvr.dig.Input.keypressed != 0 {
				return rcvr.dig.Input.keypressed
			}
			rcvr.flashywait(15)
		}
		for i = 0; i < 40; i++ {
			if rcvr.dig.Input.keypressed&0x80 == 0 && rcvr.dig.Input.keypressed != 0 {
				rcvr.dig.Pc.gwrite(x, y, '_', 3, true)
				return rcvr.dig.Input.keypressed
			}
			rcvr.flashywait(15)
		}
	}
	rcvr.gotinitflag = true
	return 0
}

func (rcvr *Scores) getinitials() {
	var k int
	var i int
	rcvr.dig.Drawing.outtext("ENTER YOUR", 100, 70, 3, true)
	rcvr.dig.Drawing.outtext(" INITIALS", 100, 90, 3, true)
	rcvr.dig.Drawing.outtext("_ _ _", 128, 130, 3, true)
	rcvr.scoreinit[0] = "..."
	rcvr.dig.Sound.killsound()
	rcvr.gotinitflag = false
	for i = 0; i < 3; i++ {
		k = 0
		for k == 0 && !rcvr.gotinitflag {
			k = rcvr.getinitial(i*24+128, 130)
			if i != 0 && k == 8 {
				i--
			}
			k = rcvr.dig.Input.getasciikey(k)
		}
		if k != 0 {
			rcvr.dig.Pc.gwrite(i*24+128, 130, k, 3, true)
			chars := []rune(rcvr.scoreinit[0])
			chars[i] = rune(k)
			rcvr.scoreinit[0] = string(chars)
		}
	}
	rcvr.dig.Input.keypressed = 0
	for i = 0; i < 20; i++ {
		rcvr.flashywait(15)
	}
	rcvr.dig.Sound.setupsound()
	rcvr.dig.Pc.gclear()
	rcvr.dig.Pc.gpal(0)
	rcvr.dig.Pc.ginten(0)
	rcvr.dig.newframe()
	rcvr.dig.Sprite.setretr(true)
}

func (rcvr *Scores) Init() {
	if !ReadFromStorage(rcvr) {
		CreateInStorage(rcvr)
	}
}

func (rcvr *Scores) initscores() {
	rcvr.addscore(0)
}

func (q *Scores) loadscores() {
	p := 1
	var i int
	var x int
	for i = 1; i < 11; i++ {
		for x = 0; x < 3; x++ {
			q.scoreinit[i] = "..."
		}
		p += 2
		for x = 0; x < 6; x++ {
			q.highbuf[x] = q.scorebuf[p]
			p++
		}
		q.scorehigh[i+1] = 0
	}
	if q.scorebuf[0] != 's' {
		for i = 0; i < 11; i++ {
			q.scorehigh[i+1] = 0
			q.scoreinit[i] = "..."
		}
	}
}

func (rcvr *Scores) numtostring(n int64) string {
	var x int
	p := ""
	for x = 0; x < 6; x++ {
		p = string(n%10) + p
		n /= 10
		if n == 0 {
			x++
			break
		}
	}
	for ; x < 6; x++ {
		p = " " + p
	}
	return p
}

func (rcvr *Scores) scorebonus() {
	rcvr.addscore(1000)
}

func (rcvr *Scores) scoreeatm() {
	rcvr.addscore(rcvr.dig.eatmsc * 200)
	rcvr.dig.eatmsc <<= 1
}

func (rcvr *Scores) scoreemerald() {
	rcvr.addscore(25)
}

func (rcvr *Scores) scoregold() {
	rcvr.addscore(500)
}

func (rcvr *Scores) scorekill() {
	rcvr.addscore(250)
}

func (rcvr *Scores) scoreoctave() {
	rcvr.addscore(250)
}

func (q *Scores) showtable() {
	var i int
	var col int
	q.dig.Drawing.outtext2("HIGH SCORES", 16, 25, 3)
	col = 2
	for i = 1; i < 11; i++ {
		q.hsbuf = fmt.Sprintf("%v%v", q.scoreinit[i]+"  ", (q.scorehigh[i+1]))
		q.dig.Drawing.outtext2(q.hsbuf, 16, 31+13*i, col)
		col = 1
	}
}

func (q *Scores) shufflehigh() {
	var i int
	var j int
	for j = 10; j > 1; j-- {
		if q.scoret < q.scorehigh[j] {
			break
		}
	}
	for i = 10; i > j; i-- {
		q.scorehigh[i+1] = q.scorehigh[i]
		q.scoreinit[i] = q.scoreinit[i-1]
	}
	q.scorehigh[j+1] = q.scoret
	q.scoreinit[j] = q.scoreinit[0]
}

func (rcvr *Scores) writecurscore(bp6 int) {
	if rcvr.dig.Main.getcplayer() == 0 {
		rcvr.writenum(rcvr.score1, 0, 0, 6, bp6)
	} else if rcvr.score2 < 100000 {
		rcvr.writenum(rcvr.score2, 236, 0, 6, bp6)
	} else {
		rcvr.writenum(rcvr.score2, 248, 0, 6, bp6)
	}
}

func (rcvr *Scores) writenum(n int64, x int, y int, w int, c int) {
	var d int
	xp := (w-1)*12 + x
	for w > 0 {
		d = int(n % 10)
		if w > 1 || d > 0 {
			rcvr.dig.Pc.gwrite(xp, y, d+'0', c, false)
		}
		n /= 10
		w--
		xp -= 12
	}
}

func (q *Scores) zeroscores() {
	q.score2 = 0
	q.score1 = 0
	q.scoret = 0
	q.nextbs1 = int64(q.bonusscore)
	q.nextbs2 = int64(q.bonusscore)
}
