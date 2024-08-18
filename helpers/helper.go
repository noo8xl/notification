package helpers

import "math/rand"

// CreateClientKey -> create unique hash for client
func CreateClientKey(length int) string {

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
