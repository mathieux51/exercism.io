package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct{}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	fmt.Println(rows)
	return nil, nil
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
