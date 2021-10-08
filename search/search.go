package search

import (
	"context"

	"github.com/gremp/essearchengine/helpers"
)

type Search struct {
	engineName     string
	apiKey         string
	url            string
	requestOptions struct {
		query  string
		page   *helpers.PageObj
		sort   []map[string]SortDirection
		Group  *GroupOptions
		Facets FacetObject
	}
}

func New(engineName, apiKey, url string) *Search {
	return &Search{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *Search) Query(query string) *Search {
	this.requestOptions.query = query

	return this
}

func (this *Search) Page(current, size int) *Search {
	this.requestOptions.page = &helpers.PageObj{
		Current: current,
		Size:    size,
	}

	return this
}

func (this *Search) Sort(sortField string, sortDirection SortDirection) *Search {
	if this.requestOptions.sort == nil {
		this.requestOptions.sort = make([]map[string]SortDirection, 0)
	}
	this.requestOptions.sort = append(this.requestOptions.sort, map[string]SortDirection{sortField: sortDirection})
	return this
}

func (this *Search) Group(groupOptions *GroupOptions) *Search {
	this.requestOptions.Group = groupOptions
	return this
}

func (this *Search) Facets(field string, facetOptions FacetOptions) *Search {
	if this.requestOptions.Facets == nil {
		this.requestOptions.Facets = make(FacetObject)
	}
	// this.requestOptions.Facets[field] = append(this.requestOptions.Facets[field], facetOptions)

	return this
}

func (this *Search) Do(ctx context.Context, target interface{}) (*helpers.ResultMeta, error) {
	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return nil, err
	}

	if err := helpers.CheckPointer(target); err != nil {
		return nil, err
	}

	return nil, nil
}
