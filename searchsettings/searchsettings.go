package searchsettings

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gremp/essearchengine/helpers"
)

type SearchSettings struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *SearchSettings {
	return &SearchSettings{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *SearchSettings) Get(ctx context.Context, engineName string) (*SearchSettingsConfig, error) {
	endpoint := fmt.Sprintf("%s/search_settings", engineName)

	resp, err := this.sendRequest(ctx, nil, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return this.decodeResponse(resp)
}

func (this *SearchSettings) Reset(ctx context.Context, engineName string) (*SearchSettingsConfig, error) {
	endpoint := fmt.Sprintf("%s/search_settings/reset", engineName)

	resp, err := this.sendRequest(ctx, nil, http.MethodPost, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return this.decodeResponse(resp)
}

func (this *SearchSettings) Update(ctx context.Context, engineName string, searchSettingsConfig *SearchSettingsConfig) (*SearchSettingsConfig, error) {
	endpoint := fmt.Sprintf("%s/search_settings", engineName)

	resp, err := this.sendRequest(ctx, searchSettingsConfig, http.MethodPut, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return this.decodeResponse(resp)
}

func (this *SearchSettings) CreateResultSetting(showRaw bool, rawSize int, showSnippet bool, snipSize int, snipFallback bool) *SingleResultSettings {
	var raw *SingleResultSettingsRaw
	var snippet *SingleResultSettingsSnippet

	if showRaw {
		raw = &SingleResultSettingsRaw{}
		if rawSize > 0 {
			raw.Size = rawSize
		}
	}

	if showSnippet {
		snippet = &SingleResultSettingsSnippet{}
		if snipSize > 20 {
			snippet.Size = snipSize
		}

		if snipFallback {
			snippet.Fallback = snipFallback
		}
	}

	return &SingleResultSettings{
		Raw:     raw,
		Snippet: snippet,
	}
}

func (this *SearchSettings) CreateSearchSetting(weight int) *SingleFieldSettings {
	return &SingleFieldSettings{
		Weight: weight,
	}
}

func (this *SearchSettings) decodeResponse(resp *http.Response) (*SearchSettingsConfig, error) {
	response := &SearchSettingsConfig{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

func (this *SearchSettings) sendRequest(ctx context.Context, payload interface{}, method string, overideEndpoint ...string) (*http.Response, error) {
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
		data, _ := io.ReadAll(response.Body)
		err = fmt.Errorf("%w with status: %d, body response was : %s", helpers.ErrGotHttpRequestError, response.StatusCode, string(data))

		return nil, err
	}
	return response, nil
}
