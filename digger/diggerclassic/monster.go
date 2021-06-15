package diggerclassic

type monsterdata struct {
	x     int
	y     int
	h     int
	v     int
	xr    int
	yr    int
	dir   int
	hdir  int
	t     int
	hnt   int
	death int
	bag   int
	dtime int
	stime int
	flag  bool
	nob   bool
	alive bool
}

func NewMonsterData() monsterdata {
	d := monsterdata{}
	return d
}
