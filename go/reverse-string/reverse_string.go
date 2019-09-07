package reverse

import "strings"

func Reverse(input string) string {
	s := strings.Split(input, "")
	start := 0
	end := len(s) - 1

	for start < end {
		temp := s[start]
		s[start] = s[end]
		s[end] = temp
		start++
		end--
	}
	return strings.Join(s, "")

}
