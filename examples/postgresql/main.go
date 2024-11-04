package main

import (
	"fmt"
	"net/url"

	criteriafromurl "github.com/AndrusGerman/go-criteria/driver/criteria-from-url"
	criteriatopostgresql "github.com/AndrusGerman/go-criteria/driver/criteria-to-postgresql"
)

func main() {

	var urlParse, err = url.Parse("http://localhost:3000/api/users?filters[0][field]=name&filters[0][operator]=CONTAINS&filters[0][value]=Javi")
	if err != nil {
		panic(err)
	}

	crit, err := criteriafromurl.NewCriteriaFromUrlConverter().ToCriteria(urlParse)
	if err != nil {
		panic(err)
	}

	var sql, params = criteriatopostgresql.NewCriteriaToPostgreSQLConverter().Convert(
		[]string{"userId"},
		"users",
		crit,
		nil,
	)

	fmt.Println(sql, params)

}
