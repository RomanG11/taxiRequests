package rand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

//
// String returns a random string based on charset
//
func String(length int) string {
	return stringWithCharset(length, charset)
}

func stringWithCharset(length int, charset string) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	seededRand = rand.New(rand.NewSource(seededRand.Int63() - 25))
	return string(b)
}
