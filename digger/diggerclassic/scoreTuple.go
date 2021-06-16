package diggerclassic

type ScoreTuple struct {
	Key   string
	Value int
}

func NewScoreTuple(key string, value int) *ScoreTuple {
	d := new(ScoreTuple)
	d.Key = key
	d.Value = value
	return d
}
