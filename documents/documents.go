package documents

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gremp/essearchengine/helpers"
)

var ErrListError = errors.New("errors on documents/list")

type Documents struct {
	engineName string
	apiKey     string
	url        string
}

func New(engineName, apiKey, url string) *Documents {
	return &Documents{
		engineName: engineName,
		apiKey:     apiKey,
		url:        url,
	}
}

func (this *Documents) Create(ctx context.Context, documents interface{}) (documentErrors *UpsertDocumentErrors, requestError error) {
	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return nil, err
	}

	return this.upsert(ctx, documents, http.MethodPost)
}

func (this *Documents) Update(ctx context.Context, documents interface{}) (documentErrors *UpsertDocumentErrors, requestError error) {
	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return nil, err
	}

	return this.upsert(ctx, documents, http.MethodPatch)
}

func (this *Documents) Delete(ctx context.Context, documentIDs []string) (unableToDeleteIDs []string, requestError error) {
	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return nil, err
	}

	resp, err := this.sendRequest(ctx, documentIDs, http.MethodDelete)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return this.filterDeleteErrors(resp.Body)
}

func (this *Documents) Get(ctx context.Context, documentIDs []string, target interface{}) error {
	if err := this.checkError(target); err != nil {
		return err
	}

	resp, err := this.sendRequest(ctx, documentIDs, http.MethodGet)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func (this *Documents) List(ctx context.Context, page, size int, target interface{}) (*helpers.ResultMeta, error) {
	if err := this.checkError(target); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("documents/list?page[current]=%d&page[size]=%d", page, size)

	resp, err := this.sendRequest(ctx, nil, http.MethodGet, endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	payload := &helpers.GenericSearchResponse{}
	payload.Results = target

	if err = json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if len(payload.Errors) != 0 {
		err = fmt.Errorf("%w. Errors: %s", ErrListError, strings.Join(payload.Errors, ", "))

		return nil, err
	}

	return &payload.Meta, nil
}

func (this *Documents) upsert(ctx context.Context, documents interface{}, method string) (documentErrors *UpsertDocumentErrors, requestError error) {
	resp, err := this.sendRequest(ctx, documents, method)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return this.filterUpsertErrors(resp.Body)
}

func (this *Documents) sendRequest(ctx context.Context, documents interface{}, method string, overideEndpoint ...string) (*http.Response, error) {
	endpoint := "documents"
	if len(overideEndpoint) == 1 {
		endpoint = overideEndpoint[0]
	}
	url := helpers.BuildURL(this.url, this.engineName, endpoint)

	payload, err := json.Marshal(documents)
	if err != nil {
		return nil, err
	}

	return helpers.DoEngineRequest(ctx, url, this.apiKey, method, payload)
}

func (this *Documents) filterDeleteErrors(body io.ReadCloser) (unableToDeleteIDs []string, requestError error) {
	response := make([]DeleteResponseItem, 0)
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	for _, item := range response {
		if !item.Deleted {
			unableToDeleteIDs = append(unableToDeleteIDs, item.ID)
		}
	}

	return unableToDeleteIDs, nil
}

func (this *Documents) filterUpsertErrors(body io.ReadCloser) (*UpsertDocumentErrors, error) {
	response := make([]UpsertResponseItem, 0)
	docErrors := &UpsertDocumentErrors{}

	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return nil, err
	}

	for _, item := range response {
		if len(item.Errors) > 0 {
			docErrors.Total++
			docErrors.Errors = append(docErrors.Errors, item)
		}
	}

	if docErrors.Total > 0 {
		return docErrors, nil
	}
	return nil, nil
}

func (this *Documents) checkError(target interface{}) error {
	if err := helpers.ValidateOptions(this.engineName, this.url, this.apiKey); err != nil {
		return err
	}

	if err := helpers.CheckPointer(target); err != nil {
		return err
	}
	return nil
}
