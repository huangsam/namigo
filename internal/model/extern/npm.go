package extern

import "time"

// NPMAPIListingResponse represents listing info from the NPM API.
type NPMAPIListingResponse struct {
	Objects []struct {
		Downloads struct {
			Monthly int `json:"monthly"`
			Weekly  int `json:"weekly"`
		} `json:"downloads"`
		Dependents  int       `json:"dependents"`
		Updated     time.Time `json:"updated"`
		SearchScore float64   `json:"searchScore"`
		Package     struct {
			Name        string        `json:"name"`
			Keywords    []interface{} `json:"keywords"`
			Version     string        `json:"version"`
			Description string        `json:"description"`
			Publisher   struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"publisher"`
			Maintainers []struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"maintainers"`
			License string    `json:"license"`
			Date    time.Time `json:"date"`
			Links   struct {
				Homepage   string `json:"homepage"`
				Repository string `json:"repository"`
				Bugs       string `json:"bugs"`
				Npm        string `json:"npm"`
			} `json:"links"`
		} `json:"package"`
		Score struct {
			Final  float64 `json:"final"`
			Detail struct {
				Popularity  int `json:"popularity"`
				Quality     int `json:"quality"`
				Maintenance int `json:"maintenance"`
			} `json:"detail"`
		} `json:"score"`
		Flags struct {
			Insecure int `json:"insecure"`
		} `json:"flags"`
	} `json:"objects"`
	Total int       `json:"total"`
	Time  time.Time `json:"time"`
}
