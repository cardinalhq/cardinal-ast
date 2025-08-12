package core

type Compute struct {
	LabelName    string        `json:"labelName"`    // Name of the computed field
	FunctionCall *FunctionCall `json:"functionCall"` // Function call to execute
}

func NewCompute(labelName string, functionCall *FunctionCall) *Compute {
	return &Compute{
		LabelName:    labelName,
		FunctionCall: functionCall,
	}
}

func (c *Compute) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["labelName"] = c.LabelName
	result["functionCall"] = c.FunctionCall.ToJsonObj()
	return result
}

type Label struct {
	Name     string `json:"name"`
	DataType string `json:"dataType"`
}

func NewLabel(name, dataType string) *Label {
	return &Label{
		Name:     name,
		DataType: dataType,
	}
}

type Literal struct {
	Value    any    `json:"value"`    
	DataType string `json:"dataType"` 
}

func NewLiteral(value any, dataType string) *Literal {
	return &Literal{
		Value:    value,
		DataType: dataType,
	}
}

type FunctionCall struct {
	Name      string `json:"name"`     
	Arguments []any  `json:"arguments"` 
}

func NewFunctionCall(name string, arguments []any) *FunctionCall {
	return &FunctionCall{
		Name:      name,
		Arguments: arguments,
	}
}

func (fc *FunctionCall) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["name"] = fc.Name

	arguments := make([]map[string]any, len(fc.Arguments))
	for i, arg := range fc.Arguments {
		argMap := make(map[string]any)

		switch a := arg.(type) {
		case *Label:
			argMap["name"] = a.Name
			argMap["type"] = "label"
			argMap["dataType"] = a.DataType
		case *Literal:
			argMap["type"] = "literal"
			argMap["value"] = a.Value
			argMap["dataType"] = a.DataType
		case *FunctionCall:
			argMap["type"] = "functionCall"
			nestedJSON := a.ToJsonObj()
			argMap["name"] = nestedJSON["name"]
			argMap["arguments"] = nestedJSON["arguments"]
		}

		arguments[i] = argMap
	}
	result["arguments"] = arguments

	return result
}
