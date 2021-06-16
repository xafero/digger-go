package diggerclassic

type bag struct {
	x        int
	y        int
	h        int
	v        int
	xr       int
	yr       int
	dir      int
	wt       int
	gt       int
	fallh    int
	wobbling bool
	unfallen bool
	exist    bool
}

func NewBag() *bag {
	d := new(bag)
	return d
}

func (t *bag) copyFrom() *bag {
	d := new(bag)
	d.x = t.x
	d.y = t.y
	d.h = t.h
	d.v = t.v
	d.xr = t.xr
	d.yr = t.yr
	d.dir = t.dir
	d.wt = t.wt
	d.gt = t.gt
	d.fallh = t.fallh
	d.wobbling = t.wobbling
	d.unfallen = t.unfallen
	d.exist = t.exist
	return d
}
