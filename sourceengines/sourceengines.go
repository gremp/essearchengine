package sourceengines

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gremp/essearchengine/helpers"
)

type SourceEngines struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *SourceEngines {
	return &SourceEngines{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
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

	response, err := helpers.DoEngineRequest(ctx, url, this.apiKey, method, payloadBytes)
	if response.StatusCode >= 400 {
		defer response.Body.Close()
		errorResponse := &helpers.EngineErrorResponse{}

		err = json.NewDecoder(response.Body).Decode(errorResponse)
		if err != nil {
			return nil, err
		}

		if helpers.IsStringInSplice(errorResponse.Errors, "Could not find engine.") {
			return nil, nil
		}

		err = fmt.Errorf("%w with status: %d, body response was : %s", helpers.ErrGotHttpRequestError, response.StatusCode, strings.Join(errorResponse.Errors, ", "))

		return nil, err
	}
	return response, nil
}
