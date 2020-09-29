---
title: "Aggregating"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 90
---

Bluge supports a powerful framework for computing aggregated values over the set of `DocumentMatches`.

## Terminology

In Bluge, the aggregation framework relies heavily on two concepts `Buckets` and `Metrics`.

A `Bucket` is simply a set of documents matching some criteria.

A `Metric` is some value (or set of values) computed over a `Bucket`.

There is one implicit bucket defined, which is the entire result set of your search.

Some aggregations (which we refer to as bucketing aggregations) define new sub-buckets inside this top-level bucket.  These sub-buckets could either be staticly defined at search time, or dynamically defined based on the data.

Other aggregations (which we refer to as metric aggregations) compute values on buckets.

## Bucketing Aggregations

### Terms Aggregation

The terms aggregation typically operates on field data.  Each term seen becomes it's own bucket, and by default the count metric is applied to each bucket.  Finally, at the conclusion of the search, these buckets are sorted by their counts descending, and the top N buckets are returned as part of the result.

For example, consider a set documents describing products.  Each product has a keyword field named `category`, indexed with the sortable option.  When a user searches the products, we can compute a terms aggregation on the `category` field, and display to the user the top 5 categories within their search results, and a count of how many products were in each category.  This is often used as a way for users to drill deeper into the results, by refining their search filter interactively.

### Numeric Range Aggregation

The numeric range aggregation also typically operates on field data.  A query time a set of buckets is statically defined, which describe interesting numeric ranges.  The aggregation by default includes the count metric, keeping track of how many documents had a numeric field value within the range.

### Date Range Aggregation

The date range aggregation also typically operates on field data.  A query time a set of buckets is statically defined, which describe interesting date ranges.  The aggregation by default includes the count metric, keeping track of how many documents had a date time field value within the range.

## Metric Aggregations

### Basic

The following basic single-value metrics are supported:

- sum
- min
- max
- avg
- weighted avg

### Special

A few special case aggregations are supported:

- count (sum of 1 per document)
- duration (time.Duration computed since the start of the search)

### Cardinality Estimation

The cardinality estimation metric can be used to count the number of distinct values seen, in a memory efficient way.

### Quantile Approximation

The quantil approximation metric can be used to approximate quantiles in a memory efficient way.

## Nesting

Buckets and Metrics can be nested in arbitrary and powerful ways.  

For example, imagine we have a set of documents describing beers.  Each beer has a field named `style` describing the style (lager, ale, lambic, etc).  Each beer also has a numeric field named `abv` describing the beer's alcohol by volume.  One could run a `MatchAll` query across the beers, compute a `Terms Aggregation` on the `style` field, and then nest the `Quantile Approximation` metric inside each of those buckets.  The result would be that we could report the median (50th) and 99th percentile ABV for each different style of beer.

## Custom Sources

All the aggregations discussed thus far operate on extendable interfaces, not directly on field values.

This allows aggregations to work on custom values computed by your application, which can themselves use field value as inputs.

It also allows for filtering out undesirable values, or replacing missing values with alternates.

## Extending the Framework

The core `Aggregation` and `Calculator` types used to define all of this functionality are exposed as interfaces, allowing your application the full power to define their own behavior. 