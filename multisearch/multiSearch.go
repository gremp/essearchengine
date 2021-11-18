package multisearch

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gremp/essearchengine/helpers"
	"github.com/gremp/essearchengine/search"
	"net/http"
)

var (
	ErrQueryAndResultNotSameLength = errors.New("query length and result length not the same size")
)

type MultiSearch struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *MultiSearch {
	return &MultiSearch{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *MultiSearch) Do(ctx context.Context, searches []*search.Search, targets []interface{}) ([]*helpers.ResultMeta, error) {

	responses := make([]*helpers.GenericSearchResponse, len(searches))
	metaArr := make([]*helpers.ResultMeta, len(searches))
	payload := &MultiSearchPayload{Queries: make([]*search.RequestOptions, len(searches))}

	if len(targets) != len(searches) {
		return nil, ErrQueryAndResultNotSameLength
	}

	for i, _ := range searches {
		payload.Queries[i] = searches[i].GetRequestOptions()
		responses[i] = &helpers.GenericSearchResponse{
			Results: &targets[i],
		}
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := helpers.BuildURL(this.url, this.engineName, "multi_search")

	response, err := helpers.DoEngineRequest(ctx, url, this.apiKey, http.MethodPost, payloadBytes)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response.Body).Decode(&responses)
	if err != nil {
		return nil, err
	}

	if len(targets) != len(responses) {
		return nil, ErrQueryAndResultNotSameLength
	}

	for i, queryRes := range responses {
		metaArr[i] = &queryRes.Meta
		targets[i] = &queryRes.Results
	}

	return metaArr, err
}
