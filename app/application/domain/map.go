package domain

type Weather struct {
	Status      string  `json:"status"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
}

type Traffic struct {
	Road   string `json:"road"`
	Level  string `json:"level"`
	Status string `json:"status"`
}

type MapEvent struct {
	Title    string `json:"title"`
	Location string `json:"location"`
	StartAt  string `json:"start_at"`
}

type MapData struct {
	Weather Weather    `json:"weather"`
	Traffic []Traffic  `json:"traffic"`
	Events  []MapEvent `json:"events"`
}
