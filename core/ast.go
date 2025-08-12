package core

// AST is the base interface for all AST nodes
type AST interface {
	ToJsonObj() map[string]any
}

// BaseExpr represents a complete query unit
type BaseExpr struct {
	ID            string        `json:"id"`
	Dataset       string        `json:"dataset"`
	Filter        QueryClause   `json:"filter"`
	Extract       *Extractor    `json:"extract,omitempty"`
	Compute       *Compute      `json:"compute,omitempty"`
	Chart         *ChartOptions `json:"chart,omitempty"`
	Limit         *int          `json:"limit,omitempty"`
	Order         *string       `json:"order,omitempty"`
	MetricType    MetricType    `json:"metricType,omitempty"`
	ReturnResults bool          `json:"returnResults,omitempty"`
}

// Represents a mathematical operation between two AST nodes
type Formula struct {
	E1 AST    `json:"e1"`
	E2 AST    `json:"e2"`
	Op string `json:"op"`
}

type ConstantExpr struct {
	Value float64 `json:"constant"`
}

// Struct for the complete query 
type ASTInput struct {
	BaseExpressions map[string]*BaseExpr `json:"baseExpressions"`
	Formulae        []string             `json:"formulae"`
}

// JSON representation of the base expression
func (be *BaseExpr) ToJsonObj() map[string]any {
	result := make(map[string]any)

	result["id"] = be.ID
	result["dataset"] = be.Dataset
	result["limit"] = be.Limit
	result["order"] = be.Order
	result["metricType"] = be.MetricType
	result["returnResults"] = be.ReturnResults

	result["filter"] = be.Filter.ToJsonObj()

	if be.Extract != nil {
		result["extract"] = be.Extract
	}
	if be.Compute != nil {
		result["compute"] = be.Compute.ToJsonObj()
	}
	if be.Chart != nil {
		result["chart"] = be.Chart.ToJsonObj()
	}

	return result
}

// JSON representation of the constant expression
func (ce *ConstantExpr) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["constant"] = ce.Value
	return result
}

// JSON representation of the formula
func (f *Formula) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["e1"] = f.E1.ToJsonObj() // Recursive call on e1
	result["e2"] = f.E2.ToJsonObj() // Recursive call on e2
	result["op"] = f.Op
	return result
}

// JSON representation of the AST input
func (ai *ASTInput) ToJsonObj() map[string]any {
	result := make(map[string]any)

	baseExprs := make(map[string]any)
	for key, baseExpr := range ai.BaseExpressions {
		baseExprs[key] = baseExpr.ToJsonObj()
	}
	result["baseExpressions"] = baseExprs

	result["formulae"] = ai.Formulae

	return result
}
