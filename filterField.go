package criteria

type FilterField interface {
	GetValue() string
}
type filterField struct {
	value string
}

func (f *filterField) GetValue() string {
	return f.value
}

func NewFilterField(value string) FilterField {
	return &filterField{value: value}
}
