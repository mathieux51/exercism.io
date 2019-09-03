package pascal

import "math"

func getValue(n, k int) int {
	return int(math.Sqrt(math.Pow(2*float64(n)-2, 2) + math.Pow((2*float64(k)-1), 2)))
}

func Triangle(n int) [][]int {
	if n == 0 {
		return [][]int{{1}}
	}

	k := n + 1
	var c = []int{}
	var t = [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			v := getValue(n, k)
			c = append(c, v)
		}
		t = append(t, c)
	}
	return t
}
