package schema

type SchemaType string

var (
	TypeText        SchemaType = "text"
	TypeNumber      SchemaType = "number"
	TypeGeolocation SchemaType = "geolocation"
	TypeDate        SchemaType = "date"
)

type SchemaDefinition map[string]SchemaType
