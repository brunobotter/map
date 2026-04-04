package service

import (
	"context"
	"errors"
	"testing"

	"github.com/brunobotter/map/application/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type eventsIntegrationStub struct {
	result []domain.Event
	err    error
}

func (e *eventsIntegrationStub) GetEvents(ctx context.Context, lat, lng float64) ([]domain.Event, error) {
	_ = ctx
	_ = lat
	_ = lng
	if e.err != nil {
		return nil, e.err
	}
	return e.result, nil
}

func TestEventsServiceGetEvents(t *testing.T) {
	expected := []domain.Event{{Name: "Feira", Lat: -23.5, Lng: -46.6}}
	svc := NewEventsService(&eventsIntegrationStub{result: expected})

	result, err := svc.GetEvents(context.Background(), -23.55052, -46.633308)
	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestEventsServiceGetEventsError(t *testing.T) {
	svc := NewEventsService(&eventsIntegrationStub{err: errors.New("api unavailable")})

	_, err := svc.GetEvents(context.Background(), -23.55052, -46.633308)
	require.Error(t, err)
}
