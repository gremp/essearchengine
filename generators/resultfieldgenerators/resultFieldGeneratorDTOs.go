package resultfieldgenerators

type ResultsFields map[string]*SingleResultField
type SingleResultField struct {
	Raw     *SingleResultFieldRaw     `json:"raw,omitempty"`
	Snippet *SingleResultFieldSnippet `json:"snippet,omitempty"`
}

type SingleResultFieldRaw struct {
	Size int `json:"size,omitempty"`
}
type SingleResultFieldSnippet struct {
	Size     int  `json:"size"`
	Fallback bool `json:"fallback,omitempty"`
}
