package sourceengines

import (
	"github.com/gremp/essearchengine/helpers"
)

type Language string

var (
	LangBrazilianPortuguese Language = "pt-br"
	LangChinese             Language = "zh"
	LangDanish              Language = "da"
	LangDutch               Language = "nl"
	LangEnglish             Language = "en"
	LangFrench              Language = "fr"
	LangGerman              Language = "de"
	LangItalian             Language = "it"
	LangJapanese            Language = "ja"
	LangKorean              Language = "ko"
	LangPortuguese          Language = "pt"
	LangRussian             Language = "ru"
	LangSpanish             Language = "es"
	LangThai                Language = "th"
	LangUniversal           Language
)

type CreateEngineReq struct {
	Name     string   `json:"name"`
	Language Language `json:"language,omitempty"`
}

type GenericEnginesRes struct {
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	Language      interface{} `json:"language"`
	DocumentCount int         `json:"document_count"`
}

type ListEnginesRes struct {
	Meta    helpers.ResultMeta  `json:"meta"`
	Results []GenericEnginesRes `json:"results"`
}

type DeleteRes struct {
	Deleted bool `json:"deleted"`
}
