package matrix

type M struct{}

func New(s string) (*M, error) {
	return nil, nil
}

func (m *M) Rows() [][]int {
	return nil
}

func (m *M) Cols() [][]int {
	return nil
}

func (m *M) Set(i, j, v int) bool {
	return true
}
