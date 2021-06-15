package diggerclassic

type ColorModel struct {
	bits int
	size int
	r    []byte
	g    []byte
	b    []byte
}

func NewColorModel(bits int, size int, r []byte, g []byte, b []byte) ColorModel {
	q := ColorModel{}
	q.bits = bits
	q.size = size
	q.r = r
	q.g = g
	q.b = b
	return q
}

func (q ColorModel) GetColor(index int) []float64 {
	r := float64(q.r[index])
	g := float64(q.g[index])
	b := float64(q.b[index])
	return []float64{r, g, b}
}
