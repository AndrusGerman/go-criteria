package criteriatomysql

import (
	"fmt"
	"strings"

	"github.com/AndrusGerman/go-criteria"
)

type CriteriaToMySqlConverter struct {
}

func NewCriteriaToMySqlConverter() *CriteriaToMySqlConverter {
	return &CriteriaToMySqlConverter{}
}

func (ctmsc *CriteriaToMySqlConverter) Convert(
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

	if criteria.HasFilters() {
		query += " WHERE "
		var whereQuery []string
		for _, value := range criteria.GetFilters().GetValue() {
			var queryPart, param = ctmsc.generateWhereQuery(value, mappings)
			whereQuery = append(whereQuery, queryPart)
			params = append(params, param)
		}
		query += strings.Join(whereQuery, " AND ")
	}
	if criteria.HasOrder() {
		query += " ORDER BY ? ?"
		params = append(params, criteria.GetOrder().GetOrderBy().GetValue(), criteria.GetOrder().GetOrderType().GetValue())
	}

	var pageSize = criteria.GetPageSize()
	if pageSize != nil {
		query += " LIMIT ?"
		params = append(params, pageSize)
	}

	var pageNumber = criteria.GetPageNumber()
	if pageSize != nil && pageNumber != nil {
		query += " OFFSET ?"
		params = append(params, (*pageSize)*((*pageNumber)-1))
	}

	return query, params
}

func (ctmsc *CriteriaToMySqlConverter) generateWhereQuery(
	filter criteria.Filter,
	mappings map[string]string,
) (queryPart string, param string) {

	var field, ok = mappings[filter.GetField().GetValue()]
	if !ok {
		field = filter.GetField().GetValue()
	}

	queryPart = fmt.Sprintf("%s ", field)
	var value = filter.GetValue().GetValue()

	if filter.GetOperator().IsContains() {
		queryPart += "LIKE ?"
		param = "%" + value + "%"
	} else if filter.GetOperator().IsNotContains() {
		queryPart += "NOT LIKE ?"
		param = "%" + value + "%"
	} else if filter.GetOperator().IsNotEquals() {
		queryPart += "!= ?"
		param = value
	} else if filter.GetOperator().IsGreaterThan() {
		queryPart += "> ?"
		param = value

	} else if filter.GetOperator().IsGreaterThanOrEqual() {
		queryPart += ">= ?"
		param = value

	} else if filter.GetOperator().IsLowerThan() {
		queryPart += "< ?"
		param = value

	} else if filter.GetOperator().IsLowerThanOrEqual() {
		queryPart += "<= ?"
		param = value

	} else {
		queryPart += fmt.Sprintf("%s ?", filter.GetOperator())
		param = value

	}
	return queryPart, param
}
