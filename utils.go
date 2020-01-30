package strand

import "math/rand"

// shuffles the provided string
func shuffle(s string) string {
	b := []byte(s)
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	return string(b)
}

// randInt returns random integer from range [min, max]
func randInt(min, max int) int {
	return rseed.Intn((max+1)-min) + min
}

// randSizedString returns string of random characters
// with length range (min, max)
func randSizedString(min int, max int, charset string) string {
	r := randInt(min, max)
	return random(r, charset)
}
