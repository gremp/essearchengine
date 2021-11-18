package helpers

import "fmt"

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
	RequestID string `json:"request_id"`
}

type GenericSearchResponse struct {
	Errors  []string    `json:"errors,omitempty"`
	Meta    ResultMeta  `json:"meta,omitempty"`
	Results interface{} `json:"results,omitempty"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (this *GeoPoint) GetStr() string {
	return fmt.Sprintf("%f, %f", this.Lat, this.Lon)
}
