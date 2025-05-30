// Напишите тест для функции, которая должна возвращать ошибку в определенных условиях

package main

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
