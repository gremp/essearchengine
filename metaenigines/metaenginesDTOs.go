package metaenigines

type CreateEngineReq struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	SourceEngines []string `json:"source_engines"`
}

type GenericSourceEnginesRes struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	SourceEngines []string `json:"source_engines"`
	DocumentCount int      `json:"document_count,omitempty"`
}
