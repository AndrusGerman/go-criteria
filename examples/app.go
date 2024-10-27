package main

import (
	"fmt"

	"github.com/AndrusGerman/go-criteria"
	criteriatomysql "github.com/AndrusGerman/go-criteria/driver/criteria-to-mysql"
)

func main() {
	var builderCriteria = criteria.NewCriteriaBuilder()

	var criteria, err = builderCriteria.
		Filters(
			criteria.NewFilters(
				[]criteria.Filter{
					criteria.NewFilter(criteria.NewFilterField("companyName"), criteria.CONTAINS, criteria.NewFilterValue("ppl")),
				},
			)).
		Order(
			criteria.NewOrder(criteria.NewOrderBy("createdAt"), criteria.ASC),
		).
		GetCriteria()
	if err != nil {
		panic(err)
	}

	var sql, params = criteriatomysql.NewCriteriaToMySqlConverter().Convert(
		[]string{"userId"},
		"companies",
		criteria,
		nil,
	)

	fmt.Println(sql, params)
}
