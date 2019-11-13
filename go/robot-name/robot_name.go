package robotname

import (
	"math/rand"
	"strings"
	"time"
)

type Robot struct {
	name string
}

// random int between 1 and 24
// map int to letter
func randomLetter() string {
	rand.Seed(time.Now().UnixNano())
	s := 'a' + rand.Intn(26)
	return strings.ToUpper(string(s))
}

func generateName() string {
	name := make([]string, 5)
	name = append(name, randomLetter())
	return strings.Join(name, "")
}

func (r Robot) Name() (string, error) {

	name := generateName()
	return name, nil
}
func (r Robot) Reset() {}
