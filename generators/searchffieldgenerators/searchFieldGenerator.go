package searchffieldgenerators

func CreateSearchField(weight int) *SingleSearchField {
	searchField := &SingleSearchField{}

	if weight != 0 {
		searchField.Weight = weight
	}
	return searchField
}
