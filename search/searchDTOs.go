package search

type SortDirection string

var DirectionAsc SortDirection = "asc"
var DirectionDesc SortDirection = "desc"

type GroupOptions struct {
	Field    string                   `json:"field"`
	Size     int                      `json:"size,omitempty"`
	Sort     map[string]SortDirection `json:"sort,omitempty"`
	Collapse bool                     `json:"collapse,omitempty"`
}

type FacetOption struct {
	Type string                   `json:"type"`
	Name string                   `json:"name,omitempty"`
	Sort map[string]SortDirection `json:"sort,omitempty"`
	Size int                      `json:"size,omitempty"`
}

type FacetOptions []*FacetOption
type FacetObject map[string]FacetOptions
