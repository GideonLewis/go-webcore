package common

type Condition struct {
	Field string      `json:"field,omitempty" query:"field"`
	Value interface{} `json:"value,omitempty" query:"value"`
}
