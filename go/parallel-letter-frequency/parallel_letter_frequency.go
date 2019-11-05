package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(sList []string) FreqMap {
	m := make(chan FreqMap)

	for _, s := range sList {
		go func(s string) {
			m <- OriginalFrequency(s)
		}(s)
	}

	f := <-m
	// How to print channels in go?
	for i := 1; i < len(sList); i++ {
		for k, v := range <-m {
			f[k] += v
		}

	}

	return f

}
