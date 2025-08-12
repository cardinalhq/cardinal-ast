package core

import (
	"fmt"
	"strings"
)

// Represents a field extracted from a regex pattern
type ExtractedField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewExtractedField(name, fieldType string) *ExtractedField {
	return &ExtractedField{
		Name: name,
		Type: fieldType,
	}
}

func (ef *ExtractedField) String() string {
	return fmt.Sprintf("%s:%s", ef.Name, ef.Type)
}

type Extractor struct {
	Regex      string           `json:"regex"`      // Regular expression pattern
	Fields     []ExtractedField `json:"fields"`     // Fields to extract
	InputField string           `json:"inputField"` // Input field to extract from (defaults to "message")
}

func NewExtractor(regex string, fields []ExtractedField, inputField string) *Extractor {
	if inputField == "" {
		inputField = "message" // Default value
	}
	return &Extractor{
		Regex:      regex,
		Fields:     fields,
		InputField: inputField,
	}
}

func (e *Extractor) String() string {
	fieldNames := make([]string, len(e.Fields))
	for i, field := range e.Fields {
		fieldNames[i] = field.Name
	}
	return fmt.Sprintf("extract(%s) from %s using regex: %s", strings.Join(fieldNames, ", "), e.InputField, e.Regex)
}

func (e *Extractor) GetFieldNames() []string {
	names := make([]string, len(e.Fields))
	for i, field := range e.Fields {
		names[i] = field.Name
	}
	return names
}

func (e *Extractor) GetFieldTypeMap() map[string]string {
	typeMap := make(map[string]string)
	for _, field := range e.Fields {
		typeMap[field.Name] = field.Type
	}
	return typeMap
}
