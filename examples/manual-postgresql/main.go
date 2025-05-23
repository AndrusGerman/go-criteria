package main

import (
	"fmt"

	"github.com/AndrusGerman/go-criteria"
	criteriatopostgresql "github.com/AndrusGerman/go-criteria/driver/criteria-to-postgresql"
)

func main() {

	var crit, err = criteria.NewCriteriaBuilder().
		Order(criteria.NewOrder(criteria.NewOrderBy("id"), criteria.ASC)).
		Filters(
			criteria.NewFilterType(
				criteria.NewFilterField("userId"),
				criteria.EQUAL,
				criteria.NewFilterValue("10"),
			),
			criteria.NewFilterType(
				criteria.NewFilterField("companyId"),
				criteria.GREATER_THAN,
				criteria.NewFilterValue("12"),
			),
			criteria.NewFilterType(
				criteria.NewFilterField("companyName"),
				criteria.CONTAINS,
				criteria.NewFilterValue("app"),
			),
		).GetCriteria()

	if err != nil {
		panic(err)
	}

	var sql, params = criteriatopostgresql.NewCriteriaToPostgreSQLConverter().Convert(
		[]string{"*"},
		"users",
		crit,
		nil,
	)

	fmt.Println(sql)
	fmt.Println(params)
}
