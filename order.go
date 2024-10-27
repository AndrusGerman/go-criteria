package criteria

type Order interface {
	GetOrderBy() OrderBy
	GetOrderType() OrderType
	IsNone() bool
}

type order struct {
	orderBy   OrderBy
	orderType OrderType
}

// GetOrderBy implements Order.
func (o *order) GetOrderBy() OrderBy {
	return o.orderBy
}

// GetOrderType implements Order.
func (o *order) GetOrderType() OrderType {
	return o.orderType
}

// IsNone implements Order.
func (o *order) IsNone() bool {
	return o.orderType.isNone()
}

func NewOrderNone() Order {
	return NewOrder(NewOrderBy(""), NONE)
}

func NewOrder(orderBy OrderBy, orderType OrderType) Order {
	return &order{
		orderBy:   orderBy,
		orderType: orderType,
	}
}

func NewOrderFromPrimitives(orderBy string, orderType string) Order {
	return NewOrder(NewOrderBy(orderBy), OrderType(orderType))
}
