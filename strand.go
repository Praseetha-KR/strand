package strand

import (
	"errors"
	"math/rand"
	"time"
)

const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz"
const numeric = "01234567890"
const special = "+-*/#!|@$%^&*_~`,.?=(){}[]"

var rseed = rand.New(rand.NewSource(time.Now().UnixNano()))

// randomString creates a string with random chars from charset
func randomString(l int, charset string) string {
	arr := make([]byte, l)
	for i := range arr {
		arr[i] = charset[rseed.Intn(len(charset))]
	}
	return string(arr)
}

// SubRandom generates random string from provided charset
func SubRandom(length int, charset string) (string, error) {
	if len(charset) == 0 {
		return "", errors.New("strand: empty charset")
	}
	return randomString(length, charset), nil
}

// AlphaNumeric generates random string of alphabet & numbers
func AlphaNumeric(length int) string {
	return randomString(length, alpha+numeric)
}

// Numeric generates random string of numbers
func Numeric(length int) string {
	return randomString(length, alpha)
}

// String generates random string of all characters
func String(length int) string {
	return randomString(length, alpha+numeric+special)
}
