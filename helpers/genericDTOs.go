package helpers

type PageObj struct {
	Current int `json:"current,omitempty"`
	Size    int `json:"size,omitempty"`
}

type ResultMeta struct {
	Page struct {
		Current int `json:"current"`
		TotalPages int `json:"total_pages"`
		TotalResults int `json:"total_results"`
		Size int `json:"size"`
	} `json:"page"`
}

type GenericSearchResponse struct {
	Errors  []string    `json:"errors,omitempty"`
	Meta    ResultMeta  `json:"meta,omitempty"`
	Results interface{} `json:"results,omitempty"`
}

