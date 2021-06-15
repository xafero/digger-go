package diggerclassic

type runnable func()

type Thread struct {
	cb runnable
}

func NewThread(f runnable) *Thread {
	d := new(Thread)
	d.cb = f
	return d
}

func (q Thread) Stop() {
}

func (q Thread) Start() {
	go q.cb()
}
