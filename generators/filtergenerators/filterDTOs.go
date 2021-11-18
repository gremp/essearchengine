package filtergenerators

type Filters struct {
	All  []IFilter `json:"all,omitempty"`
	Any  []IFilter `json:"any,omitempty"`
	None []IFilter `json:"none,omitempty"`
}

type IFilter interface {
	GetFilter() interface{}
	ToMapString() map[string]interface{}
}

type ValueFilter struct {
	Filter *ValueFilterObject
}

type GeoFilter struct {
	Filter *GeoFilterObject
}

type RangeFilter struct {
	Filter *RangeFilterObject
}

func (this *ValueFilter) GetFilter() interface{} {
	return this.Filter
}

func (this *ValueFilter) ToMapString() map[string]interface{} {
	ret := make(map[string]interface{})
	for key, value := range *this.Filter {
		ret[key] = value
	}
	return ret
}

func (this *Filters) GetFilter() interface{} {
	filtersAll := getFiltersSanitized(this.All)
	filtersAny := getFiltersSanitized(this.Any)
	filtersNone := getFiltersSanitized(this.None)
	return map[string]interface{}{
		"all":  filtersAll,
		"any":  filtersAny,
		"none": filtersNone,
	}
}

func getFiltersSanitized(incFilters []IFilter) []map[string]interface{} {
	filters := make([]map[string]interface{}, 0)
	for _, filter := range incFilters {
		filterMap := make(map[string]interface{})
		for key, value := range filter.ToMapString() {
			filterMap[key] = value
		}
		filters = append(filters, filterMap)
	}
	return filters

}

func (this *GeoFilter) GetFilter() interface{} {
	return this.Filter
}

func (this *GeoFilter) ToMapString() map[string]interface{} {
	ret := make(map[string]interface{})
	for key, value := range *this.Filter {
		ret[key] = value
	}
	return ret
}

func (this *RangeFilter) GetFilter() interface{} {
	return this.Filter
}

func (this *RangeFilter) ToMapString() map[string]interface{} {
	ret := make(map[string]interface{})
	for key, value := range *this.Filter {
		ret[key] = value
	}
	return ret
}

type ValueFilterObject map[string]interface{}
type RangeFilterObject map[string]*RangeFilterSingleObject

type RangeFilterSingleObject struct {
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

type GeoFilterObject map[string]*GeoFilterSingleObject

type GeoFilterSingleObject struct {
	Center   string  `json:"center"`
	Unit     GeoUnit `json:"unit"`
	Distance float64 `json:"distance,omitempty"`
	From     int     `json:"from,omitempty"`
	To       int     `json:"to,omitempty"`
}

type GeoUnit string

var GeoUnits = struct {
	Milimeter  GeoUnit
	Centimeter GeoUnit
	Meter      GeoUnit
	Kilometer  GeoUnit
	Inch       GeoUnit
	Feet       GeoUnit
	Yard       GeoUnit
	Mile       GeoUnit
}{
	Milimeter:  "mm",
	Centimeter: "cm",
	Meter:      "m",
	Kilometer:  "km",
	Inch:       "in",
	Feet:       "ft",
	Yard:       "yd",
	Mile:       "mi",
}
