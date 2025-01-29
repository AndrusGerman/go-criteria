package criteria

type criteria struct {
	filters    Filters
	order      Order
	pageSize   *int
	pageNumber *int
}

// HasFilters implements Criteria.
func (c *criteria) HasFilters() bool {
	return !c.filters.IsEmpty()
}

// HasOrder implements Criteria.
func (c *criteria) HasOrder() bool {
	return !c.order.IsNone()
}

// GetFilters implements Criteria.
func (c *criteria) GetFilters() Filters {
	return c.filters
}

// GetOrder implements Criteria.
func (c *criteria) GetOrder() Order {
	return c.order
}

// GetPageNumber implements Criteria.
func (c *criteria) GetPageNumber() *int {
	return c.pageNumber
}

// GetPageSize implements Criteria.
func (c *criteria) GetPageSize() *int {
	return c.pageSize
}

type Criteria interface {
	GetFilters() Filters
	GetOrder() Order
	GetPageSize() *int
	GetPageNumber() *int
	HasOrder() bool
	HasFilters() bool
}

type CriteriaBuilder struct {
	filters    Filters
	order      Order
	pageSize   *int
	pageNumber *int
}

func NewCriteriaBuilder() *CriteriaBuilder {
	return &CriteriaBuilder{}
}

func (cb *CriteriaBuilder) Filters(filters Filters) *CriteriaBuilder {
	cb.filters = filters
	return cb
}

func (cb *CriteriaBuilder) FiltersPrimitive(filters []FiltersPrimitive) *CriteriaBuilder {
	return cb.Filters(NewFiltersFromPrimitives(filters))
}

func (cb *CriteriaBuilder) Order(order Order) *CriteriaBuilder {
	cb.order = order
	return cb
}

func (cb *CriteriaBuilder) OrderPrimitive(orderBy string, orderType string) *CriteriaBuilder {
	return cb.Order(NewOrderFromPrimitives(orderBy, orderType))
}

func (cb *CriteriaBuilder) PageSize(pageSize int) *CriteriaBuilder {
	cb.pageSize = &pageSize
	return cb
}

func (cb *CriteriaBuilder) PageNumber(pageNumber int) *CriteriaBuilder {
	cb.pageNumber = &pageNumber
	return cb
}
func (cb *CriteriaBuilder) Reset() *CriteriaBuilder {
	cb.pageNumber = nil
	cb.filters = nil
	cb.order = nil
	cb.pageNumber = nil
	cb.pageSize = nil
	return cb
}

func (cb *CriteriaBuilder) Copy() *CriteriaBuilder {
	return &(*cb)
}

func (cb *CriteriaBuilder) GetCriteria() (Criteria, error) {
	if cb.filters == nil {
		cb.filters = NewFiltersNone()
	}
	if cb.order == nil {
		cb.order = NewOrderNone()
	}

	if cb.pageNumber != nil && cb.pageSize == nil {
		return nil, ERR_INVALID_CRITERIA_BUILDER_PAGER
	}

	return &criteria{
		filters:    cb.filters,
		order:      cb.order,
		pageSize:   cb.pageSize,
		pageNumber: cb.pageNumber,
	}, nil
}

func (cb *CriteriaBuilder) MustGetCriteria() Criteria {
	criteria, err := cb.GetCriteria()
	if err != nil {
		panic(err)
	}
	return criteria
}

func EmptyCriteria() Criteria {
	return NewCriteriaBuilder().MustGetCriteria()
}
