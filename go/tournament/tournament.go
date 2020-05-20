package tournament

import (
	"fmt"
	"io"
)

type Teams struct {
	TeamName struct {
		MP int `json:"MP"`
		W  int `json:"W"`
		D  int `json:"D"`
		L  int `json:"L"`
		P  int `json:"P"`
	} `json:"team name"`
}

func Tally(r io.Reader, w io.Writer) error {
	var inMemoryTally map[string]Teams
	fmt.Println(inMemoryTally)
	return nil
}
