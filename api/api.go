package api

type StandardReturn struct {
	Type       string `json:"type"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}
