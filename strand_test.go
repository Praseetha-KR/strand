package strand

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"gotest.tools/v3/assert"
)

func TestString(t *testing.T) {
	l := 40
	pat := regexp.MustCompile("^[[:alnum:]+\\-\\*/#!|@$%\\^&_~\\`,\\.:;?=\\(\\){}\\[\\]<>]{" + strconv.Itoa(l) + "}$")

	// length & charset match
	r1 := String(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := String(l)
	assert.Assert(t, pat.MatchString(r2))

	// immediate randoms should not be equal
	assert.Assert(t, r1 != r2)

	// zero or negative length should be handled
	r3 := String(0)
	assert.Equal(t, r3, "")

	r4 := String(-1)
	assert.Equal(t, r4, "")
}

func TestAplha(t *testing.T) {
	l := 10
	pat := regexp.MustCompile(`^[[:alpha:]]{` + strconv.Itoa(l) + `}$`)

	r1 := Alpha(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := Alpha(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestAplhaLower(t *testing.T) {
	l := 10
	pat := regexp.MustCompile(`^[[:lower:]]{` + strconv.Itoa(l) + `}$`)

	r1 := AlphaLower(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := AlphaLower(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestAplhaUpper(t *testing.T) {
	l := 10
	pat := regexp.MustCompile(`^[[:upper:]]{` + strconv.Itoa(l) + `}$`)

	r1 := AlphaUpper(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := AlphaUpper(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestAplhaNumeric(t *testing.T) {
	l := 40
	pat := regexp.MustCompile(`^[[:alnum:]]{` + strconv.Itoa(l) + `}$`)

	r1 := AlphaNumeric(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := AlphaNumeric(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestAplhaLowerNumeric(t *testing.T) {
	l := 40
	pat := regexp.MustCompile(`^[a-z0-9]{` + strconv.Itoa(l) + `}$`)

	r1 := AlphaLowerNumeric(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := AlphaLowerNumeric(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestAplhaUpperNumeric(t *testing.T) {
	l := 40
	pat := regexp.MustCompile(`^[A-Z0-9]{` + strconv.Itoa(l) + `}$`)

	r1 := AlphaUpperNumeric(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := AlphaUpperNumeric(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestNumeric(t *testing.T) {
	l := 18
	pat := regexp.MustCompile(`^[[:digit:]]{` + strconv.Itoa(l) + `}$`)

	r1 := Numeric(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := Numeric(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestURLSafe(t *testing.T) {
	l := 20
	pat := regexp.MustCompile(`^[[:alnum:]-._~]{` + strconv.Itoa(l) + `}$`)

	r1 := URLSafe(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := URLSafe(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestHex(t *testing.T) {
	l := 18
	pat := regexp.MustCompile(`^[[:xdigit:]]{` + strconv.Itoa(l) + `}$`)

	r1 := Hex(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := Hex(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestBinary(t *testing.T) {
	l := 12
	pat := regexp.MustCompile(`^[0-1]{` + strconv.Itoa(l) + `}$`)

	r1 := Binary(l)
	assert.Assert(t, pat.MatchString(r1))

	r2 := Binary(l)
	assert.Assert(t, pat.MatchString(r2))

	assert.Assert(t, r1 != r2)
}

func TestSimplePassword(t *testing.T) {
	pat := regexp.MustCompile("^[[:alnum:]+\\-\\*/#!|@$%\\^&_~\\`,\\.:;?=\\(\\){}\\[\\]<>]{8}$")

	// length & charset match
	r := SimplePassword()
	assert.Assert(t, pat.MatchString(r))

	// charset composition check
	cn, cl, cu, cs := 0, 0, 0, 0
	for _, c := range r {
		if unicode.IsDigit(c) {
			cn++
		}
		if unicode.IsLower(c) {
			cl++
		}
		if unicode.IsUpper(c) {
			cu++
		}
		if strings.Contains(special, string(c)) {
			cs++
		}
	}
	assert.Equal(t, cn, 2)
	assert.Equal(t, cl, 2)
	assert.Equal(t, cu, 2)
	assert.Equal(t, cs, 2)

	// immediate randoms should not be equal
	r1 := SimplePassword()
	assert.Assert(t, r != r1)
}

func TestURLSafeSimplePassword(t *testing.T) {
	pat := regexp.MustCompile("^[[:alnum:]\\-\\._~]{8}$")

	r := URLSafeSimplePassword()
	assert.Assert(t, pat.MatchString(r))

	cn, cl, cu, cs := 0, 0, 0, 0
	for _, c := range r {
		if unicode.IsDigit(c) {
			cn++
		}
		if unicode.IsLower(c) {
			cl++
		}
		if unicode.IsUpper(c) {
			cu++
		}
		if strings.Contains(urlsafe, string(c)) {
			cs++
		}
	}
	assert.Equal(t, cn, 2)
	assert.Equal(t, cl, 2)
	assert.Equal(t, cu, 2)
	assert.Equal(t, cs, 2)

	r1 := URLSafeSimplePassword()
	assert.Assert(t, r != r1)
}

func TestPassword(t *testing.T) {
	pat := regexp.MustCompile("^[[:alnum:]+\\-\\*/#!|@$%\\^&_~\\`,\\.:;?=\\(\\){}\\[\\]<>]{12,40}$")

	r := Password()
	assert.Assert(t, pat.MatchString(r))

	cn, cl, cu, cs := 0, 0, 0, 0
	for _, c := range r {
		if unicode.IsDigit(c) {
			cn++
		}
		if unicode.IsLower(c) {
			cl++
		}
		if unicode.IsUpper(c) {
			cu++
		}
		if strings.Contains(special, string(c)) {
			cs++
		}
	}
	assert.Assert(t, cn >= 3 && cn <= 10)
	assert.Assert(t, cl >= 3 && cl <= 14)
	assert.Assert(t, cu >= 3 && cu <= 10)
	assert.Assert(t, cs >= 3 && cs <= 6)

	r1 := Password()
	assert.Assert(t, r != r1)
}

func TestURLSafePassword(t *testing.T) {
	pat := regexp.MustCompile("^[[:alnum:]\\-\\._~]{12,40}$")

	r := URLSafePassword()
	assert.Assert(t, pat.MatchString(r))

	cn, cl, cu, cs := 0, 0, 0, 0
	for _, c := range r {
		if unicode.IsDigit(c) {
			cn++
		}
		if unicode.IsLower(c) {
			cl++
		}
		if unicode.IsUpper(c) {
			cu++
		}
		if strings.Contains(special, string(c)) {
			cs++
		}
	}
	assert.Assert(t, cn >= 3 && cn <= 10)
	assert.Assert(t, cl >= 3 && cl <= 14)
	assert.Assert(t, cu >= 3 && cu <= 10)
	assert.Assert(t, cs >= 3 && cs <= 6)

	r1 := URLSafePassword()
	assert.Assert(t, r != r1)
}

func TestFrom(t *testing.T) {
	l := 25
	pat := regexp.MustCompile("^[@abc123]{" + strconv.Itoa(l) + "}$")

	// length & charset match
	r1, err := From("abc@123", l)
	assert.Assert(t, pat.MatchString(r1))
	assert.Assert(t, err == nil)

	r2, err := From("abc@123", l)
	assert.Assert(t, pat.MatchString(r1))
	assert.Assert(t, err == nil)

	// immediate randoms should not be equal
	assert.Assert(t, r1 != r2)

	// unsupported characters should error
	r, err := From(" ", 3)
	assert.Error(t, err, "strand: unsupported character  ")
	assert.Equal(t, r, "")

	r, err = From("🤖", 4)
	assert.Error(t, err, "strand: unsupported character 🤖")
	assert.Equal(t, r, "")

	// empty characters should error
	r, err = From("", 3)
	assert.Error(t, err, "strand: empty characters")
	assert.Equal(t, r, "")

	r, err = From("", 0)
	assert.Error(t, err, "strand: empty characters")
	assert.Equal(t, r, "")

	// negative length should error
	r, err = From("abc", -1)
	assert.Error(t, err, "strand: length should be greater than 0")
	assert.Equal(t, r, "")

	r, err = From("abc", 0)
	assert.Assert(t, err == nil)
	assert.Equal(t, r, "")
}
