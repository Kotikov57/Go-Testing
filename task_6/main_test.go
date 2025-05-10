// Используйте моки для тестирования функции, которая взаимодействует с внешним сервисом

package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtApiReqSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":"ok"}`)
	}))
	defer server.Close()
	client := &http.Client{}
	resp, err := ExtApiReq(client, server.URL)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestExtApiReqFail(t *testing.T) {
	client := &http.Client{}
	resp, err := ExtApiReq(client, "http://localhost:0")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, "connection error", err.Error())
}
