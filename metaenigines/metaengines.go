package metaenigines

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gremp/essearchengine/helpers"
	"net/http"
)

type MetaEngines struct {
	engineName string
	apiKey     string
	url        string
}

func New(apiKey, url string) *MetaEngines {
	return &MetaEngines{
		apiKey: apiKey,
		url:    url,
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

	return helpers.DoEngineRequest(ctx, url, this.apiKey, method, payloadBytes)
}
