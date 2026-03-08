package requests

import "github.com/brunobotter/map/api/http"

type CreateTenantRequest struct {
	http.HttpRequest
	Name string `json:"name"`
}

type TenantRequest struct {
	http.HttpRequest
	Name string `json:"name"`
}
