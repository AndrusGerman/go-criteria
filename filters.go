package criteria

type filters struct {
	value []Filter
}

// GetValue implements Filters.
func (f *filters) GetValue() []Filter {
	return f.value
}

// IsEmpty implements Filters.
func (f *filters) IsEmpty() bool {
	return len(f.value) == 0
}

// ToPrimitives implements Filters.
func (f *filters) ToPrimitives() []FiltersPrimitive {
	var filtersPrimitive = make([]FiltersPrimitive, len(f.value))
	for i := range f.value {
		filtersPrimitive[i] = f.value[i].ToPrimitives()
	}
	return filtersPrimitive
}

type Filters interface {
	GetValue() []Filter
	ToPrimitives() []FiltersPrimitive
	IsEmpty() bool
}

func NewFilters(value []Filter) Filters {
	return &filters{value: value}
}

func NewFiltersNone() Filters {
	return &filters{value: make([]Filter, 0)}
}

func NewFiltersFromPrimitives(filtersPrimitive []FiltersPrimitive) []Filter {
	var filters []Filter = make([]Filter, len(filtersPrimitive))
	for i := range filtersPrimitive {
		filters[i] = NewFilterFromPrimitives(
			filtersPrimitive[i].GetField(),
			filtersPrimitive[i].GetOperator(),
			filtersPrimitive[i].GetValue(),
		)
	}
	return filters
}
