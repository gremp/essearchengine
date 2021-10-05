package appsearch

import (
	"errors"
	"sync"

	"github.com/gremp/essearchengine/documents"
	"github.com/gremp/essearchengine/search"
)

type AppSearch struct {
	defaultEngineName string
	url               string
	apiKey            string
}

type OverideOptions struct {
	EngineName string
	URL        string
	APIKey     string
}

var (
	once                    sync.Once
	instance                *AppSearch
	ErrHasNotInitialized    = errors.New("lib is not initialized. Run Init first")
	ErrMissingConfiguration = errors.New("missing either url or engine name or api key")
)

func GetInstance() (*AppSearch, error) {
	if instance == nil {
		return nil, ErrHasNotInitialized
	}

	return instance, nil
}

func Init(url, defaultEngineName, apiKey string) *AppSearch {

	once.Do(func() {
		instance = &AppSearch{
			defaultEngineName: defaultEngineName,
			url:               url,
			apiKey:            apiKey,
		}

	})

	return instance
}

// Documents creates a new Documents instance for inserting/updating/deleting/getting a document
// API used https://www.elastic.co/guide/en/app-search/7.15/documents.html
// Options is an object that can be passed to overide the default configuration of engine name api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Documents(options ...OverideOptions) *documents.Documents {
	engineName, url, apiKey := this.getOptions(options...)

	return documents.New(engineName, apiKey, url)
}

// Search creates a new Search instance for searching documents
// API used https://www.elastic.co/guide/en/app-search/7.15/search.html
// Options is an object that can be passed to overide the default configuration of engine name api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Search(options ...OverideOptions) *search.Search {
	engineName, url, apiKey := this.getOptions(options...)

	return search.New(engineName, apiKey, url)
}

func (this *AppSearch) getOptions(options ...OverideOptions) (engineName, url, apiKey string) {
	engineName = this.defaultEngineName
	url = this.url
	apiKey = this.apiKey

	if len(options) > 0 {
		if options[0].URL != "" {
			url = options[0].URL
		}

		if options[0].APIKey != "" {
			apiKey = options[0].APIKey
		}

		if options[0].EngineName != "" {
			engineName = options[0].EngineName
		}
	}

	if engineName == "" || url == "" || apiKey == "" {
		panic(ErrMissingConfiguration)
	}
	return engineName, url, apiKey
}
