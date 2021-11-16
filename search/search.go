package search

import (
	"context"
	"github.com/gremp/essearchengine/helpers"
)

type Search struct {
	engineName     string
	apiKey         string
	url            string
	requestOptions *RequestOptions
}

func New(engineName, apiKey, url string) *Search {
	return &Search{
		engineName:     engineName,
		apiKey:         apiKey,
		url:            url,
		requestOptions: &RequestOptions{},
	}
}

func (this *Search) Query(query string) *Search {
	this.requestOptions.Query = query

	return this
}

func (this *Search) Page(current, size int) *Search {
	this.requestOptions.Page = &helpers.PageObj{
		Current: current,
		Size:    size,
	}

	return this
}

func (this *Search) Sort(sortField string, sortDirection SortDirection) *Search {
	if this.requestOptions.Sort == nil {
		this.requestOptions.Sort = make([]map[string]SortDirection, 0)
	}
	this.requestOptions.Sort = append(this.requestOptions.Sort, map[string]SortDirection{sortField: sortDirection})
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

	//payloadBytes, err := json.Marshal(payload)
	//if err != nil {
	//	return nil, err
	//}
	//
	//response, err := helpers.DoEngineRequest(ctx, url, this.apiKey, method, payloadBytes)

	return nil, nil
}

func (this *Search) GetRequestOptions() *RequestOptions {
	return this.requestOptions
}
