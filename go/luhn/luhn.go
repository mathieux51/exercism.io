package luhn

import (
	"strconv"
)

// Valid ...
func Valid(input string) bool {
	var sum int
	var count int

	for n := len(input) - 1; n >= 0; n-- {
		d, err := strconv.Atoi(string(input[n]))

		if err != nil {
			if input[n] == ' ' {
				continue
			}
			return false
		}

		if count%2 != 0 {
			d = d * 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
		count++
	}

	return count > 1 && sum%10 == 0
}
