package query

import (
	"fmt"
	"strings"
)

type ArangoQueryBuilder struct {
	query       strings.Builder
	counterName string
}

func NewForQuery(repositoryName string, counterName string) *ArangoQueryBuilder {
	aqb := &ArangoQueryBuilder{
		counterName: counterName,
	}

	aqb.query.WriteString(fmt.Sprintf("FOR %s IN %s", aqb.counterName, repositoryName))

	return aqb
}

func (aqb *ArangoQueryBuilder) Search() *ArangoQuerySearch {
	aqb.query.WriteString(" ")
	aqb.query.WriteString("SEARCH")
	return NewArangoQuerySearch(aqb)
}

func (aqb *ArangoQueryBuilder) Filter(fieldName string, condition string, value interface{}) *ArangoQueryFilter {
	aqb.query.WriteString(" ")

	switch value.(type) {
	case string:
		aqb.query.WriteString(fmt.Sprintf("FILTER %s.%s %s %q", aqb.counterName, fieldName, condition, value))
	default:
		aqb.query.WriteString(fmt.Sprintf("FILTER %s.%s %s %v", aqb.counterName, fieldName, condition, value))
	}

	return NewArangoQueryFilter(aqb)
}

func (aqb *ArangoQueryBuilder) LIMIT(offset int, count int) *ArangoQueryBuilder {
	aqb.query.WriteString(" ")
	aqb.query.WriteString(fmt.Sprintf("LIMIT %d,%d", offset, count))
	return aqb
}

func (aqb *ArangoQueryBuilder) SortBM25(desc bool) *ArangoQueryBuilder {
	aqb.query.WriteString(" ")
	aqb.query.WriteString(fmt.Sprintf("SORT BM25(%s)", aqb.counterName))

	if desc {
		aqb.query.WriteString(" ")
		aqb.query.WriteString("DESC")
	}

	return aqb
}

func (aqb *ArangoQueryBuilder) Sort(fieldName string, desc bool) *ArangoQueryBuilder {
	aqb.query.WriteString(" ")
	aqb.query.WriteString(fmt.Sprintf("SORT %s.%s", aqb.counterName, fieldName))

	if desc {
		aqb.query.WriteString(" ")
		aqb.query.WriteString("DESC")
	}

	return aqb
}

func (aqb *ArangoQueryBuilder) SortBM25WithFreqScaling(desc bool, k float32, b float32) *ArangoQueryBuilder {
	aqb.query.WriteString(" ")
	aqb.query.WriteString(fmt.Sprintf("SORT BM25(%s, %.2f, %.2f)", aqb.counterName, k, b))

	if desc {
		aqb.query.WriteString(" ")
		aqb.query.WriteString("DESC")
	}

	return aqb
}

func (aqb *ArangoQueryBuilder) Return() *ArangoQueryBuilder {
	aqb.query.WriteString(" ")
	aqb.query.WriteString(fmt.Sprintf("RETURN %s", aqb.counterName))
	return aqb
}

func (aqb *ArangoQueryBuilder) String() string {
	return aqb.query.String()
}

type ArangoQueryFilter struct {
	arangoQueryBuilder *ArangoQueryBuilder
}

func NewArangoQueryFilter(queryBuilder *ArangoQueryBuilder) *ArangoQueryFilter {
	return &ArangoQueryFilter{
		arangoQueryBuilder: queryBuilder,
	}
}

func (aqf *ArangoQueryFilter) And(fieldName string, condition string, value interface{}) *ArangoQueryFilter {
	aqf.arangoQueryBuilder.query.WriteString(" ")
	aqf.arangoQueryBuilder.query.WriteString("AND")
	aqf.arangoQueryBuilder.query.WriteString(" ")

	switch value.(type) {
	case string:
		aqf.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %q", aqf.arangoQueryBuilder.counterName, fieldName, condition, value))
	default:
		aqf.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %v", aqf.arangoQueryBuilder.counterName, fieldName, condition, value))
	}

	return aqf
}

func (aqf *ArangoQueryFilter) Or(fieldName string, condition string, value interface{}) *ArangoQueryFilter {
	aqf.arangoQueryBuilder.query.WriteString(" ")
	aqf.arangoQueryBuilder.query.WriteString("OR")
	aqf.arangoQueryBuilder.query.WriteString(" ")

	switch value.(type) {
	case string:
		aqf.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %q", aqf.arangoQueryBuilder.counterName, fieldName, condition, value))
	default:
		aqf.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %v", aqf.arangoQueryBuilder.counterName, fieldName, condition, value))
	}
	return aqf
}

func (aqf *ArangoQueryFilter) Done() *ArangoQueryBuilder {
	return aqf.arangoQueryBuilder
}

type ArangoQuerySearch struct {
	arangoQueryBuilder *ArangoQueryBuilder
}

func NewArangoQuerySearch(queryBuilder *ArangoQueryBuilder) *ArangoQuerySearch {
	return &ArangoQuerySearch{
		arangoQueryBuilder: queryBuilder,
	}
}

func (aqs *ArangoQuerySearch) Phrase(fieldName string, searchKeyword string, analyzer string) *ArangoQuerySearch {
	aqs.arangoQueryBuilder.query.WriteString(" ")
	aqs.arangoQueryBuilder.query.WriteString(fmt.Sprintf("PHRASE(%s.%s, %q, %q)", aqs.arangoQueryBuilder.counterName, fieldName, searchKeyword, analyzer))
	return aqs
}

func (aqs *ArangoQuerySearch) Condition(fieldName string, condition string, value interface{}) *ArangoQuerySearch {
	aqs.arangoQueryBuilder.query.WriteString(" ")

	switch value.(type) {
	case string:
		aqs.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %q", aqs.arangoQueryBuilder.counterName, fieldName, condition, value))
	default:
		aqs.arangoQueryBuilder.query.WriteString(fmt.Sprintf("%s.%s %s %v", aqs.arangoQueryBuilder.counterName, fieldName, condition, value))
	}

	return aqs
}

func (aqs *ArangoQuerySearch) Or() *ArangoQuerySearch {
	aqs.arangoQueryBuilder.query.WriteString(" ")
	aqs.arangoQueryBuilder.query.WriteString("OR")
	return aqs
}

func (aqs *ArangoQuerySearch) And() *ArangoQuerySearch {
	aqs.arangoQueryBuilder.query.WriteString(" ")
	aqs.arangoQueryBuilder.query.WriteString("AND")
	return aqs
}

func (aqs *ArangoQuerySearch) Done() *ArangoQueryBuilder {
	return aqs.arangoQueryBuilder
}