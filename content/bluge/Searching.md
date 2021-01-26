---
title: "Searching"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 30
---

## Config

To start working with a Bluge index, one begins by creating the appropriate `Config` structure.  To create a default config structure for working with an index stored on the filesystem use the following:

```
config := bluge.DefaultConfig(path)
```

## Readers

Next, to search an index with this config structure, we need to open a `Reader`, which can be done as follows:

```
reader, err := bluge.OpenReader(beerCfg)
if err != nil {
	log.Fatalf("unable to open reader: %v", err)
}
```

The `Reader` represents a stable snapshot of the index a point in time.  This means that changes made to the index after the reader is obtained never affect the results returned by this reader.  This also means that this `Reader` is holding onto resources and **MUST** be closed when it is no longer needed.

```
err = reader.Close()
if err != nil {
	log.Fatalf("error closing reader: %v", err)
}
```

## Queries

Now that we have a `Reader` ready for searching, we need to describe to Bluge what documents we're looking for.  This is done with an object satisfying Bluge's `Query` interface.  A common query used for finding text matches in the index is a `MatchQuery`:

```
q := NewMatchQuery("bluge")
```

This creates a MatchQuery looking for the term `bluge`.  We have not specified a field restriction, so this will search the `DefaultSearchField` in the `Config` structure.  To explicitly restrict this to a particular field we can use:

```
q.SetField("name")
```

## Search Requests

While the `Query` describes *what* we're looking for, the `Request` describes how we want the search executed.  For example, do we want the top 10 hits?  

```
req := NewTopNSearch(10, q)
```

The first argument is the number of hits we want, and the second is the `Query`.

Alternatively, maybe want to see page 2 of the results by skipping over the first 10 hits?  

```
req := NewTopNSearch(10, q).SetFrom(10)
```

Or do we want to see all matches?

```
req := NewAllMatches(q)
```

All of these return a structure satisfying Bluge's `SearchRequest` interface, so we work with them the same.  To execute the search:

```
dmi, err := reader.Search(context.Background(), req)
if err != nil {
	log.Fatalf("error executing search: %v", err)
}
```

The first argument is a `context.Context`, which allows you to cancel (or timout) a search if needed.  The second argument is the `SearchRequest` we just built.  Executing a search returns a `DocumentMatchIterator`, which allows us to iterate through the `DocumentMatches` that satisfy the request.

## Iterating through Results

The standard pattern to iterate through search results looks like this:

```
next, err := dmi.Next()
for err == nil && next != nil {
	err = next.VisitStoredFields(func(field string, value []byte) bool {
		if field == "_id" {
			fmt.Println(string(value))
		}
		return true
	})
	if err != nil {
		log.Fatalf("error accessing stored fields: %v", err)
	}
	next, err = dmi.Next()
}
if err != nil {
	log.Fatalf("error iterating results: %v", err)
}
```

Each time we invoke `Next()` we are returned the next `DocumentMatch`.  `DocumentMatches` each have an identifying number, however this number **ONLY** has meaning inside the context of this reader's snapshot.

Here we see the `VisitStoreFields` helper method being used to invoke a callback for each stored field.  In this example, we only look for a field named `_id` which allows us to print the documents actual identifier.

After the last hit has been seen by the iterator, `nil` is returned by `Next()`.

## Standard Aggregations

Often it is desirable to return a set of standard aggregations along with the search results.  These include:
- Total Number Documents that matched the Query (not the number returned by this request)
- Maximum score of all matches (not the highest returned by this request)
- Search Duration (time.Duration to execute this request)

To include this information, add the following to your `SearchRequest`:

```
req := bluge.NewTopNSearch(10, q).
		WithStandardAggregations()
```

Then, after iterating through the all of the `DocumentMatches` you can access the values from the `Aggregations` structure:

```
total := dmi.Aggregations().Count()
maxScore := dmi.Aggregations().Metric("max_score")
duration := dmi.Aggregations().Duration()
```

For more information, see the page on [Aggregations](/bluge/aggregations/)

## Near-Real-Time Readers

If your application already has a `Writer` open to modify an index, you can get a reader using a method on the Writer:

```
reader, err := writer.Reader()
if err != nil {
	log.Fatalf("error accessing reader: %v", err)
}
```

This `Reader` is a special reader that has access to portions of the index not yet persisted to disk.

