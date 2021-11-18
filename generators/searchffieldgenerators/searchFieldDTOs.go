package searchffieldgenerators

type SearchFields map[string]*SingleSearchField
type SingleSearchField struct {
	Weight int `json:"weight"`
}
