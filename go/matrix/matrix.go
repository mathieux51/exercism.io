package matrix

import (
	"errors"
	"fmt"
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

		if len(row) == 0 {
			return nil, errors.New("empty row")
		}

		fields := strings.Fields(row)
		fmt.Println(fields)
		for _, f := range fields {
			v, err := strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
			fmt.Println(i, v)
			// m.slice[i] = append(m.slice[i], v)
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
