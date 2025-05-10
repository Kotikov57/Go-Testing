// Используйте testing пакет для тестирования функции, зависящей от времени

package main

import (
	"time"
)

func TimeFunc(t time.Time) string {
	hour := t.Hour()
	if hour < 12 {
		return "Доброе утро"
	} else if hour < 18 {
		return "Добрый день"
	}
	return "Добрый вечер"
}
