package utils

import "math/rand"

// RandomCurrency generates a valid random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
