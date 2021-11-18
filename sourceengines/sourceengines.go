package sourceengines

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gremp/essearchengine/helpers"
	"net/http"
)

type SourceEngines struct {
	engineName string
	apiKey     string
	url        string
}

func New(apiKey, url string) *SourceEngines {
	return &SourceEngines{
		apiKey: apiKey,
		url:    url,
	}
}

func (this *SourceEngines) Get(ctx context.Context, engineName string) (*GenericEnginesRes, error) {
	if err := helpers.ValidateOptions("always_true", this.url, this.apiKey); err != nil {
		return nil, err
	}

	resp, err := this.sendRequest(ctx, nil, http.MethodGet, engineName)
	if err != nil {
		return nil, err
	}

	return this.decodeGenericResponse(resp)
}

func (this *SourceEngines) Create(ctx context.Context, engineName string, language Language) (*GenericEnginesRes, error) {
	if err := helpers.ValidateOptions("always_true", this.url, this.apiKey); err != nil {
		return nil, err
	}

	request := &CreateEngineReq{
		Name: engineName,
	}

	if language != "" {
		request.Language = language
	}

	resp, err := this.sendRequest(ctx, request, http.MethodPost)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return this.decodeGenericResponse(resp)
}

func (this *SourceEngines) List(ctx context.Context, page int, size int) (*ListEnginesRes, error) {

	endpoint := fmt.Sprintf("?page[current]=%d&page[size]=%d", page, size)

	resp, err := this.sendRequest(ctx, nil, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &ListEnginesRes{}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (this *SourceEngines) Delete(ctx context.Context, engineName string) (*DeleteRes, error) {
	resp, err := this.sendRequest(ctx, nil, http.MethodDelete, engineName)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &DeleteRes{}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (this *SourceEngines) decodeGenericResponse(resp *http.Response) (*GenericEnginesRes, error) {
	response := &GenericEnginesRes{}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (this *SourceEngines) sendRequest(ctx context.Context, payload interface{}, method string, overideEndpoint ...string) (*http.Response, error) {
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
