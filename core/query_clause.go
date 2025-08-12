package core

import (
	"fmt"
	"strings"
)

// QueryClause is the interface for all query clause types
type QueryClause interface {
	String() string
	ToJsonObj() map[string]any
}

type Filter struct {
	K         string   `json:"k"`         
	V         []string `json:"v"`         
	Op        string   `json:"op"`       
	Extracted bool     `json:"extracted"` 
	Computed  bool     `json:"computed"`  
	DataType  string   `json:"dataType"` 
}

func (f *Filter) String() string {
	if len(f.V) == 0 {
		return ""
	}

	switch f.Op {
	case OP_EQ:
		return fmt.Sprintf("%s = %s", f.K, f.V[0])
	case OP_GT:
		return fmt.Sprintf("%s > %s", f.K, f.V[0])
	case OP_GE:
		return fmt.Sprintf("%s >= %s", f.K, f.V[0])
	case OP_LT:
		return fmt.Sprintf("%s < %s", f.K, f.V[0])
	case OP_LE:
		return fmt.Sprintf("%s <= %s", f.K, f.V[0])
	case OP_REGEX:
		return fmt.Sprintf("regexMatches(%s, %s)", f.K, f.V[0])
	case OP_CONTAINS:
		return fmt.Sprintf("%s contains %s", f.K, f.V[0])
	case OP_IN:
		return fmt.Sprintf("%s in (%s)", f.K, strings.Join(f.V, ", "))
	default:
		return ""
	}
}

func (f *Filter) LabelName() string {
	return f.K
}

func (f *Filter) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["k"] = f.K
	result["v"] = f.V
	result["op"] = f.Op
	result["extracted"] = f.Extracted
	result["computed"] = f.Computed
	result["dataType"] = f.DataType
	return result
}

// Represents a logical combination of two query clauses
type BinaryClause struct {
	Q1 QueryClause `json:"q1"` 
	Q2 QueryClause `json:"q2"` 
	Op string      `json:"op"` 
}

func (bc *BinaryClause) String() string {
	if bc.Q1 == nil || bc.Q2 == nil {
		return ""
	}
	return fmt.Sprintf("(%s %s %s)", bc.Q1.String(), bc.Op, bc.Q2.String())
}

func (bc *BinaryClause) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["q1"] = bc.Q1.ToJsonObj() 
	result["q2"] = bc.Q2.ToJsonObj() 
	result["op"] = bc.Op
	return result
}

// NotClause represents the negation of a query clause
type NotClause struct {
	Not QueryClause `json:"not"`
}

func (nc *NotClause) String() string {
	if nc.Not == nil {
		return ""
	}
	return fmt.Sprintf("not(%s)", nc.Not.String())
}

func (nc *NotClause) ToJsonObj() map[string]any {
	result := make(map[string]any)
	result["not"] = nc.Not.ToJsonObj()
	return result
}
