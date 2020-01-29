package strand

import (
	"errors"
	"math/rand"
	"time"
)

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const alpha = alphaUpper + alphaLower
const numeric = "0123456789"
const special = "+-*/#!|@$%^&*_~`,.:;?=(){}[]<>"
const urlsafe = "-_~."

func charMap() map[string]bool {
	m := make(map[string]bool)
	charset := alpha + numeric + special
	for _, c := range charset {
		m[string(c)] = true
	}
	return m
}

var supportedChars = charMap()
var rseed = rand.New(rand.NewSource(time.Now().UnixNano()))

// random string from charset
func random(length int, charset string) string {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = charset[rseed.Intn(len(charset))]
	}
	return string(arr)
}

// String generates random string of all supported characters
func String(length int) string {
	return random(length, alpha+numeric+special)
}

// Alpha generates random string of alphabet
func Alpha(length int) string {
	return random(length, alpha)
}

// AlphaLower generates random string of alphabets lowercase
func AlphaLower(length int) string {
	return random(length, alphaLower)
}

// AlphaUpper generates random string of alphabets uppercase
func AlphaUpper(length int) string {
	return random(length, alphaUpper)
}

// AlphaNumeric generates random string of alphabet & digits
func AlphaNumeric(length int) string {
	return random(length, alpha+numeric)
}

// AlphaLowerNumeric generates random string of lowercase alphabets & digits
func AlphaLowerNumeric(length int) string {
	return random(length, alphaLower+numeric)
}

// AlphaUpperNumeric generates random string of uppercase alphabets & digits
func AlphaUpperNumeric(length int) string {
	return random(length, alphaUpper+numeric)
}

// Numeric generates random string of digits
func Numeric(length int) string {
	return random(length, numeric)
}

// URLSafe generates random string of url-safe characters
func URLSafe(length int) string {
	return random(length, alpha+numeric+urlsafe)
}

// Hex generates random string of Hexadecimal characters
func Hex(length int) string {
	return random(length, alphaLower[0:6]+numeric)
}

// Binary generates random string of binary digits
func Binary(length int) string {
	return random(length, numeric[0:2])
}

// SimplePassword generates password of length 8
func SimplePassword() string {
	return shuffle(random(2, special) + random(2, alphaLower) + random(2, numeric) + random(2, alphaUpper))
}

// URLSafeSimplePassword generates url-safe  password of length 8
func URLSafeSimplePassword() string {
	return shuffle(random(2, urlsafe) + random(2, alphaLower) + random(2, numeric) + random(2, alphaUpper))
}

// Password generates password of length range (12, 40)
func Password() string {
	b1 := randSizedString(3, 6, special)
	b2 := randSizedString(3, 10, numeric)
	b3 := randSizedString(3, 14, alphaLower)
	b4 := randSizedString(3, 10, alphaUpper)
	return shuffle(b1 + b3 + b2 + b4)
}

// URLSafePassword generates url-safe password of length range (12, 40)
func URLSafePassword() string {
	b1 := randSizedString(3, 6, urlsafe)
	b2 := randSizedString(3, 10, numeric)
	b3 := randSizedString(3, 14, alphaLower)
	b4 := randSizedString(3, 10, alphaUpper)
	return shuffle(b1 + b3 + b2 + b4)
}

// From the provided charset generate
func From(characters string, length int) (string, error) {
	if len(characters) == 0 {
		return "", errors.New("strand: empty characters")
	}
	for _, c := range characters {
		if _, ok := supportedChars[string(c)]; !ok {
			return "", errors.New("strand: unsupported character " + string(c))
		}
	}
	return random(length, characters), nil
}
