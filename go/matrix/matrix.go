package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	slice [][]int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	m := &Matrix{}
	for i, row := range rows {
		fields := strings.Fields(row)

		if len(fields) == 0 {
			return nil, errors.New("empty row")
		}

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
	m.slice[i][j] = v
	return true
}
