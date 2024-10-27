package criteriatomongodb

import (
	gocriteria "github.com/AndrusGerman/go-criteria"
)

type CriteriaToMongodb struct {
}

func NewCriteriaToMongodb() *CriteriaToMongodb {
	return &CriteriaToMongodb{}
}

func (ctmsc *CriteriaToMongodb) Convert(
	fieldsToSelect []string,
	criteria gocriteria.Criteria,
	mappings map[string]string,
) []map[string]any {
	var query []map[string]any
	var params []any
	if mappings == nil {
		mappings = make(map[string]string)
	}

	if criteria.HasFilters() {
		var groupFilterValue = make(map[string][]gocriteria.Filter)
		for _, value := range criteria.GetFilters().GetValue() {

			var field, ok = mappings[value.GetField().GetValue()]
			if !ok {
				field = value.GetField().GetValue()
			}
			groupFilterValue[field] = append(groupFilterValue[value.GetField().GetValue()], value)
		}

		var match = make(map[string]any)

		for key, filters := range groupFilterValue {
			match[key] = ctmsc.createQueryForField(filters)
		}
		query = append(query, map[string]any{
			"$match": match,
		})

	}
	if criteria.HasOrder() {
		query = append(query, map[string]any{
			"$sort": map[string]any{
				criteria.GetOrder().GetOrderBy().GetValue(): ctmsc.generateSortTypeMongo(criteria.GetOrder().GetOrderType()),
			},
		})
	}

	var pageSize = criteria.GetPageSize()
	var pageNumber = criteria.GetPageNumber()
	if pageSize != nil && pageNumber != nil {
		query = append(query, map[string]any{
			"$skip": *pageSize,
		})
		params = append(params, (*pageSize)*((*pageNumber)-1))
	}

	if pageSize != nil {
		query = append(query, map[string]any{
			"$limit": *pageSize,
		})
	}

	var project = make(map[string]any)
	for _, field := range fieldsToSelect {
		project[field] = 1
	}
	query = append(query, map[string]any{
		"$project": project,
	})

	return query

	//return query, params
}

func (ctmsc *CriteriaToMongodb) generateSortTypeMongo(orderType gocriteria.OrderType) int {
	switch orderType {
	case gocriteria.ASC:
		return -1
	case gocriteria.DESC:
		return 1
	default:
		return 0
	}
}

func (ctmsc *CriteriaToMongodb) createQueryForField(filters []gocriteria.Filter) map[string]any {
	var queryField = make(map[string]any)
	for _, filter := range filters {
		var field, value = ctmsc.generateWhereQuery(filter)
		queryField[field] = value
	}

	return queryField
}

func (ctmsc *CriteriaToMongodb) generateWhereQuery(
	filter gocriteria.Filter,
) (fieldQuery string, valueQuery any) {
	var filterValue = filter.GetValue().GetValue()

	if filter.GetOperator().IsContains() {
		fieldQuery = "$regex"
		valueQuery = filterValue
	} else if filter.GetOperator().IsNotContains() {
	} else if filter.GetOperator().IsNotEquals() {
		fieldQuery = "$ne"
		valueQuery = filterValue
	} else if filter.GetOperator().IsGreaterThan() {
		fieldQuery = "$gt"
		valueQuery = filterValue

	} else if filter.GetOperator().IsGreaterThanOrEqual() {
		fieldQuery = "$gte"
		valueQuery = filterValue

	} else if filter.GetOperator().IsLowerThan() {
		fieldQuery = "$lt"
		valueQuery = filterValue
	} else if filter.GetOperator().IsLowerThanOrEqual() {
		fieldQuery = "$lte"
		valueQuery = filterValue
	} else {
		fieldQuery = "$eq"
		valueQuery = filterValue
	}
	return fieldQuery, valueQuery
}
