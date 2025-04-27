package criteria

type filtersPrimitive struct {
	Field    string
	Operator string
	Value    any
}

type FiltersPrimitive interface {
	GetField() string
	GetOperator() string
	GetValue() any
}

type filter struct {
	field    FilterField
	operator Operator
	value    FilterValue
}

func (f *filtersPrimitive) GetField() string {
	return f.Field
}
func (f *filtersPrimitive) GetValue() any {
	return f.Value
}
func (f *filtersPrimitive) GetOperator() string {
	return f.Operator
}

func (f *filter) GetField() FilterField {
	return f.field
}

func (f *filter) GetOperator() Operator {
	return f.operator
}

func (f *filter) GetValue() FilterValue {
	return f.value
}

type Filter interface {
	ToPrimitives() FiltersPrimitive
	GetField() FilterField
	GetOperator() Operator
	GetValue() FilterValue
}

func (f *filter) ToPrimitives() FiltersPrimitive {
	return &filtersPrimitive{
		Field:    f.field.String(),
		Operator: string(f.operator),
		Value:    f.value.GetValue(),
	}
}

func NewFilter(field string, operator Operator, value any) Filter {
	return &filter{
		field:    NewFilterField(field),
		operator: operator,
		value:    NewFilterValue(value),
	}
}

func NewFilterType(field FilterField, operator Operator, value FilterValue) Filter {
	return &filter{
		field:    field,
		operator: operator,
		value:    value,
	}
}

func NewFilterFromPrimitives(field string, operator string, value any) Filter {
	return &filter{
		field:    NewFilterField(field),
		operator: Operator(operator),
		value:    NewFilterValue(value),
	}
}

func NewFilterPrimitive(field string, operator string, value string) FiltersPrimitive {
	return &filtersPrimitive{
		Field:    field,
		Operator: operator,
		Value:    value,
	}
}
