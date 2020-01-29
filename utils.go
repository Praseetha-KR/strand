package strand

import "math/rand"

func shuffle(s string) string {
	b := []byte(s)
	rand.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})
	return string(b)
}

func randInt(min, max int) int {
	return rseed.Intn(max-min) + min
}

func randSizedString(min int, max int, charset string) string {
	r := randInt(min, max)
	return random(r, charset)
}
