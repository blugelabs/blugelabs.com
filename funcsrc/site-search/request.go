package main

import (
	"fmt"
	"github.com/blugelabs/bluge/search"
	"github.com/blugelabs/bluge/search/aggregations"
	"time"

	"github.com/blugelabs/bluge"
	querystr "github.com/blugelabs/query_string"
)

const resultsPerPage = 10
const roundDurationTo = 500 * time.Microsecond

type Filter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchRequest struct {
	Query   string    `json:"query"`
	Filters []*Filter `json:"filters"`
	Page    int       `json:"page"`
}

func (r *SearchRequest) buildFilterClauses() (rv []bluge.Query) {
	for _, filter := range r.Filters {
		switch filter.Name {
		case "type":
			rv = append(rv, bluge.NewTermQuery(filter.Value).SetField(filter.Name))
		}
	}

	return rv
}

func (r *SearchRequest) SizeOffset() (size, offset int) {
	return resultsPerPage, (r.Page - 1) * resultsPerPage
}

func (r *SearchRequest) BlugeRequest() (bluge.SearchRequest, error) {
	userQuery, err := querystr.ParseQueryString(r.Query, querystr.DefaultOptions())
	if err != nil {
		return nil, fmt.Errorf("errror parsing query string '%s': %v", r.Query, err)
	}

	if r.Page < 1 {
		r.Page = 1
	}

	size, offset := r.SizeOffset()

	filters := r.buildFilterClauses()

	q := bluge.NewBooleanQuery().
		AddMust(userQuery).
		AddMust(filters...)

	blugeRequest := bluge.NewTopNSearch(size, q).
		IncludeLocations().
		WithStandardAggregations().
		SetFrom(offset).
		ExplainScores()

	blugeRequest.AddAggregation("type", aggregations.NewTermsAggregation(search.Field("type"), 5))

	return blugeRequest, nil
}