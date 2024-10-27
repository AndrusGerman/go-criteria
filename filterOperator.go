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

func (o Operator) isContains() bool {
	return o == CONTAINS
}
func (o Operator) isNotContains() bool {
	return o == NOT_CONTAINS
}
func (o Operator) isNotEquals() bool {
	return o == NOT_EQUAL
}
func (o Operator) isGreaterThan() bool {
	return o == GREATER_THAN
}
func (o Operator) isGreaterThanOrEqual() bool {
	return o == GREATER_THAN_OR_EQUAL
}
func (o Operator) isLowerThan() bool {
	return o == LOWER_THAN
}
func (o Operator) isLowerThanOrEqual() bool {
	return o == LOWER_THAN_OR_EQUAL
}
