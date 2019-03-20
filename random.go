package random

import (
	"math/rand"
	"time"
)

var letterRunes = []rune{}

// Default Runes
func init() {
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
}

// Set New Runes
func SetNewRunes(newRunes []rune) {
	letterRunes = newRunes
}

// Random Runes to String
func RandStringRunes(n int) string {
	privateRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[privateRand.Intn(len(letterRunes))]
	}
	return string(b)
}
