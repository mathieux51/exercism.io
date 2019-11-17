package robotname

import (
	"math/rand"
	"strconv"
	"strings"
)

type Robot struct {
	name string
}

var names = make(map[string]struct{})

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomLetter() string {
	s := 'a' + randInt(0, 26)
	return strings.ToUpper(string(s))
}

func generateName() string {
	name := make([]string, 5)
	name = append(name, randomLetter())
	name = append(name, randomLetter())
	name = append(name, strconv.Itoa(randInt(0, 9)))
	name = append(name, strconv.Itoa(randInt(0, 9)))
	name = append(name, strconv.Itoa(randInt(0, 9)))

	var n string = strings.Join(name, "")
	_, notUnique := names[n]

	if notUnique {
		return generateName()
	}

	names[n] = struct{}{}

	return n
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.name = generateName()

	return r.name, nil
}

func (r *Robot) Reset() {
	delete(names, r.name)
	r.name = ""
}
