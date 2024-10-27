package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	criteriafromurl "github.com/AndrusGerman/go-criteria/driver/criteria-from-url"
	criteriatomongodb "github.com/AndrusGerman/go-criteria/driver/criteria-to-mongodb"
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

	var query = criteriatomongodb.NewCriteriaToMongodb().Convert(
		[]string{"userId"},
		crit,
		nil,
	)

	json.NewEncoder(os.Stdout).Encode(query)
}

func mainX() {

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
