package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences must have equal length")
	}

	var distance int
	as := strings.Split(a, "")
	bs := strings.Split(b, "")

	for i := range as {
		if as[i] != bs[i] {
			distance += 1
		}
	}

	return distance, nil
}
