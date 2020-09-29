---
title: "Introduction"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 10
---

# What is Bluge?

Bluge is an indexing/search library for Go.

Developers can insert/update/delete documents in the index.  Documents are composed of fields.  Bluge support fields with text, numeric, date, and geo point values.

Developers can then search the index to find matching documents.  Bluge supports many types of queries, which can be composed into more complex queries with boolean operators.

Bluge supports scoring full-text searches using the [BM25](https://en.wikipedia.org/wiki/Okapi_BM25) ranking function.  Scoring can be customized on each query clause.

Bluge supports highlighting <mark>matching values</mark> in the original text.

By default, Bluge sorts the results by score, descending.  Bluge allows you to customize the sort order, sorting on fields indexed or custom functions using field values.

Bluge supports an advanced aggregation framework, allowing the developer to compute aggregated values over the set of matching documents.

# History

Bluge originated as an evolution of the [Bleve](https://github.com/blevesearch/bleve) search library. Bleve has seen wide adoption, but that success comes with some downside.  As companies ship products using Bleve, they require backwards compatibility, and this has led to a slower more incremental progression.

Bluge is an attempt to break out of this model, making many breaking changes all at once. 
  