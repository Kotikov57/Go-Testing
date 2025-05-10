// Напишите тест для HTTP-обработчика, проверяющего код ответа и содержимое ответа

package main

import (
	"fmt"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Test message")
}
