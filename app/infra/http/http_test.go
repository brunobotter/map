package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brunobotter/map/main/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHttpClientDoRetriesOnServerError(t *testing.T) {
	attempts := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	client := NewHttpClient(&config.Config{Env: "test"}, nil)
	req, err := client.NewRequest(http.MethodGet, ts.URL, nil)
	require.NoError(t, err)

	resp, err := client.Do(context.Background(), "test", req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.Status())
	assert.Equal(t, 2, attempts)
}

func TestHttpClientDoTimeoutCancelsRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := NewHttpClient(&config.Config{Env: "test"}, nil)
	req, err := client.NewRequest(http.MethodGet, ts.URL, nil)
	require.NoError(t, err)

	start := time.Now()
	_, err = client.Do(context.Background(), "test", req)
	require.Error(t, err)
	assert.Less(t, time.Since(start), 6*time.Second)
}
