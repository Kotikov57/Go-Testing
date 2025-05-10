// Напишите тест для вашего HTTP-сервера, проверяющий, 
// что он правильно обрабатывает запросы и возвращает ожидаемые результаты

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(SimpleHandler))
	defer server.Close()
	resp, err := http.Get(server.URL + "?name=Alice")
	require.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, "Hello, Alice!", string(body))
}

func TestHelloHandler_MissingName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	SimpleHandler(rr, req)
	require.Equal(t, http.StatusBadRequest, rr.Code)
	require.Equal(t, "name is required\n", rr.Body.String())
}
