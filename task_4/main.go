// Используйте пакет testing/quick для автоматизации тестирования генератора случайных данных

package main

import (
	"math/rand"
)

const charset = "123456789"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
