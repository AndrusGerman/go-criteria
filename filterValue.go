package criteria

type FilterValue interface {
	GetValue() string
}
type filterValue struct {
	value string
}

func (f *filterValue) GetValue() string {
	return f.value
}

func NewFilterValue(value string) FilterValue {
	return &filterValue{value: value}
}
