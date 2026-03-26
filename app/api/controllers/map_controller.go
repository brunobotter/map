package controllers

import (
	api "github.com/brunobotter/map/api/http"
	"github.com/brunobotter/map/application/usecase"

	"github.com/brunobotter/map/api/requests"
)

type MapHandler struct {
	mapUsecase usecase.MapUsecase
}

func NewMapHandler() *MapHandler {
	return &MapHandler{}
}

func (h *MapHandler) MapData(reque requests.MapRequest) *api.HttpResponse {
	return api.Ok("")
}
