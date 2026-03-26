package controllers

import (
	"net/http"

	api "github.com/brunobotter/map/api/http"
	"github.com/brunobotter/map/application/usecase"

	"github.com/brunobotter/map/api/requests"
)

type MapHandler struct {
	mapUsecase usecase.MapUsecase
}

func NewMapHandler(mapUsecase usecase.MapUsecase) *MapHandler {
	return &MapHandler{mapUsecase: mapUsecase}
}

func (h *MapHandler) MapData(reque requests.MapRequest) *api.HttpResponse {
	result, err := h.mapUsecase.Execute(reque.Context())
	if err != nil {
		return &api.HttpResponse{
			StatusCode: http.StatusOK,
			Body: map[string]any{
				"weather": map[string]any{},
				"traffic": []any{},
				"events":  []any{},
			},
		}
	}

	return &api.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       result,
	}
}
