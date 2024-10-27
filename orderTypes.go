package criteria

type OrderType string

const (
	ASC  OrderType = "ASC"
	DESC OrderType = "DESC"
	NONE OrderType = "NONE"
)

func (o OrderType) isNone() bool {
	return o == NONE
}

func (o OrderType) GetValue() string {
	return string(o)
}
