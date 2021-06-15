package diggerclassic

type Sprite struct {
	dig         *digger
	retrflag    bool
	sprrdrwf    []bool
	sprrecf     []bool
	sprenf      []bool
	sprch       []int
	sprmov      [][]int16
	sprx        []int
	spry        []int
	sprwid      []int
	sprhei      []int
	sprbwid     []int
	sprbhei     []int
	sprnch      []int
	sprnwid     []int
	sprnhei     []int
	sprnbwid    []int
	sprnbhei    []int
	defsprorder []int
	sprorder    []int
}

func NewSprite(d *digger) *Sprite {
	rcvr := new(Sprite)

	rcvr.defsprorder = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // [16]
	rcvr.sprx = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [17]
	rcvr.spry = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [17]
	rcvr.sprwid = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}         // [17]
	rcvr.sprhei = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}         // [17]
	rcvr.sprbwid = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [16]
	rcvr.sprbhei = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [16]
	rcvr.sprnch = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}            // [16]
	rcvr.sprnwid = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [16]
	rcvr.sprnhei = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}           // [16]
	rcvr.sprnbwid = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}          // [16]
	rcvr.sprnbhei = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rcvr.sprch = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	rcvr.sprmov = [][]int16{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}                                       // []
	rcvr.sprrdrwf = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false} // [17]
	rcvr.sprrecf = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}  // [17]
	rcvr.sprenf = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}          // [16]

	rcvr.retrflag = true
	rcvr.sprorder = rcvr.defsprorder
	rcvr.dig = d
	return rcvr
}

func (q *Sprite) bcollide(bx int, si int) bool {
	if q.sprx[bx] >= q.sprx[si] {
		if q.sprx[bx]+q.sprbwid[bx] > q.sprwid[si]*4+q.sprx[si]-q.sprbwid[si]-1 {
			return false
		}
	} else if q.sprx[si]+q.sprbwid[si] > q.sprwid[bx]*4+q.sprx[bx]-q.sprbwid[bx]-1 {
		return false
	}
	if q.spry[bx] >= q.spry[si] {
		if q.spry[bx]+q.sprbhei[bx] <= q.sprhei[si]+q.spry[si]-q.sprbhei[si]-1 {
			return true
		}
		return false
	}
	if q.spry[si]+q.sprbhei[si] <= q.sprhei[bx]+q.spry[bx]-q.sprbhei[bx]-1 {
		return true
	}
	return false
}

func (q *Sprite) bcollides(bx int) int {
	si := bx
	ax := 0
	dx := 0
	bx = 0
	for {
		if q.sprenf[bx] && bx != si {
			if q.bcollide(bx, si) {
				ax |= 1 << dx
			}
			q.sprx[bx] += 320
			q.spry[bx] -= 2
			if q.bcollide(bx, si) {
				ax |= 1 << dx
			}
			q.sprx[bx] -= 640
			q.spry[bx] += 4
			if q.bcollide(bx, si) {
				ax |= 1 << dx
			}
			q.sprx[bx] += 320
			q.spry[bx] -= 2
		}
		bx++
		dx++
		if !(dx != 16) {
			break
		}
	}
	return ax
}

func (q *Sprite) clearrdrwf() {
	var i int
	q.clearrecf()
	for i = 0; i < 17; i++ {
		q.sprrdrwf[i] = false
	}
}

func (q *Sprite) clearrecf() {
	var i int
	for i = 0; i < 17; i++ {
		q.sprrecf[i] = false
	}
}

func (q *Sprite) collide(bx int, si int) bool {
	if q.sprx[bx] >= q.sprx[si] {
		if q.sprx[bx] > q.sprwid[si]*4+q.sprx[si]-1 {
			return false
		}
	} else if q.sprx[si] > q.sprwid[bx]*4+q.sprx[bx]-1 {
		return false
	}
	if q.spry[bx] >= q.spry[si] {
		if q.spry[bx] <= q.sprhei[si]+q.spry[si]-1 {
			return true
		}
		return false
	}
	if q.spry[si] <= q.sprhei[bx]+q.spry[bx]-1 {
		return true
	}
	return false
}

func (q *Sprite) createspr(n int, ch int, mov []int16, wid int, hei int, bwid int, bhei int) {
	q.sprnch[n&15] = ch
	q.sprch[n&15] = ch
	q.sprmov[n&15] = mov
	q.sprnwid[n&15] = wid
	q.sprwid[n&15] = wid
	q.sprnhei[n&15] = hei
	q.sprhei[n&15] = hei
	q.sprnbwid[n&15] = bwid
	q.sprbwid[n&15] = bwid
	q.sprnbhei[n&15] = bhei
	q.sprbhei[n&15] = bhei
	q.sprenf[n&15] = false
}

func (q *Sprite) drawmiscspr(x int, y int, ch int, wid int, hei int) {
	q.sprx[16] = x & -4
	q.spry[16] = y
	q.sprch[16] = ch
	q.sprwid[16] = wid
	q.sprhei[16] = hei
	q.dig.Pc.gputim(q.sprx[16], q.spry[16], q.sprch[16], q.sprwid[16], q.sprhei[16])
}

