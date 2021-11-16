package searchsettings

import "github.com/gremp/essearchengine/helpers"

type SearchSettingsConfig struct {
	SearchFields helpers.SearchFields           `json:"search_fields"`
	ResultFields helpers.ResultsFields          `json:"result_fields"`
	Precision    int                            `json:"precision"`
	Boosts       map[string]SearchSettingsBoost `json:"boosts"`
}

type SearchSettingsBoost struct {
	Type   string   `json:"type"`
	Factor float64  `json:"factor"`
	Value  []string `json:"value"`
}
