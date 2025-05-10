// Используйте моки для тестирования функции, которая взаимодействует с внешним сервисом

package main

import (
	"errors"
	"net/http"
)

func ExtApiReq(client *http.Client, url string) (*http.Response, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.New("connection error")
	}
	return resp, nil
}
