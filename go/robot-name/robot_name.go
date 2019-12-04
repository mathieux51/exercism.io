package robotname

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
)

type Robot struct {
	name string
	m    sync.Mutex
}

var names = make(map[string]struct{})

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomLetter() string {
	s := 'a' + randInt(0, 26)
	return strings.ToUpper(string(s))
}

func generateName(m *sync.Mutex) string {
	name := make([]string, 5)
	name = append(name, randomLetter())
	name = append(name, randomLetter())
	name = append(name, strconv.Itoa(randInt(0, 9)))
	name = append(name, strconv.Itoa(randInt(0, 9)))
	name = append(name, strconv.Itoa(randInt(0, 9)))

	var n string = strings.Join(name, "")
	m.Lock()
	defer m.Unlock()
	_, notUnique := names[n]

	if notUnique {
		return generateName(m)
	}

	names[n] = struct{}{}

	return n
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	r.name = generateName(&r.m)

	return r.name, nil
}

func (r *Robot) Reset() {
	r.m.Lock()
	defer r.m.Unlock()
	delete(names, r.name)
	r.name = ""
}
