package main

import (
	"encoding/json"
	"os"

	"github.com/AndrusGerman/go-criteria"
	criteriatomongodb "github.com/AndrusGerman/go-criteria/driver/criteria-to-mongodb"
)

func main() {

	var crit, err = criteria.NewCriteriaBuilder().
		Order(
			criteria.NewOrder(
				criteria.NewOrderBy("_id"),
				criteria.ASC,
			),
		).
		Filters(
			criteria.NewFilterType(
				criteria.NewFilterField("title"),
				criteria.EQUAL,
				criteria.NewFilterValue("booking title"),
			),
			criteria.NewFilterType(
				criteria.NewFilterField("status"),
				criteria.NOT_EQUAL,
				criteria.NewFilterValue("Canceled"),
			),
			criteria.NewFilterType(
				criteria.NewFilterField("capacity"),
				criteria.LOWER_THAN,
				criteria.NewFilterValue(2),
			),
		).GetCriteria()

	if err != nil {
		panic(err)
	}

	var query = criteriatomongodb.NewCriteriaToMongodb().Convert(
		[]string{},
		crit,
		nil,
	)

	json.NewEncoder(os.Stdout).Encode(query)
}
