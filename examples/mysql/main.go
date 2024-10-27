package main

import (
	"fmt"
	"net/url"

	criteriafromurl "github.com/AndrusGerman/go-criteria/driver/criteria-from-url"
	criteriatomysql "github.com/AndrusGerman/go-criteria/driver/criteria-to-mysql"
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

	var sql, params = criteriatomysql.NewCriteriaToMySqlConverter().Convert(
		[]string{"userId"},
		"users",
		crit,
		nil,
	)

	fmt.Println(sql, params)

}
