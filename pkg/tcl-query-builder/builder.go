package tcl_query_builder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	conditions []string
	orders     []struct {
		key string
		asc bool
	}
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) AddCondition(condition string) {
	qb.conditions = append(qb.conditions, condition)
}

// Order добавляет сортировку в массив, принимает 2 параметра: ключ и буливое asc.
func (qb *QueryBuilder) Order(key string, asc bool) *QueryBuilder {
	qb.orders = append(qb.orders, struct {
		key string
		asc bool
	}{key, asc})

	return qb
}

// Query возвращает строку запроса, которую можно использовать в запросе к Temporal.
func (qb *QueryBuilder) Query() string {
	var query string

	if len(qb.conditions) > 0 {
		query = strings.Join(qb.conditions, " and ")
	}

	if len(qb.orders) > 0 {
		var orders []string
		for _, order := range qb.orders {
			orders = append(orders, fmt.Sprintf("%s %s", order.key, orderDirection(order.asc)))
		}

		query = fmt.Sprintf("%s order by %s", query, strings.Join(orders, ", "))
	}

	return query
}

func Eq(field string, value interface{}) string {
	return fmt.Sprintf("%s = %s", field, valueToString(value))
}

func Neq(field string, value interface{}) string {
	return fmt.Sprintf("%s != %s", field, valueToString(value))
}

func Gte(field string, value interface{}) string {
	return fmt.Sprintf("%s >= %s", field, valueToString(value))
}

func Lte(field string, value interface{}) string {
	return fmt.Sprintf("%s <= %s", field, valueToString(value))
}

func valueToString(value interface{}) string {
	switch value.(type) {
	case string:
		value = strings.Replace(value.(string), "'", "''", -1)
		return fmt.Sprintf("'%s'", value)
	case int, int32, int64:
		return fmt.Sprintf("%d", value)
	default:
		return fmt.Sprintf("%v", value)
	}

	return ""
}

func orderDirection(asc bool) string {
	if asc {
		return "asc"
	}

	return "desc"
}
