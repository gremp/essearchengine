package documents

type UpsertResponseItem struct {
	ID string `json:"id"`
	Errors []string `json:"errors"`
}

type DeleteResponseItem struct {
	ID string `json:"id"`
	Deleted bool `json:"deleted"`
}

type UpsertDocumentErrors struct {
	Total int
	Errors []UpsertResponseItem
}

