package criteria

import "errors"

var (
	ERR_INVALID_CRITERIA_BUILDER_PAGER       = errors.New("Page size is required when page number is defined")
	ERR_CRITERIA_BUILDER_FILTERS_NOT_DEFINED = errors.New("Filters is required")
	ERR_CRITERIA_BUILDER_ORDER_NOT_DEFINED   = errors.New("Order is required")
)
