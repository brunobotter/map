package response

import "time"

type MapResponse struct {
	TenantId string
	Name     string
	CreateAt time.Time
}
