package util

import (
	"math/rand"
	"strings"
)

func RandomInt(min, max int64) int64 {
	if min >= max {
		panic("invalid range")
	}
	return min + rand.Int63n(max-min+1) // Generates a random integer between min and max (inclusive).
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // Selects a random character from the alphabet.
		sb.WriteByte(c)             // Appends the character to the string builder.
	}
	return sb.String() // Returns the generated random string.
}

func RandomAccountID() int64 {
	return RandomInt(0, 10000)
}

func RandomOwner() string {
	return RandomString(6) // Generates a random owner name with 6 characters.
}

func RandomBalance() int64 {
	return RandomInt(0, 1000) // Generates a random money amount between 0 and 1000.
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "AUD", "JPY"} // List of supported currencies.
	n := len(currencies)
	return currencies[rand.Intn(n)] // Selects a random currency from the list.
}
