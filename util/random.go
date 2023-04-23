package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())

}

// generate a random interger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var stringBuilder strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)
	}
	return stringBuilder.String()

}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 { // in cents
	return RandomInt(20, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}