package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	numbers  = "1234567890"
)

func init() {
	rand.Seed(int64(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomPhoneNumber() string {
	ans := strings.Builder{}

	ans.WriteByte(8)

	k := len(numbers)

	for i := 0; i < 7; i++ {
		c := numbers[rand.Intn(k)]
		ans.WriteByte(c)
	}

	return ans.String()
}

func RandomString(n int) string {
	ans := strings.Builder{}

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		ans.WriteByte(c)
	}

	return ans.String()
}

func RandomUsername() string {
	return RandomString(8)
}
