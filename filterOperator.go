package criteria

type Operator string

const (
	EQUAL                 Operator = "="
	NOT_EQUAL             Operator = "!="
	GREATER_THAN          Operator = ">"
	GREATER_THAN_OR_EQUAL Operator = ">="
	LOWER_THAN            Operator = "<"
	LOWER_THAN_OR_EQUAL   Operator = "<="
	CONTAINS              Operator = "CONTAINS"
	NOT_CONTAINS          Operator = "NOT_CONTAINS"
)

func (o Operator) IsContains() bool {
	return o == CONTAINS
}
func (o Operator) IsNotContains() bool {
	return o == NOT_CONTAINS
}
func (o Operator) IsNotEquals() bool {
	return o == NOT_EQUAL
}
func (o Operator) IsGreaterThan() bool {
	return o == GREATER_THAN
}
func (o Operator) IsGreaterThanOrEqual() bool {
	return o == GREATER_THAN_OR_EQUAL
}
func (o Operator) IsLowerThan() bool {
	return o == LOWER_THAN
}
func (o Operator) IsLowerThanOrEqual() bool {
	return o == LOWER_THAN_OR_EQUAL
}
