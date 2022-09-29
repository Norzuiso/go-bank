package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func GenerateRandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func GenerateRandomOwner() string {
	return GenerateRandomString(6)
}

func GenerateRandomAmountMoney() int64 {
	return GenerateRandomInt(0, 1000)
}

func GenerateRandomCurrency() string {
	currencies := []string{"EUR", "USD", "MX"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
