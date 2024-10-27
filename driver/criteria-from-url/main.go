package criteriafromurl

import (
	gourl "net/url"
	"strconv"

	"regexp"

	"github.com/AndrusGerman/go-criteria"
)

type CriteriaFromUrlConverter struct {
}

func NewCriteriaFromUrlConverter() *CriteriaFromUrlConverter {
	return &CriteriaFromUrlConverter{}
}

func (cfuc *CriteriaFromUrlConverter) ToCriteria(url *gourl.URL) (criteria.Criteria, error) {
	var filters = cfuc.ParseFilters(url.Query())

	var criteriaBuilder = criteria.NewCriteriaBuilder().
		FiltersPrimitive(filters).
		OrderPrimitive(
			url.Query().Get("orderBy"),
			url.Query().Get("order"),
		)

	if !url.Query().Has("order") {
		criteriaBuilder.Order(criteria.NewOrderNone())
	}

	if url.Query().Has("pageSize") {
		var pageSize, err = strconv.Atoi(url.Query().Get("pageSize"))
		if err != nil {
			return nil, err
		}
		criteriaBuilder.PageSize(pageSize)
	}

	if url.Query().Has("pageNumber") {
		var pageNumber, err = strconv.Atoi(url.Query().Get("pageSize"))
		if err != nil {
			return nil, err
		}
		criteriaBuilder.PageNumber(pageNumber)

	}

	return criteriaBuilder.GetCriteria()
}

func (cfuc *CriteriaFromUrlConverter) ToFiltersPrimitives(url *gourl.URL) []criteria.FiltersPrimitive {
	return cfuc.ParseFilters(url.Query())
}

func (cfuc *CriteriaFromUrlConverter) ParseFilters(searchParams gourl.Values) []criteria.FiltersPrimitive {

	var tempFilters = make(map[string]map[string]string)

	for key := range searchParams {
		re := regexp.MustCompile(`filters\[(\d+)]\[(.+)]`)

		matches := re.FindAllStringSubmatch(key, -1)

		if len(matches) > 0 {
			var index = matches[0][1]
			var property = matches[0][2]

			if tempFilters[index] == nil {
				tempFilters[index] = make(map[string]string)
			}
			tempFilters[index][property] = searchParams.Get(key)
		}
	}

	var filtersPrimitives []criteria.FiltersPrimitive

	for _, filter := range tempFilters {
		if filter["field"] == "" && filter["operator"] == "" && filter["value"] == "" {
			continue
		}

		filtersPrimitives = append(filtersPrimitives, criteria.NewFilterPrimitive(
			filter["field"],
			filter["operator"],
			filter["value"],
		))
	}

	return filtersPrimitives
}
