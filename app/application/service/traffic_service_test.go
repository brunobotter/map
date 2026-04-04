package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrafficServiceGetIncidents(t *testing.T) {
	svc := NewTrafficService()

	result, err := svc.GetIncidents(context.Background(), -23.55052, -46.633308)
	require.NoError(t, err)
	require.Len(t, result, 2)

	assert.Equal(t, "accident", result[0].Type)
	assert.Equal(t, "jam", result[1].Type)
	assert.NotZero(t, result[0].Lat)
	assert.NotZero(t, result[0].Lng)
}