func (q *Sprite) drawspr(n int, x int, y int) int {
	var bx int
	var t1 int
	var t2 int
	var t3 int
	var t4 int
	bx = n & 15
	x &= -4
	q.clearrdrwf()
	q.setrdrwflgs(bx)
	t1 = q.sprx[bx]
	t2 = q.spry[bx]
	t3 = q.sprwid[bx]
	t4 = q.sprhei[bx]
	q.sprx[bx] = x
	q.spry[bx] = y
	q.sprwid[bx] = q.sprnwid[bx]
	q.sprhei[bx] = q.sprnhei[bx]
	q.clearrecf()
	q.setrdrwflgs(bx)
	q.sprhei[bx] = t4
	q.sprwid[bx] = t3
	q.spry[bx] = t2
	q.sprx[bx] = t1
	q.sprrdrwf[bx] = true
	q.putis()
	q.sprx[bx] = x
	q.spry[bx] = y
	q.sprch[bx] = q.sprnch[bx]
	q.sprwid[bx] = q.sprnwid[bx]
	q.sprhei[bx] = q.sprnhei[bx]
	q.sprbwid[bx] = q.sprnbwid[bx]
	q.sprbhei[bx] = q.sprnbhei[bx]
	q.dig.Pc.ggeti(q.sprx[bx], q.spry[bx], q.sprmov[bx], q.sprwid[bx], q.sprhei[bx])
	q.putims()
	return q.bcollides(bx)
}

func (q *Sprite) erasespr(n int) {
	bx := n & 15
	q.dig.Pc.gputi(q.sprx[bx], q.spry[bx], q.sprmov[bx], q.sprwid[bx], q.sprhei[bx], true)
	q.sprenf[bx] = false
	q.clearrdrwf()
	q.setrdrwflgs(bx)
	q.putims()
}

func (q *Sprite) getis() {
	var i int
	for i = 0; i < 16; i++ {
		if q.sprrdrwf[i] {
			q.dig.Pc.ggeti(q.sprx[i], q.spry[i], q.sprmov[i], q.sprwid[i], q.sprhei[i])
		}
	}
	q.putims()
}

func (q *Sprite) initmiscspr(x int, y int, wid int, hei int) {
	q.sprx[16] = x
	q.spry[16] = y
	q.sprwid[16] = wid
	q.sprhei[16] = hei
	q.clearrdrwf()
	q.setrdrwflgs(16)
	q.putis()
}

func (q *Sprite) initspr(n int, ch int, wid int, hei int, bwid int, bhei int) {
	q.sprnch[n&15] = ch
	q.sprnwid[n&15] = wid
	q.sprnhei[n&15] = hei
	q.sprnbwid[n&15] = bwid
	q.sprnbhei[n&15] = bhei
}

func (q *Sprite) movedrawspr(n int, x int, y int) int {
	bx := n & 15
	q.sprx[bx] = x & -4
	q.spry[bx] = y
	q.sprch[bx] = q.sprnch[bx]
	q.sprwid[bx] = q.sprnwid[bx]
	q.sprhei[bx] = q.sprnhei[bx]
	q.sprbwid[bx] = q.sprnbwid[bx]
	q.sprbhei[bx] = q.sprnbhei[bx]
	q.clearrdrwf()
	q.setrdrwflgs(bx)
	q.putis()
	q.dig.Pc.ggeti(q.sprx[bx], q.spry[bx], q.sprmov[bx], q.sprwid[bx], q.sprhei[bx])
	q.sprenf[bx] = true
	q.sprrdrwf[bx] = true
	q.putims()
	return q.bcollides(bx)
}

func (q *Sprite) putims() {
	var i int
	var j int
	for i = 0; i < 16; i++ {
		j = q.sprorder[i]
		if q.sprrdrwf[j] {
			q.dig.Pc.gputim(q.sprx[j], q.spry[j], q.sprch[j], q.sprwid[j], q.sprhei[j])
		}
	}
}

func (q *Sprite) putis() {
	var i int
	for i = 0; i < 16; i++ {
		if q.sprrdrwf[i] {
			q.dig.Pc.gputi2(q.sprx[i], q.spry[i], q.sprmov[i], q.sprwid[i], q.sprhei[i])
		}
	}
}

func (q *Sprite) setrdrwflgs(n int) {
	var i int
	if !q.sprrecf[n] {
		q.sprrecf[n] = true
		for i = 0; i < 16; i++ {
			if q.sprenf[i] && i != n {
				if q.collide(i, n) {
					q.sprrdrwf[i] = true
					q.setrdrwflgs(i)
				}
				q.sprx[i] += 320
				q.spry[i] -= 2
				if q.collide(i, n) {
					q.sprrdrwf[i] = true
					q.setrdrwflgs(i)
				}
				q.sprx[i] -= 640
				q.spry[i] += 4
				if q.collide(i, n) {
					q.sprrdrwf[i] = true
					q.setrdrwflgs(i)
				}
				q.sprx[i] += 320
				q.spry[i] -= 2
			}
		}
	}
}

func (rcvr *Sprite) setretr(f bool) {
	rcvr.retrflag = f
}

func (rcvr *Sprite) setsprorder(newsprorder []int) {
	if newsprorder == nil {
		rcvr.sprorder = rcvr.defsprorder
	} else {
		rcvr.sprorder = newsprorder
	}
}
