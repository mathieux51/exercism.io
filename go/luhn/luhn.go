package luhn

import (
	"regexp"
)

func Valid(s string) bool {
	r := regexp.MustCompile(`^([0-9]{2,4} ?){1,}$`)
	isValid := r.MatchString(s)
	if !isValid {
		return isValid
	}

	str := strings.Replace(s, " ", "", -1)

	for _, r range str {
		
	}

	return true
}
