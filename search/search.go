package search

import (
	"context"
	"encoding/json"
	"github.com/gremp/essearchengine/generators/boostgenerators"
	"github.com/gremp/essearchengine/generators/filtergenerators"
	"github.com/gremp/essearchengine/generators/resultfieldgenerators"
	"github.com/gremp/essearchengine/generators/searchffieldgenerators"
	"github.com/gremp/essearchengine/helpers"
	"net/http"
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

	this.requestOptions.Facets[field] = facetOptions

	return this
}

func (this *Search) Filters(filters *filtergenerators.Filters) *Search {
	this.requestOptions.Filters = filters.GetFilter()
	return this
}

func (this *Search) Precision(precision int) *Search {
	this.requestOptions.Precision = precision

	return this
}

func (this *Search) SearchField(field string, searchField *searchffieldgenerators.SingleSearchField) *Search {
	if this.requestOptions.SearchFields == nil {
		searchFieldsMain := make(searchffieldgenerators.SearchFields)
		this.requestOptions.SearchFields = searchFieldsMain
	}

	this.requestOptions.SearchFields[field] = searchField
	return this
}

func (this *Search) ResultField(field string, resultField *resultfieldgenerators.SingleResultField) *Search {
	if this.requestOptions.ResultFields == nil {
		this.requestOptions.ResultFields = make(resultfieldgenerators.ResultsFields)
	}
	this.requestOptions.ResultFields[field] = resultField

	return this
}

func (this *Search) Analytics(tags ...string) *Search {
	this.requestOptions.Analytics = tags

	return this
}

func (this *Search) Boost(field string, boost boostgenerators.BoostSingleObject) *Search {
	if this.requestOptions.Boosts == nil {
		this.requestOptions.Boosts = make(boostgenerators.BoostObject)
	}

	this.requestOptions.Boosts[field] = boost

	return this
}

func (this *Search) Do(ctx context.Context, target interface{}) (*helpers.ResultMeta, error) {

	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return nil, err
	}

	if err := helpers.CheckPointer(target); err != nil {
		return nil, err
	}

	payloadBytes, err := json.Marshal(this.requestOptions)
	if err != nil {
		return nil, err
	}

	url := helpers.BuildURL(this.url, this.engineName, "search")
	response, err := helpers.DoEngineRequest(ctx, url, this.apiKey, http.MethodPost, payloadBytes)
	if err != nil {
		return nil, err
	}

	payload := &helpers.GenericSearchResponse{}
	payload.Results = target

	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload.Meta, nil
}

func (this *Search) GetRequestOptions() *RequestOptions {
	return this.requestOptions
}

func (this *Search) Clone() (*Search, error) {
	res, err := json.Marshal(this.requestOptions)
	if err != nil {
		return nil, err
	}

	search := New(this.engineName, this.apiKey, this.url)
	requestOptions := &RequestOptions{}

	err = json.Unmarshal(res, requestOptions)
	if err != nil {
		return nil, err
	}

	search.SetRequestOptions(requestOptions)

	return search, nil
}

func (this *Search) SetRequestOptions(options *RequestOptions) {
	this.requestOptions = options
}
