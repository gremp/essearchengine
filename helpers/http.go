package helpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)
var ErrGotHttpRequestError = errors.New("got http request error")

func DoEngineRequest(ctx context.Context, url, apiKey, method string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := http.DefaultClient

	response, err :=  client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		data, _ := io.ReadAll(response.Body)
		err := fmt.Errorf("%w with status: %d, body response was : %s", ErrGotHttpRequestError, response.StatusCode, string(data))

		return nil, err
	}

	return response, nil
}

func BuildURL(baseURL, engineName, endpoint string) string {
	urlTemplate :="%s/api/as/v1/engines/%s/%s"
	return fmt.Sprintf(urlTemplate, baseURL, engineName, endpoint)
}