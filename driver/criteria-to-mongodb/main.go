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
	if mappings == nil {
		mappings = make(map[string]string)
	}

	if criteria.HasFilters() {
		var andElements []map[string]any

		for _, filter := range criteria.GetFilters().GetValue() {
			var field, ok = mappings[filter.GetField().String()]
			if !ok {
				field = filter.GetField().String()
			}
			var elemtQuery, filterValue = ctmsc.generateWhereQuery(filter)
			andElements = append(andElements, map[string]any{
				field: map[string]any{
					elemtQuery: filterValue,
				},
			})
		}

		query = append(query, map[string]any{
			"$match": map[string]any{
				"$and": andElements,
			},
		})

	}
	if criteria.HasOrder() {
		query = append(query, map[string]any{
			"$sort": map[string]any{
				criteria.GetOrder().GetOrderBy().GetByField(): ctmsc.generateSortTypeMongo(criteria.GetOrder().GetOrderType()),
			},
		})
	}

	// page number
	var pageSizePointer = criteria.GetPageSize()
	var pageNumberPointer = criteria.GetPageNumber()
	if pageSizePointer != nil && pageNumberPointer != nil {
		var pageSize = *pageSizePointer
		var pageNumber = *pageNumberPointer
		query = append(query, map[string]any{
			"$skip": pageSize * (pageNumber - 1),
		})
	}

	// page size
	if pageSizePointer != nil {
		var pageSize = *pageSizePointer
		query = append(query, map[string]any{
			"$limit": pageSize,
		})
	}

	// project result
	if fieldsToSelect == nil || len(fieldsToSelect) == 0 {
		return query
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

func (ctmsc *CriteriaToMongodb) generateWhereQuery(
	filter gocriteria.Filter,
) (fieldQuery string, valueQuery any) {
	var filterValue = filter.GetValue().GetValue()

	if filter.GetOperator().IsContains() {
		fieldQuery = "$regex"
		valueQuery = filterValue
	} else if filter.GetOperator().IsNotContains() {
		fieldQuery = "$not"
		valueQuery = map[string]any{
			"$regex": filterValue,
		}
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
