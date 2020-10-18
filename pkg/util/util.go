package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt creates random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvyxwz"

// RandomStr creates random string with the length provided
func RandomStr(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < k; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner creates a random string with fixed length of 6
func RandomOwner() string {
	return RandomStr(6)
}

// RandomMoney creates a random number between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
