package random

import (
	"time"

	"golang.org/x/exp/rand"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	chars := []rune("qwertyuiopasdfghjkklzxcvbnm12345677890")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}
	return string(b)
}
