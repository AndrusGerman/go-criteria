package criteria

type FilterValue interface {
	GetValue() any
}
type filterValue struct {
	value any
}

func (f *filterValue) GetValue() any {
	return f.value
}

func NewFilterValue(value any) FilterValue {
	return &filterValue{value: value}
}
