package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	elements       [][]int
	rowLen, colLen int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	length := make(map[int]struct{})
	var colLen int
	var elements [][]int

	for i, row := range rows {

		if len(row) == 0 {
			return nil, errors.New("empty row")
		}

		fields := strings.Fields(row)

		if i == 0 {
			colLen = len(fields)
			length[colLen] = struct{}{}
		}

		if _, ok := length[len(fields)]; !ok {
			return nil, errors.New("uneven matrice")
		}

		r := []int{}
		for _, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			r = append(r, v)
		}
		elements = append(elements, r)
	}

	m := Matrix{elements: elements, rowLen: len(rows), colLen: colLen}

	return &m, nil
}

func (m *Matrix) Rows() [][]int {
	return nil
}

func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Set(i, j, v int) bool {
	m.elements[i][j] = v
	return true
}
