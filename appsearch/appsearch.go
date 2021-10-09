package appsearch

import (
	"errors"
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
func (this *AppSearch) Documents(options ...*OverideOptions) *documents.Documents {
	engineName, url, apiKey := this.getOptions(options...)

	return documents.New(engineName, apiKey, url)
}

// Search creates a new Search instance for searching documents
// API used https://www.elastic.co/guide/en/app-search/7.15/search.html
// Options is an object that can be passed to overide the default configuration of engine name api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Search(options ...*OverideOptions) *search.Search {
	engineName, url, apiKey := this.getOptions(options...)

	return search.New(engineName, apiKey, url)
}

// MetaEngines creates a new MetaEngines instance for creating/updating/deleting meta engines
// API used https://www.elastic.co/guide/en/app-search/7.15/meta-engines.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) MetaEngines(options ...*OverideOptions) *metaenigines.MetaEngines {
	engineName, url, apiKey := this.getOptions(options...)

	return metaenigines.New(engineName, apiKey, url)
}

// SourceEngines creates a new SourceEngines instance for creating/updating/deleting meta engines
// API used https://www.elastic.co/guide/en/app-search/7.15/engines.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) SourceEngines(options ...*OverideOptions) *sourceengines.SourceEngines {
	engineName, url, apiKey := this.getOptions(options...)

	return sourceengines.New(engineName, apiKey, url)
}

// Schema creates a new Schema instance for getting and editing a schema of an engine
// API used https://www.elastic.co/guide/en/app-search/7.15/schema.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) Schema(options ...*OverideOptions) *schema.Schema {
	engineName, url, apiKey := this.getOptions(options...)

	return schema.New(engineName, apiKey, url)
}

// SearchSettings creates a new SearchSettings instance for getting and editing search settings, result settings, precision and boost
// API used https://www.elastic.co/guide/en/app-search/7.15/search-settings.html
// Options is an object that can be passed to overide the default configuration of api key and url.
// Only the 1st Option object is used
func (this *AppSearch) SearchSettings(options ...*OverideOptions) *searchsettings.SearchSettings {
	engineName, url, apiKey := this.getOptions(options...)

	return searchsettings.New(engineName, apiKey, url)
}

func (this *AppSearch) getOptions(options ...*OverideOptions) (engineName, url, apiKey string) {
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

	return engineName, url, apiKey
}
