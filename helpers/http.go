package helpers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

	return client.Do(req)
}

func BuildURL(baseURL, engineName, endpoint string) string {
	urlTemplate := "%s/api/as/v1/engines/%s/%s"
	return fmt.Sprintf(urlTemplate, baseURL, engineName, endpoint)
}
