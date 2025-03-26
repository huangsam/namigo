package extern

// PyPIAPIListResponse represents list info from the PyPI API.
type PyPIAPIListResponse struct {
	Meta struct {
		LastSerial int    `json:"_last-serial"`
		APIVersion string `json:"api-version"`
	} `json:"meta"`
	Projects []struct {
		LastSerial int    `json:"_last-serial"`
		Name       string `json:"name"`
	} `json:"projects"`
}

// PyPIAPIDetailResponse represents detailed info from the PyPI API.
type PyPIAPIDetailResponse struct {
	Info struct {
		Author      string `json:"author"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Version     string `json:"version"`
	} `json:"info"`
}
