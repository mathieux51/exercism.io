package acronym

import (
	"strings"
)

// Abbreviate ...
func Abbreviate(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", "", -1)
	words := strings.Fields(s)
	var a []string
	for i := range words {
		firstLetter := string(words[i][0])
		u := strings.ToUpper(firstLetter)
		a = append(a, u)
	}
	return strings.Join(a, "")
}
