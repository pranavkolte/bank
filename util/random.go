package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstywxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

}

// generate random number between min max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// generate random string with lenth of n
func randomString(n int) string {
	var stringBuilder strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		stringBuilder.WriteByte(c)

	}

	return stringBuilder.String()

}

// genrate random owener name
func RandomOwner() string {
	return "test-" + randomString(7)
}

// generate random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// genrate Random Currency
func RandomCurrency() string {
	currencies := []string{USD, INR, EUR}
	return currencies[rand.Intn(len(currencies))]
}
