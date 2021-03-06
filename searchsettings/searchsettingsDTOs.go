package searchsettings

import (
	"github.com/gremp/essearchengine/generators/resultfieldgenerators"
	"github.com/gremp/essearchengine/generators/searchffieldgenerators"
)

type SearchSettingsConfig struct {
	SearchFields searchffieldgenerators.SearchFields `json:"search_fields"`
	ResultFields resultfieldgenerators.ResultsFields `json:"result_fields"`
	Precision    int                                 `json:"precision"`
	Boosts       map[string]SearchSettingsBoost      `json:"boosts"`
}

type SearchSettingsBoost struct {
	Type   string   `json:"type"`
	Factor float64  `json:"factor"`
	Value  []string `json:"value"`
}
