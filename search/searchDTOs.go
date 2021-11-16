package search

import "github.com/gremp/essearchengine/helpers"

type RequestOptions struct {
	Query        string                     `json:"query"`
	Page         *helpers.PageObj           `json:"page,omitempty"`
	Sort         []map[string]SortDirection `json:"sort,omitempty"`
	Group        *GroupOptions              `json:"group,omitempty"`
	Facets       FacetObject                `json:"facets,omitempty"`
	Filters      FiltersObject              `json:"filters,omitempty"`
	Precision    int                        `json:"precision,omitempty"`
	Boosts       BoostsObject               `json:"boosts,omitempty"`
	SearchFields *helpers.SearchFields      `json:"search_fields,omitempty"`
	ResultFields *helpers.ResultsFields     `json:"result_fields,omitempty"`
}

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

type FiltersObject map[string]interface{}

type BoostsObject map[string]BoostOption

type BoostOption struct {
	Type      string `json:"type"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
	Factor    int    `json:"factor"`
}
