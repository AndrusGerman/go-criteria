package criteria

type OrderBy interface {
	GetValue() string
}
type orderBy struct {
	value string
}

func (o *orderBy) GetValue() string {
	return o.value
}

func NewOrderBy(value string) OrderBy {
	return &orderBy{value: value}
}
