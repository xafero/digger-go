package diggerclassic

type game struct {
	lives   int
	level   int
	dead    bool
	levdone bool
}

func NewGame() game {
	d := game{}
	return d
}
