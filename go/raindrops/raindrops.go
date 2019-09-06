package raindrops

import "strconv"

func Convert(i int) string {
	var s string
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if s != "" {
		return s
	}
	return strconv.Itoa(i)
}
