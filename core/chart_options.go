package core

type ChartType string

const (
	CHART_TYPE_COUNT ChartType = "count"
	CHART_TYPE_RATE  ChartType = "rate"
)

type ChartOptions struct {
	Aggregation       string    `json:"aggregation"`
	GroupBys          []string  `json:"groupBys"`          
	Type              ChartType `json:"type"`           
	RollupAggregation *string   `json:"rollup,omitempty"`   
	FieldName         *string   `json:"fieldName,omitempty"`
	FieldType         *string   `json:"fieldType,omitempty"`
}

func (co *ChartOptions) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["aggregation"] = co.Aggregation
	result["groupBys"] = co.GroupBys
	result["type"] = co.Type

	if co.RollupAggregation != nil {
		result["rollup"] = co.RollupAggregation
	}
	if co.FieldName != nil {
		result["fieldName"] = co.FieldName
	}
	if co.FieldType != nil {
		result["fieldType"] = co.FieldType
	}

	return result
}
