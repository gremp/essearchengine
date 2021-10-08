package searchsettings

type SearchSettingsConfig struct {
	SearchFields SearchFields                   `json:"search_fields"`
	ResultFields ResultsFields                  `json:"result_fields"`
	Precision    int                            `json:"precision"`
	Boosts       map[string]SearchSettingsBoost `json:"boosts"`
}

type SearchSettingsBoost struct {
	Type   string   `json:"type"`
	Factor float64  `json:"factor"`
	Value  []string `json:"value"`
}

type SingleFieldSettings struct {
	Weight int `json:"weight"`
}

type SingleResultSettings struct {
	Raw     *SingleResultSettingsRaw     `json:"raw,omitempty"`
	Snippet *SingleResultSettingsSnippet `json:"snippet,omitempty"`
}

type SingleResultSettingsRaw struct {
	Size int `json:"size,omitempty"`
}
type SingleResultSettingsSnippet struct {
	Size     int  `json:"size"`
	Fallback bool `json:"fallback,omitempty"`
}

type SearchFields map[string]*SingleFieldSettings
type ResultsFields map[string]*SingleResultSettings
