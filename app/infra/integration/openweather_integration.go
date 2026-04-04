package integration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/brunobotter/map/application/domain"
	httpContract "github.com/brunobotter/map/application/http"
	appIntegration "github.com/brunobotter/map/application/integration"
	"github.com/brunobotter/map/main/config"
)

const defaultOpenWeatherBaseURL = "https://api.openweathermap.org"

type openWeatherIntegration struct {
	httpClient httpContract.Client
	cfg        *config.Config
}

func NewOpenWeatherIntegration(httpClient httpContract.Client, cfg *config.Config) appIntegration.WeatherIntegration {
	return &openWeatherIntegration{httpClient: httpClient, cfg: cfg}
}

func (s *openWeatherIntegration) GetWeather(ctx context.Context, lat, lng float64) (domain.Weather, error) {
	if s.cfg.Weather.APIKey == "" {
		return domain.Weather{}, fmt.Errorf("missing openweather api key")
	}

	baseURL := normalizeBaseURL(s.cfg.Weather.BaseURL)
	endpoint := fmt.Sprintf("%s/data/2.5/weather", baseURL)
	query := url.Values{}
	query.Set("lat", strconv.FormatFloat(lat, 'f', 6, 64))
	query.Set("lon", strconv.FormatFloat(lng, 'f', 6, 64))
	query.Set("appid", s.cfg.Weather.APIKey)
	query.Set("units", "metric")

	req, err := s.httpClient.NewRequestWithContext(ctx, http.MethodGet, endpoint+"?"+query.Encode(), nil)
	if err != nil {
		return domain.Weather{}, err
	}

	resp, err := s.httpClient.Do(ctx, "openweather", req)
	if err != nil {
		return domain.Weather{}, err
	}

	if resp.Status() != http.StatusOK {
		return domain.Weather{}, fmt.Errorf("openweather returned status %d", resp.Status())
	}

	body, err := resp.Body()
	if err != nil {
		return domain.Weather{}, err
	}

	payload := struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Main string `json:"main"`
		} `json:"weather"`
	}{}

	if err = json.Unmarshal(body, &payload); err != nil {
		return domain.Weather{}, err
	}

	condition := "unknown"
	if len(payload.Weather) > 0 {
		condition = payload.Weather[0].Main
	}

	return domain.Weather{Temperature: payload.Main.Temp, Status: condition, Unit: "C"}, nil
}
func normalizeBaseURL(baseURL string) string {
	trimmed := strings.TrimSpace(baseURL)
	if trimmed == "" {
		return defaultOpenWeatherBaseURL
	}

	if !strings.HasPrefix(trimmed, "http://") && !strings.HasPrefix(trimmed, "https://") {
		trimmed = "https://" + trimmed
	}

	return strings.TrimRight(trimmed, "/")
}
