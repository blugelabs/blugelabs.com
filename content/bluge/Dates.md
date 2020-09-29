---
title: "Dates"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 50
---

Bluge includes support indexing and searching `DateTime` (time.Time) values.  It is important to understand that one should only use the Date field type when one needs to perform `DateRange` queries.  If you only intend to perform exact matches, you should convert your date to a string representation and use the `keyword` analyzer.

## Indexing

To start we'll create a document with identifier `a`:

```
doc := bluge.NewDocument("a")
```

Now, let's add a `DateTime` value to this document:

```
now := time.Now()
doc.AddField(bluge.NewDateTimeField("updated", now)
```

The first argument to the `NewDateTimeField` name is the field name `age`.  The second argument is the `time.Time` value to be indexed.

**NOTE**: Because Bluge's internal representation encodes signed 64-bit nanoseconds relative to the Unix epoch (`1970-01-01 00:00:00 +0000 UTC`), accurate encoding of `time.Time` values has nanosecond precision in the range (`1677-09-21 00:12:43.145224192 +0000 UTC` to `2262-04-11 23:47:16.854775807 +0000 UTC`)

## Searching

To search for documents containing a `DateTime` value within a range, we use the `NewDateRangeQuery`.  Here is an example, if we wanted to search for values after `2020-01-01` and before `2021-01-01`:

```
start2020, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
start2021, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
q := bluge.NewDateRangeQuery(start2020, start20201)
```

The first argument is the minimum date time (inclusive), and the second argument is the maximum date time (exclusive).  If you need to control the inclusive/exclusive boundaries, use the longer form:

```
NewDateRangeInclusiveQuery(start, end time.Time, startInclusive, endInclusive bool)
```

## Partially Open Ranges

To perform a `DateRangeQuery` with a partially open range, the time.Time zero-value can be used as a minimum or maximum value.