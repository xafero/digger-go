package diggerclassic

type game struct {
	lives   int
	level   int
	dead    bool
	levdone bool
}

func NewGame() *game {
	d := new(game)
	return d
}
