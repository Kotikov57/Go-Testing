// Напишите тест для вашего HTTP-сервера, проверяющий, 
// что он правильно обрабатывает запросы и возвращает ожидаемые результаты

package main

import (
	"fmt"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
