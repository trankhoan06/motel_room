package common

import (
	"math/rand"
	"time"
)

var letter = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

func Randomreques(length int) string {
	b := make([]rune, length)
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	for i := range b {
		b[i] = letter[r.Intn(9999999%len(letter))]
	}
	return string(b)
}
func GetSalt(length int) string {
	if length < 1 {
		length = 50
	}
	return Randomreques(length)
}
