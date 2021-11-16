package helpers

type PageObj struct {
	Current int `json:"current,omitempty"`
	Size    int `json:"size,omitempty"`
}

type ResultMeta struct {
	Page struct {
		Current      int `json:"current"`
		TotalPages   int `json:"total_pages"`
		TotalResults int `json:"total_results"`
		Size         int `json:"size"`
	} `json:"page"`
}

type GenericSearchResponse struct {
	Errors  []string    `json:"errors,omitempty"`
	Meta    ResultMeta  `json:"meta,omitempty"`
	Results interface{} `json:"results,omitempty"`
}

type EngineErrorResponse struct {
	Errors []string `json:"errors,omitempty"`
}

type SearchFields map[string]*SingleFieldSettings
type ResultsFields map[string]*SingleResultSettings

type SingleFieldSettings struct {
	Weight int `json:"weight"`
}

type SingleResultSettings struct {
	Raw     *SingleResultSettingsRaw     `json:"raw,omitempty"`
	Snippet *SingleResultSettingsSnippet `json:"snippet,omitempty"`
}

type SingleResultSettingsRaw struct {
	Size int `json:"size,omitempty"`
}
type SingleResultSettingsSnippet struct {
	Size     int  `json:"size"`
	Fallback bool `json:"fallback,omitempty"`
}
