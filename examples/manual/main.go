package main

import (
	"fmt"

	"github.com/AndrusGerman/go-criteria"
)

func main() {

	var crit, err = criteria.NewCriteriaBuilder().
		Order(criteria.NewOrderNone()).
		Filters(
			criteria.NewFilters(
				[]criteria.Filter{
					criteria.NewFilter(
						criteria.NewFilterField("userId"),
						criteria.EQUAL,
						criteria.NewFilterValue("10"),
					),
					criteria.NewFilter(
						criteria.NewFilterField("companyId"),
						criteria.GREATER_THAN,
						criteria.NewFilterValue("12"),
					),
					criteria.NewFilter(
						criteria.NewFilterField("companyName"),
						criteria.CONTAINS,
						criteria.NewFilterValue("app"),
					),
				},
			),
		).GetCriteria()

	if err != nil {
		panic(err)
	}

	fmt.Println("Criteria: ", crit)
}
