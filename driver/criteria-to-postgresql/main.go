package criteriatopostgresql

import (
	"fmt"
	"strings"

	"github.com/AndrusGerman/go-criteria"
)

type CriteriaToPostgreSQLConverter struct {
}

func NewCriteriaToPostgreSQLConverter() *CriteriaToPostgreSQLConverter {
	return &CriteriaToPostgreSQLConverter{}
}

func (ctmsc *CriteriaToPostgreSQLConverter) Convert(
	fieldsToSelect []string,
	tableName string,
	criteria criteria.Criteria,
	mappings map[string]string,
) (string, []any) {
	var query = fmt.Sprintf("SELECT %s FROM %s", strings.Join(fieldsToSelect, ", "), tableName)
	var params []any
	if mappings == nil {
		mappings = make(map[string]string)
	}

	var queryIndex = 1
	if criteria.HasFilters() {
		query += " WHERE "
		var whereQuery []string
		for _, value := range criteria.GetFilters().GetValue() {
			var queryPart, param = ctmsc.generateWhereQuery(value, mappings, queryIndex)
			whereQuery = append(whereQuery, queryPart)
			params = append(params, param)
			queryIndex++
		}
		query += strings.Join(whereQuery, " AND ")
	}

	if criteria.HasOrder() {
		query += fmt.Sprintf(" ORDER BY %s %s", criteria.GetOrder().GetOrderBy().GetByField(), criteria.GetOrder().GetOrderType().GetValue())
	}

	var pageSize = criteria.GetPageSize()
	if pageSize != nil {
		query += fmt.Sprintf(" LIMIT $%d", queryIndex)
		queryIndex++
		params = append(params, pageSize)
	}

	var pageNumber = criteria.GetPageNumber()
	if pageSize != nil && pageNumber != nil {
		query += fmt.Sprintf(" OFFSET $%d", queryIndex)
		queryIndex++
		params = append(params, (*pageSize)*((*pageNumber)-1))
	}

	return query, params
}

func (ctmsc *CriteriaToPostgreSQLConverter) generateWhereQuery(
	filter criteria.Filter,
	mappings map[string]string,
	queryIndex int,
) (queryPart string, param any) {

	var field, ok = mappings[filter.GetField().String()]
	if !ok {
		field = filter.GetField().String()
	}

	queryPart = fmt.Sprintf("%s ", field)
	var value = filter.GetValue().GetValue()

	var queryName = fmt.Sprintf("$%d", queryIndex)

	if filter.GetOperator().IsContains() {
		queryPart += "LIKE " + queryName
		param = "%" + fmt.Sprint(value) + "%"
	} else if filter.GetOperator().IsNotContains() {
		queryPart += "NOT LIKE " + queryName
		param = "%" + fmt.Sprint(value) + "%"
	} else if filter.GetOperator().IsNotEquals() {
		queryPart += "!= " + queryName
		param = value
	} else if filter.GetOperator().IsGreaterThan() {
		queryPart += "> " + queryName
		param = value
	} else if filter.GetOperator().IsGreaterThanOrEqual() {
		queryPart += ">= " + queryName
		param = value

	} else if filter.GetOperator().IsLowerThan() {
		queryPart += "< " + queryName
		param = value

	} else if filter.GetOperator().IsLowerThanOrEqual() {
		queryPart += "<= " + queryName
		param = value

	} else {
		queryPart += fmt.Sprintf("%s %s", filter.GetOperator(), queryName)
		param = value

	}
	return queryPart, param
}
