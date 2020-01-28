package strand

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz"
const numeric = "01234567890"
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

// randomString creates a string with random chars from charset
func randomString(l int, charset string) string {
	arr := make([]byte, l)
	for i := range arr {
		arr[i] = charset[rseed.Intn(len(charset))]
	}
	return string(arr)
}

// String generates random string of all characters
func String(length int) string {
	return randomString(length, alpha+numeric+special)
}

// Alpha generates random string of numbers
func Alpha(length int) string {
	return randomString(length, alpha)
}

// Numeric generates random string of numbers
func Numeric(length int) string {
	return randomString(length, numeric)
}

// AlphaNumeric generates random string of alphabet & numbers
func AlphaNumeric(length int) string {
	return randomString(length, alpha+numeric)
}

// URLSafe generates random string of url-safe characters
func URLSafe(length int) string {
	return randomString(length, alpha+numeric+urlsafe)
}

// RandomFrom generates random string from the provided charset
func RandomFrom(charset string, length int) (string, error) {
	if len(charset) == 0 {
		return "", errors.New("strand: empty charset")
	}
	for _, c := range charset {
		if !strings.Contains(charset, string(c)) {
			return "", errors.New("strand: unsupported character in charset")
		}
	}
	return randomString(length, charset), nil
}
