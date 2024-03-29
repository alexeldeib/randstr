/*
Simple random strings.

Don't use for anything where security is important, as it relies on math/rand, not cryptographicaly secure generators.

 Credit to https://stackoverflow.com/a/31832326 for the default implementation.
*/
package randstr

import (
	mrand "math/rand"
	"strings"
	"time"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	LowerBytes    = "abcdefghijklmnopqrstuvwxyz0123456789"
	SafeBytes     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// NewBoring generates a string of n bytes from the known set of random safe characters.
func NewBoring(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = SafeBytes[mrand.Intn(len(SafeBytes))]
	}
	return string(b)
}

// New generates a string of n characters from the known set of random characters.
func New(n int) string {
	return NewFromAlphabet(n, SafeBytes)
}

// NewFromAlphabet generates a string of n characters from the known set of random characters.
func NewFromAlphabet(n int, alphabet string) string {
	src := mrand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(alphabet) {
			sb.WriteByte(alphabet[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

// NewLowerCaseAlphaNumeric generates a random string with lowercase alphanumeric characters.
func NewLowerCaseAlphaNumeric(n int) string {
	return NewFromAlphabet(n, LowerBytes)
}
