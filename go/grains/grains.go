package grains

import (
	"errors"
)

func Square(input int) (uint64, error) {
	var count uint64

	if input <= 0 || input > 64 {
		return count, errors.New("input should be in between 0 and 64")
	}

	for i := 0; i < input; i++ {

		count = count * 2

		if count == 0 {
			count = 1
		}

	}
	return count, nil
}

func Total() uint64 {
	var count uint64
	var sum uint64
	for i := 0; i < 64; i++ {

		count = count * 2

		if count == 0 {
			count = 1
		}

		sum += count
	}
	return sum
}
