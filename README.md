<h1 align="center">
  🎼 Criteria for Go
</h1>

<!-- <p align="center">
    <a href="https://github.com/CodelyTV"><img src="https://img.shields.io/badge/Codely-OS-green.svg?style=flat-square" alt="Codely Open Source projects"/></a>
    <a href="https://pro.codely.com"><img src="https://img.shields.io/badge/Codely-Pro-black.svg?style=flat-square" alt="Codely Pro courses"/></a>
</p> -->

## 📥 Installation

To install the base criteria dependency, run the following command:

```sh
go get github.com/AndrusGerman/go-criteria
```


## 🎉 Examples

* MYSQL from URL
```go
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
```

out: `SELECT userId FROM users WHERE name LIKE ?` params: `[%Javi%]`


* CUSTOM criteria
```go
var crit, err = criteria.NewCriteriaBuilder().
	Order(criteria.NewOrderNone()).
	Filters(
		criteria.NewFilter(
			"userId",
			criteria.EQUAL,
			"10",
		),
		criteria.NewFilter(
			"companyId",
			criteria.GREATER_THAN,
			"12",
		),
		criteria.NewFilter(
			"companyName",
			criteria.CONTAINS,
			"app",
		),
	).GetCriteria()
```

* criteria to MongoDB
```go
var query = criteriatomongodb.NewCriteriaToMongodb().Convert(
	[]string{},
	crit,
	nil,
)
```

## 🐂 Thanks

This package is initially inspired by the implementation of typescript-criteria created by codelyTv
