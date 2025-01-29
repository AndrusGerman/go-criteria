package criteria

type FilterField string

func (f FilterField) String() string {
	return string(f)
}

func NewFilterField(value string) FilterField {
	return FilterField(value)
}
