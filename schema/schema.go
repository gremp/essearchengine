package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gremp/essearchengine/helpers"
)

type Schema struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *Schema {
	return &Schema{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *Schema) Get(ctx context.Context, engineName string) (map[string]SchemaType, error) {
	endpoint := fmt.Sprintf("%s/%s", engineName, "schema")
	resp, err := this.sendRequest(ctx, nil, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	return this.decodeResponse(resp)
}

func (this *Schema) Update(ctx context.Context, engineName string, updateValues map[string]SchemaType) (map[string]SchemaType, error) {
	endpoint := fmt.Sprintf("%s/%s", engineName, "schema")
	resp, err := this.sendRequest(ctx, updateValues, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	return this.decodeResponse(resp)
}

func (this *Schema) decodeResponse(resp *http.Response) (map[string]SchemaType, error) {
	response := make(map[string]SchemaType)

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil

}

func (this *Schema) sendRequest(ctx context.Context, payload interface{}, method string, overideEndpoint ...string) (*http.Response, error) {
	url := fmt.Sprintf("%s/api/as/v1/engines", this.url)
	if len(overideEndpoint) == 1 {
		url = fmt.Sprintf("%s/%s", url, overideEndpoint[0])
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return helpers.DoEngineRequest(ctx, url, this.apiKey, method, payloadBytes)
}
