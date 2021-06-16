package diggerclassic

import "time"

type Pc struct {
	dig           *digger
	source        []*Refresher
	currentSource *Refresher
	width         int
	height        int
	size          int
	pixels        []int
	pal           [][][]byte
}

func NewPc(d *digger) *Pc {
	rcvr := new(Pc)

	palA := [][]byte{
		{0, 0x00, 0xAA, 0xAA},
		{0, 0xAA, 0x00, 0x54},
		{0, 0x00, 0x00, 0x00},
	}
	palB := [][]byte{
		{0, 0x54, 0xFF, 0xFF},
		{0, 0xFF, 0x54, 0xFF},
		{0, 0x54, 0x54, 0x54},
	}
	rcvr.pal = [][][]byte{palA, palB}

	rcvr.source = make([]*Refresher, 2)
	rcvr.width = 320
	rcvr.height = 200
	rcvr.size = rcvr.width * rcvr.height
	rcvr.dig = d
	return rcvr
}

func (rcvr *Pc) gclear() {
	for i := 0; i < rcvr.size; i++ {
		rcvr.pixels[i] = 0
	}
	rcvr.currentSource.NewPixelsAll()
}

func (rcvr *Pc) gethrt() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func (rcvr *Pc) getkips() int {
	return 0
}

func (q *Pc) ggeti(x int, y int, p []int16, w int, h int) {
	src := 0
	dest := y*q.width + x&0xfffc
	for i := 0; i < h; i++ {
		d := dest
		for j := 0; j < w; j++ {
			p[src] = int16(((q.pixels[d]<<2|q.pixels[d+1])<<2|q.pixels[d+2])<<2 | q.pixels[d+3])
			src++
			d += 4
			if src == len(p) {
				return
			}
		}
		dest += q.width
	}
}

func (q *Pc) ggetpix(x int, y int) int {
	ofs := (q.width*y + x) & 0xfffc
	return ((q.pixels[ofs]<<2|q.pixels[ofs+1])<<2|q.pixels[ofs+2])<<2 | q.pixels[ofs+3]
}

func (rcvr *Pc) ginit() {
}

func (q *Pc) ginten(inten int) {
	q.currentSource = q.source[inten&1]
	q.currentSource.NewPixelsAll()
}

func (rcvr *Pc) gpal(pal int) {
}

func (q *Pc) gputi(x int, y int, p []int16, w int, h int, b bool) {
	src := 0
	dest := y*q.width + x&0xfffc
	for i := 0; i < h; i++ {
		d := dest
		for j := 0; j < w; j++ {
			px := p[src]
			src++
			q.pixels[d+3] = int(px & 3)
			px >>= 2
			q.pixels[d+2] = int(px & 3)
			px >>= 2
			q.pixels[d+1] = int(px & 3)
			px >>= 2
			q.pixels[d] = int(px & 3)
			d += 4
			if src == len(p) {
				return
			}
		}
		dest += q.width
	}
}

func (rcvr *Pc) gputi2(x int, y int, p []int16, w int, h int) {
	rcvr.gputi(x, y, p, w, h, true)
}

func (rcvr *Pc) gputim(x int, y int, ch int, w int, h int) {
	spr := cgatable[ch*2]
	msk := cgatable[ch*2+1]
	src := 0
	dest := y*rcvr.width + x&0xfffc
	for i := 0; i < h; i++ {
		d := dest
		for j := 0; j < w; j++ {
			px := spr[src]
			mx := msk[src]
			src++
			if mx&3 == 0 {
				rcvr.pixels[d+3] = int(px) & 3
			}
			px >>= 2
			if mx&(3<<2) == 0 {
				rcvr.pixels[d+2] = int(px) & 3
			}
			px >>= 2
			if mx&(3<<4) == 0 {
				rcvr.pixels[d+1] = int(px) & 3
			}
			px >>= 2
			if mx&(3<<6) == 0 {
				rcvr.pixels[d] = int(px) & 3
			}
			d += 4
			if src == len(spr) || src == len(msk) {
				return
			}
		}
		dest += rcvr.width
	}
}

func (rcvr *Pc) gtitle() {
	src := 0
	dest := 0
	for {
		if src >= len(cgatitledat) {
			break
		}
		b := cgatitledat[src]
		src++
		var l int
		var c int
		if b == 0xfe {
			l = int(cgatitledat[src])
			src++
			if l == 0 {
				l = 256
			}
			c = int(cgatitledat[src])
			src++
		} else {
			l = 1
			c = int(b)
		}
		for i := 0; i < l; i++ {
			px := c
			adst := 0
			if dest < 32768 {
				adst = (dest/320)*640 + dest%320
			} else {
				adst = 320 + (dest-32768)/320*640 + (dest-32768)%320
			}
			rcvr.pixels[adst+3] = px & 3
			px >>= 2
			rcvr.pixels[adst+2] = px & 3
			px >>= 2
			rcvr.pixels[adst+1] = px & 3
			px >>= 2
			rcvr.pixels[adst+0] = px & 3
			dest += 4
			if dest >= 65535 {
				break
			}
		}
		if dest >= 65535 {
			break
		}
	}
}

func (rcvr *Pc) gwrite(x int, y int, ch int, c int, upd bool) {
	dest := x + y*rcvr.width
	ofs := 0
	color := c & 3

	ch -= 32
	if ch < 0 || ch > 0x5f {
		return
	}

	chartab := ascii2cga[ch]
	if chartab == nil {
		return
	}

	for i := 0; i < 12; i++ {
		d := dest
		for j := 0; j < 3; j++ {
			px := chartab[ofs]
			ofs++
			rcvr.pixels[d+3] = int(px) & color
			px >>= 2
			rcvr.pixels[d+2] = int(px) & color
			px >>= 2
			rcvr.pixels[d+1] = int(px) & color
			px >>= 2
			rcvr.pixels[d] = int(px) & color
			d += 4
		}
		dest += rcvr.width
	}
	if upd {
		// Force complete update for high score
		rcvr.currentSource.NewPixelsAll( /* x, y, 12, 12 */ )
	}
}

func (rcvr *Pc) gwrite2(x int, y int, ch int, c int) {
	rcvr.gwrite(x, y, ch, c, false)
}
