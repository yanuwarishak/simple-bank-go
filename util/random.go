package util

import (
	"fmt"
	"math/rand"
	"strings"
)

// RandomInt generate random integer
func RandomInt(min, max int64) int64 {
	if min >= max {
		panic("invalid range")
	}
	return min + rand.Int63n(max-min+1) // Generates a random integer between min and max (inclusive).
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generate random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // Selects a random character from the alphabet.
		sb.WriteByte(c)             // Appends the character to the string builder.
	}
	return sb.String() // Returns the generated random string.
}

// RandomAccountID generate ID ranging from min value to max value
func RandomAccountID() int64 {
	return RandomInt(0, 10000)
}

// RandomOwner generate owner string name with 6 characteres
func RandomOwner() string {
	return RandomString(6)
}

// RandomBalance generates a random money amount between 0 and 1000.
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency selects a random currency from the list.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "AUD", "JPY"} // List of supported currencies.
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generate random string that follows email pattern
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
