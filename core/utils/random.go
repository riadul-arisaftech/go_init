package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numberBytes = "0123456789"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// helper function for creating random email for testing purpose
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

// helper function for generating OTP
func GenerateOTP() string {
	sb := strings.Builder{}
	sb.Grow(4)
	for i := 1; i < 5; i++ {
		idx := rand.Int63() % int64(len(numberBytes))
		sb.WriteByte(numberBytes[idx])
	}
	return sb.String()
}
