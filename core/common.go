package core

import (
	"fmt"
	"strings"
)

type MetricType string

const (
	METRIC_TYPE_RATE      MetricType = "rate"
	METRIC_TYPE_COUNTER   MetricType = "count"
	METRIC_TYPE_GAUGE     MetricType = "gauge"
	METRIC_TYPE_HISTOGRAM MetricType = "histogram"
)

// Query operators
const (
	OP_EXISTS     = "exists"
	OP_NOT_EQUALS = "!="
	OP_REGEX      = "regex"
	OP_IN         = "in"
	OP_NOT_IN     = "not_in"
	OP_CONTAINS   = "contains"
	OP_EQ         = "eq"
	OP_HAS        = "has"
	OP_GT         = "gt"
	OP_GE         = "ge"
	OP_LT         = "lt"
	OP_LE         = "le"
)

// Data types
const (
	DATA_TYPE_STRING    = "string"
	DATA_TYPE_NUMBER    = "number"
	DATA_TYPE_DURATION  = "duration"
	DATA_TYPE_DATA_SIZE = "datasize"
)

// Aggregation functions
const (
	AGG_MAX   = "max"
	AGG_MIN   = "min"
	AGG_SUM   = "sum"
	AGG_COUNT = "count"
	AGG_AVG   = "avg"
)

// Datasets
const (
	DATASET_LOGS    = "logs"
	DATASET_METRICS = "metrics"
	DATASET_SPANS   = "spans"
)

// Order directions
const (
	ORDER_ASCENDING  = "ASC"
	ORDER_DESCENDING = "DESC"
)

func (m MetricType) String() string {
	return string(m)
}

func FromStr(str string) (MetricType, error) {
	switch strings.ToLower(strings.TrimSpace(str)) {
	case string(METRIC_TYPE_RATE):
		return METRIC_TYPE_RATE, nil
	case string(METRIC_TYPE_COUNTER), "counter":
		return METRIC_TYPE_COUNTER, nil
	case string(METRIC_TYPE_GAUGE):
		return METRIC_TYPE_GAUGE, nil
	case string(METRIC_TYPE_HISTOGRAM):
		return METRIC_TYPE_HISTOGRAM, nil
	default:
		return "", fmt.Errorf("unknown metric type %q", str)
	}
}
