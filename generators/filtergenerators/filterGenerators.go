package filtergenerators

import (
	"github.com/gremp/essearchengine/helpers"
)

func CreateFilters(all []IFilter, any []IFilter, none []IFilter) *Filters {
	return &Filters{
		All:  all,
		Any:  any,
		None: none,
	}
}

func CreateValueFilter(key string, value interface{}) *ValueFilter {
	filter := make(ValueFilterObject)
	filterObj := &ValueFilter{Filter: &filter}

	filter[key] = value

	return filterObj
}

func CreateRangeFilter(key string, from interface{}, to interface{}) *RangeFilter {
	filter := make(RangeFilterObject)
	filterObj := &RangeFilter{Filter: &filter}

	filter[key] = &RangeFilterSingleObject{
		From: from,
		To:   to,
	}

	return filterObj
}

func CreateGeoFilter(key string, center helpers.GeoPoint, unit GeoUnit, distance *float64, from *int, to *int) *GeoFilter {
	filter := make(GeoFilterObject)
	filterObj := &GeoFilter{Filter: &filter}

	filter[key] = &GeoFilterSingleObject{
		Center: center.GetStr(),
		Unit:   unit,
	}

	if distance != nil {
		filter[key].Distance = *distance
	}
	if from != nil {
		filter[key].From = *from
	}
	if to != nil {
		filter[key].To = *to
	}

	return filterObj
}
