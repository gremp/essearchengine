package multisearch

import "github.com/gremp/essearchengine/search"

type MultiSearchPayload struct {
	Queries []*search.RequestOptions `json:"queries"`
}
