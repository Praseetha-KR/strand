package strand

import (
	"sort"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

// --- Test Helpers ---

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, c := range s {
		r = append(r, c)
	}
	return r
}

func isAnagram(s1, s2 string) bool {
	var r1 ByRune = stringToRuneSlice(s1)
	var r2 ByRune = stringToRuneSlice(s2)

	sort.Sort(r1)
	sort.Sort(r2)

	return string(r1) == string(r2)
}

// --- /Test Helpers ---

func TestShuffle(t *testing.T) {
	s := "Hello, World!"
	s1 := shuffle(s)

	assert.Assert(t, s != s1)
	assert.Equal(t, len(s), len(s1))
	assert.Assert(t, isAnagram(s, s1))
}

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := randInt(1, 3)
		assert.Assert(t, n == 1 || n == 2 || n == 3)
	}
}

func TestRandSizedString(t *testing.T) {
	s := randSizedString(1, 13, "abc")
	l := len(s)
	assert.Assert(t, l >= 1 && l <= 13)

	// should only contain charset
	s1 := ""
	for _, c := range s {
		if !strings.Contains("abc", string(c)) {
			s1 = s1 + string(c)
		}
		assert.Assert(t, s1 == "")
	}
}
