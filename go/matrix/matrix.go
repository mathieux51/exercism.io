package matrix

import (
	"strconv"
	"strings"
)

type Matrix struct{}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	// var cols = 0
	m := &Matrix{}
	for i, row := range rows {
		fields := strings.Fields(row)
		for j, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			m.Set(i, j, v)
		}

	}
	return m, nil
}

func (m *Matrix) Rows() [][]int {
	return nil
}

func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Set(i, j, v int) bool {
	return true
}
