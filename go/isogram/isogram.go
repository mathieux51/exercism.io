package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {

	unique := make(map[string]struct{})
	noHyphens := strings.Replace(input, "-", " ", -1)
	lower := strings.ToLower(noHyphens)
	letters := strings.Split(lower, "")
	for _, s := range letters {

		// reset map if new word
		if s == " " {
			unique = make(map[string]struct{})
		}

		if _, ok := unique[s]; ok {
			return false
		}
		unique[s] = struct{}{}
	}
	return true
}
