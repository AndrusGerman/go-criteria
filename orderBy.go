package criteria

type OrderBy interface {
	GetByField() string
}
type orderBy struct {
	byValue string
}

func (o *orderBy) GetByField() string {
	return o.byValue
}

func NewOrderBy(value string) OrderBy {
	return &orderBy{byValue: value}
}
