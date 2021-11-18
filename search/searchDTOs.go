package search

import (
	"github.com/gremp/essearchengine/generators/boostgenerators"
	"github.com/gremp/essearchengine/generators/resultfieldgenerators"
	"github.com/gremp/essearchengine/generators/searchffieldgenerators"
	"github.com/gremp/essearchengine/helpers"
)

type RequestOptions struct {
	Query        string                              `json:"query"`
	Page         *helpers.PageObj                    `json:"page,omitempty"`
	Sort         []map[string]SortDirection          `json:"sort,omitempty"`
	Group        *GroupOptions                       `json:"group,omitempty"`
	Facets       FacetObject                         `json:"facets,omitempty"`
	Filters      interface{}                         `json:"filters,omitempty"`
	Precision    int                                 `json:"precision,omitempty"`
	Boosts       boostgenerators.BoostObject         `json:"boosts,omitempty"`
	SearchFields searchffieldgenerators.SearchFields `json:"search_fields,omitempty"`
	ResultFields resultfieldgenerators.ResultsFields `json:"result_fields,omitempty"`
	Analytics    []string                            `json:"analytics,omitempty"`
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
