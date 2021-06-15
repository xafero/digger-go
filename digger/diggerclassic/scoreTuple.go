package diggerclassic

type ScoreTuple struct {
	Key   string
	Value int
}

func NewScoreTuple(key string, value int) ScoreTuple {
	d := ScoreTuple{key, value}
	return d
}
