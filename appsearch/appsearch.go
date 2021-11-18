package appsearch

import (
	"errors"
	"github.com/gremp/essearchengine/multisearch"
	"sync"

	"github.com/gremp/essearchengine/documents"
	"github.com/gremp/essearchengine/metaenigines"
	"github.com/gremp/essearchengine/schema"
	"github.com/gremp/essearchengine/search"
	"github.com/gremp/essearchengine/searchsettings"
	"github.com/gremp/essearchengine/sourceengines"
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
	once                 sync.Once
	instance             *AppSearch
	ErrHasNotInitialized = errors.New("lib is not initialized. Run Init first")
)

func GetInstance() (*AppSearch, error) {
	if instance == nil {
		return nil, ErrHasNotInitialized
	}

	return instance, nil
}

func Init(url, apiKey string) *AppSearch {

	once.Do(func() {
		instance = &AppSearch{
			url:    url,
			apiKey: apiKey,
		}

	})

	return instance
}

// Documents creates a new Documents instance for inserting/updating/deleting/getting a document
// API used https://www.elastic.co/guide/en/app-search/7.15/documents.html
// Options is an object that can be passed to overide the default configuration of engine name api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Documents(engineName string) *documents.Documents {
	return documents.New(engineName, this.apiKey, this.url)
}

// Search creates a new Search instance for searching documents
// API used https://www.elastic.co/guide/en/app-search/7.15/search.html
// Options is an object that can be passed to overide the default configuration of engine name api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Search(engineName string) *search.Search {
	return search.New(engineName, this.apiKey, this.url)
}

// MetaEngines creates a new MetaEngines instance for creating/updating/deleting meta engines
// API used https://www.elastic.co/guide/en/app-search/7.15/meta-engines.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) MetaEngines() *metaenigines.MetaEngines {
	return metaenigines.New(this.apiKey, this.url)
}

// SourceEngines creates a new SourceEngines instance for creating/updating/deleting meta engines
// API used https://www.elastic.co/guide/en/app-search/7.15/engines.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) SourceEngines() *sourceengines.SourceEngines {
	return sourceengines.New(this.apiKey, this.url)
}

// Schema creates a new Schema instance for getting and editing a schema of an engine
// API used https://www.elastic.co/guide/en/app-search/7.15/schema.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Schema(engineName string) *schema.Schema {
	return schema.New(engineName, this.apiKey, this.url)
}

// SearchSettings creates a new SearchSettings instance for getting and editing search settings, result settings, precision and boost
// API used https://www.elastic.co/guide/en/app-search/7.15/search-settings.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) SearchSettings(engineName string) *searchsettings.SearchSettings {
	return searchsettings.New(engineName, this.apiKey, this.url)
}

func (this *AppSearch) MultiSearch(engineName string) *multisearch.MultiSearch {
	return multisearch.New(engineName, this.apiKey, this.url)
}
