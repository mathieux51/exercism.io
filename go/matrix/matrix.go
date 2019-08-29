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
	m := Matrix{}
	var length int
	for i, row := range rows {

		if len(row) == 0 {
			return nil, errors.New("empty row")
		}

		fields := strings.Fields(row)

		if i == 0 {
			length = len(fields)
		}

		if i > 0 && length != len(fields) {
			return nil, errors.New("funky matrice")
		}

		r := []int{}
		for _, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			r = append(r, v)
		}
		m.slice = append(m.slice, r)
	}
	return &m, nil
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
