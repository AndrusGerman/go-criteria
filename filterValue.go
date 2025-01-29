package criteria

type FilterValue interface {
	GetValue() any
	IsEmpty() bool
}

type filterValue struct {
	value any
}

func (f *filterValue) GetValue() any {
	return f.value
}

func (f *filterValue) IsEmpty() bool {
	return f.value == nil
}

func NewFilterValue(value any) FilterValue {
	return &filterValue{value: value}
}
