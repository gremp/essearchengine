package appsearch

import (
	"fmt"
	"github.com/gremp/essearchengine/generators/boostgenerators"
	"github.com/gremp/essearchengine/generators/filtergenerators"
	"github.com/gremp/essearchengine/generators/resultfieldgenerators"
	"github.com/gremp/essearchengine/generators/searchffieldgenerators"
	"github.com/gremp/essearchengine/search"
	"testing"
)

func TestGetInstance(t *testing.T) {
	engine := Init(
		"https://vsale-dev-elastic.ent.eu-west-3.aws.elastic-cloud.com",
		"products-do-not-touch-products-meta-engine",
		"private-dg2m9bnw3zk11x9sdukpwbik",
	)

	searchStruct := engine.Search(&OverideOptions{})
	filterAll := make([]filtergenerators.IFilter, 0)

	filter := filtergenerators.CreateValueFilter("shopID", 2)
	filterAll = append(filterAll, filter)
	filters := filtergenerators.CreateFilters(filterAll, nil, nil)
	options := searchStruct.
		SearchField("name", searchffieldgenerators.CreateSearchField(0)).
		SearchField("surname", searchffieldgenerators.CreateSearchField(10)).
		ResultField("name", resultfieldgenerators.CreateResultField(true, 0, false, 0, false)).
		ResultField("surname", resultfieldgenerators.CreateResultField(true, 20, true, 60, true)).
		Filters(filters).
		Analytics("tag1", "tag2").
		Boost("name", boostgenerators.CreateValueBoost("test", boostgenerators.BoostOperations.Multiply, 1.0)).
		Boost("value", boostgenerators.CreateFunctionalBoost(boostgenerators.BoostFunctions.Exponential, boostgenerators.BoostOperations.Multiply, 1.0)).
		Page(1, 20).
		Sort("name", search.DirectionAsc).
		Group(&search.GroupOptions{
			Field:    "tets",
			Size:     2,
			Sort:     map[string]search.SortDirection{"name": search.DirectionAsc},
			Collapse: false,
		}).
		Query("iphone").
		GetRequestOptions()

	fmt.Println(options)

}
