package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	infraHttp "github.com/brunobotter/map/infra/http"
	"github.com/brunobotter/map/main/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenWeatherIntegrationGetWeather(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/data/2.5/weather", r.URL.Path)
		assert.Equal(t, "token", r.URL.Query().Get("appid"))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"main":{"temp":22.4},"weather":[{"main":"Clouds"}]}`))
	}))
	defer ts.Close()

	cfg := &config.Config{
		Env: "test",
		Weather: config.WeatherConfig{
			BaseURL: ts.URL,
			APIKey:  "token",
		},
	}

	client := infraHttp.NewHttpClient(cfg, nil)
	integration := NewOpenWeatherIntegration(client, cfg)

	result, err := integration.GetWeather(context.Background(), -23.55052, -46.633308)
	require.NoError(t, err)
	assert.Equal(t, 22.4, result.Temperature)
	assert.Equal(t, "Clouds", result.Status)
	assert.Equal(t, "C", result.Unit)
}
