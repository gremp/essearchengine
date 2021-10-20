package metaenigines

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gremp/essearchengine/helpers"
)

var (
	ErrEngineDoesNotExist = errors.New("engine does not exist")
)

type MetaEngines struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *MetaEngines {
	return &MetaEngines{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *MetaEngines) Create(ctx context.Context, metaEngineName string, sourceEngines []string) (*GenericSourceEnginesRes, error) {
	if err := helpers.ValidateOptions("always_true", this.url, this.apiKey); err != nil {
		return nil, err
	}

	request := &CreateEngineReq{
		Name:          metaEngineName,
		Type:          "meta",
		SourceEngines: sourceEngines,
	}

	resp, err := this.sendRequest(ctx, request, http.MethodPost)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &GenericSourceEnginesRes{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (this *MetaEngines) GetInfo(ctx context.Context, metaEngineName string) (*GenericSourceEnginesRes, error) {
	resp, err := this.sendRequest(ctx, nil, http.MethodGet, metaEngineName)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &GenericSourceEnginesRes{}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (this *MetaEngines) AddEngines(ctx context.Context, metaEngineName string, sourceEngines []string) (*GenericSourceEnginesRes, error) {
	return this.modifyEngines(ctx, metaEngineName, sourceEngines, http.MethodPost)
}

func (this *MetaEngines) DeleteEngines(ctx context.Context, metaEngineName string, sourceEngines []string) (*GenericSourceEnginesRes, error) {
	return this.modifyEngines(ctx, metaEngineName, sourceEngines, http.MethodDelete)
}

func (this *MetaEngines) modifyEngines(ctx context.Context, metaEngineName string, sourceEngines []string, method string) (*GenericSourceEnginesRes, error) {
	endpoint := fmt.Sprintf("%s/source_engines", metaEngineName)
	resp, err := this.sendRequest(ctx, sourceEngines, method, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := &GenericSourceEnginesRes{}

	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (this *MetaEngines) sendRequest(ctx context.Context, payload interface{}, method string, overideEndpoint ...string) (*http.Response, error) {
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

		if helpers.IsStringInSplice(errorResponse.Errors, ErrEngineDoesNotExist.Error()) {
			return nil, ErrEngineDoesNotExist
		}


		err = fmt.Errorf("%w with status: %d, body response was : %s", helpers.ErrGotHttpRequestError, response.StatusCode, strings.Join(errorResponse.Errors, ", "))

		return nil, err
	}
	return response, nil
}
