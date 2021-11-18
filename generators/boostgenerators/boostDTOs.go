package boostgenerators

type BoostObject map[string]BoostSingleObject

type BoostSingleObject struct {
	Type      BoostType      `json:"type"`
	Value     interface{}    `json:"value"`
	Operation BoostOperation `json:"operation"`
	Factor    float64        `json:"factor"`
	Function  BoostFunction  `json:"function"`
	Center    string         `json:"center"`
}

type BoostType string
type BoostOperation string
type BoostFunction string

var (
	BoostTypes = struct {
		Value      BoostType
		Functional BoostType
		Proximity  BoostType
	}{
		Value:      "value",
		Functional: "functional",
		Proximity:  "proximity",
	}

	BoostOperations = struct {
		Add      BoostOperation
		Multiply BoostOperation
	}{
		Add:      "add",
		Multiply: "multiply",
	}

	BoostFunctions = struct {
		Linear      BoostFunction
		Exponential BoostFunction
		Logarithmic BoostFunction
		Gaussian    BoostFunction
	}{
		Linear:      "linear",
		Exponential: "exponential",
		Logarithmic: "logarithmic",
		Gaussian:    "gaussian",
	}
)
