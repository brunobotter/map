package domain

type Weather struct {
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
}

type Traffic struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Type string  `json:"type"`
}

type MapData struct {
	Weather Weather   `json:"weather"`
	Traffic []Traffic `json:"traffic"`
	Events  []Event   `json:"events"`
}
