package boostgenerators

import (
	"github.com/gremp/essearchengine/helpers"
)

func CreateValueBoost(value interface{}, operation BoostOperation, factor float64) BoostSingleObject {
	return BoostSingleObject{
		Type:      BoostTypes.Value,
		Value:     value,
		Operation: operation,
		Factor:    factor,
	}
}

func CreateFunctionalBoost(function BoostFunction, operation BoostOperation, factor float64) BoostSingleObject {
	return BoostSingleObject{
		Type:      BoostTypes.Functional,
		Operation: operation,
		Factor:    factor,
		Function:  function,
	}
}

func CreateProximityBoost(function BoostFunction, center helpers.GeoPoint, factor float64) BoostSingleObject {
	return BoostSingleObject{
		Type:     BoostTypes.Proximity,
		Center:   center.GetStr(),
		Factor:   factor,
		Function: function,
	}
}
