package main

import (
	"fmt"

	"github.com/AndrusGerman/go-criteria"
)

func main() {
	var builderCriteria = criteria.NewCriteriaBuilder()

	var criteria, err = builderCriteria.
		Filters(
			criteria.NewFilters(
				[]criteria.Filter{
					criteria.NewFilter(criteria.NewFilterField("X"), criteria.CONTAINS, criteria.NewFilterValue("Y")),
				},
			)).
		Order(
			criteria.NewOrder(criteria.NewOrderBy("X"), criteria.ASC),
		).
		GetCriteria()
	if err != nil {
		panic(err)
	}

	fmt.Println(criteria.GetOrder().GetOrderBy())
}
