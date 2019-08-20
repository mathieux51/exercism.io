package matrix

import (
	"fmt"
	"strings"
)

type Matrix struct{}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	// var cols = 0

	for _, row := range rows {
		fmt.Println(row)
	}

	// fmt.Printf("len(rows): %v\n", len(rows))
	// fmt.Printf("cols: %v\n", cols)
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
