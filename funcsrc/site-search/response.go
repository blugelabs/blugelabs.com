package main

import (
	"fmt"
	"github.com/blugelabs/bluge/search/highlight"
	"math"

	"github.com/blugelabs/bluge/search"
)

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type DocumentMatch struct {
	Document Document     `json:"document"`
	Score    float64             `json:"score"`
	Expl     *search.Explanation `json:"explanation"`
	ID       string              `json:"id"`
}

type AggregationValue struct {
	DisplayName string `json:"display_name"`
	FilterName  string `json:"filter_name"`
	Count       uint64 `json:"count"`
	Filtered    bool   `json:"filtered"`
}

type Aggregation struct {
	DisplayName string              `json:"display_name"`
	FilterName  string              `json:"filter_name"`
	Values      []*AggregationValue `json:"values"`
}

type SearchResponse struct {
	Query        string                  `json:"query"`
	Total        uint64                  `json:"total"`
	TopScore     float64                 `json:"top_score"`
	Hits         []*DocumentMatch        `json:"hits"`
	Duration     string                  `json:"duration"`
	Aggregations map[string]*Aggregation `json:"aggregations"`
	Message      string                  `json:"message"`
	PreviousPage int                     `json:"previousPage,omitempty"`
	NextPage     int                     `json:"nextPage,omitempty"`
}

func NewSearchResponse(query string, dmi search.DocumentMatchIterator,
	highlighter *highlight.SimpleHighlighter, searchRequest *SearchRequest) (*SearchResponse, error) {
	rv := &SearchResponse{
		Query: query,
	}

	next, err := dmi.Next()
	for err == nil && next != nil {
		var dm DocumentMatch
		err = next.VisitStoredFields(func(field string, value []byte) bool {
			if field == "_id" {
				dm.ID = string(value)
			} else if field == "title" {
				dm.Document.Title = string(value)
			} else if field == "content" {
				dm.Document.Content = string(value)
				if contentLocations, ok := next.Locations["content"]; ok {
					fragment := highlighter.BestFragment(contentLocations, value)
					if len(fragment) > 0 {
						dm.Document.Content = fragment
					}
				}

			} else if field == "type" {
				dm.Document.Type = string(value)
			}
			return true
		})
		if err != nil {
			return nil, fmt.Errorf("error visiting stored fields: %v", err)
		}
		dm.Score = next.Score
		dm.Expl = next.Explanation
		rv.Hits = append(rv.Hits, &dm)
		next, err = dmi.Next()
	}
	if err != nil {
		return nil, fmt.Errorf("error iterating matches: %v", err)
	}

	rv.AddAggregations(dmi.Aggregations(), searchRequest.Filters)

	return rv, nil
}

func (s *SearchResponse) AddPaging(aggs *search.Bucket, page int) {
	numPages := int(math.Ceil(float64(aggs.Count()) / float64(resultsPerPage)))
	if numPages > page {
		s.NextPage = page + 1
	}
	if page != 1 {
		s.PreviousPage = page - 1
	}

	if page != 1 {
		s.Message = fmt.Sprintf("Page %d of ", page)
	}
	s.Message += fmt.Sprintf("%d results (%s)", aggs.Count(),
		aggs.Duration().Round(roundDurationTo))
}

func (s *SearchResponse) AddAggregations(aggs *search.Bucket, filters []*Filter) {
	s.Total = aggs.Count()
	s.TopScore = aggs.Metric("max_score")
	s.Duration = aggs.Duration().String()

	s.Aggregations = make(map[string]*Aggregation)
	s.buildAggregation(aggs, "type", filters)
}

func (s *SearchResponse) buildAggregation(aggs *search.Bucket, name string, filters []*Filter) {
	agg := &Aggregation{
		DisplayName: displayName(name),
		FilterName:  name,
	}

	for _, bucket := range aggs.Buckets(name) {
		aggVal := &AggregationValue{
			DisplayName: displayName(bucket.Name()),
			FilterName:  bucket.Name(),
			Count:       bucket.Count(),
		}
		for _, f := range filters {
			if f.Name == name && f.Value == bucket.Name() {
				aggVal.Filtered = true
			}
		}
		agg.Values = append(agg.Values, aggVal)
	}

	s.Aggregations[name] = agg
}

func displayName(in string) string {
	switch in {
	case "type":
		return "Type"
	case "blog":
		return "Blog"
	case "page":
		return "Page"
	case "bluge":
		return "Bluge Docs"
	}
	return in
}
