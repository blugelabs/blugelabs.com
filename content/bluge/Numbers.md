---
title: "Numbers"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 40
---

Bluge includes support indexing and searching numeric values.  It is important to understand that one should only use the Numeric field type when one needs to perform NumericRange queries.  If you only intend to perform exact matches, you should convert your number to a string representation and use the `keyword` analyzer.

## Indexing

To start we'll create a document with identifier `a`:

```
doc := bluge.NewDocument("a")
```

Now, let's add a numeric value to this document:

```
doc.AddField(bluge.NewNumericField("age", 0.1))
```

The first argument to the `NewNumericField` name is the field name `age`.  The second argument is the `float64` value to be indexed.

**NOTE**: Because Bluge's internal representation encodes `float64` values, accurate encoding of type converted integer values is limited to &#177;2<sup>53</sup>.

## Searching

To search for documents containing a numeric value within a range, we use the `NumericRangeQuery`.  Here is an example, if we wanted to search for values greater than or equal to zero, and less than one:

```
q := NewNumericRangeQuery(0, 1)
```

The first argument is the minimum value (inclusive), and the second argument is the maximum value (exclusive).  If you need to control the inclusive/exclusive boundaries, use the longer form:

```
NewNumericRangeInclusiveQuery(min, max float64, minInclusive, maxInclusive bool)
```

## Partially Open Ranges

To perform a `NumericRangeQuery` with a partially open range, the following constant values are provided:

```
bluge.MinNumeric
bluge.MaxNumeric
```
