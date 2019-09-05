// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

func Split(r rune) bool {
	return r == ' ' || r == '-'
}

// Abbreviate should have a comment documenting it.
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
