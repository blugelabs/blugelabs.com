---
title: "Sorting"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 70
---

By default, search results in Bluge are sorted by the `DocumentMatch`'s score, in descending order.  However, users can specify their own sort order.

## Preparing the Index

In order to sort on the value of a field, that field must have been indexed with the correction options.  Here we see a keyword field named `style` indexed with the additional `Sortable()` property:

```
doc.AddField(bluge.NewKeywordField("style", style).Sortable())
```

## Custom Sort

First, create your `SearchRequest` as usual:

```
req := bluge.NewTopNSearch(10, query)
```

Then, add a custom sort order:

```
req.SortBy([]string{"-name"})
```

The `SortBy` method takes a slice of strings describing the custom sort order.  In this simple form, strings are interpreted as field names, and the prefix `-` indictates descending order.  A special value of `_score` is also recognized as the `DocumentMatch`'s score value.

This can be extended to sort by multiple field values.  Here we sort by name ascending, score descending, then finally by document id.

```
req.SortBy([]string{"name", "-_score", "_id"})
```

## Advanced Sorting

The slice of strings format we just saw is convenient, but does not allow for the full capabilities.  A more advanced form can be used:

```
req.SortByCustom(search.SortOrder{
		search.SortBy(search.Field("name")).Desc().MissingFirst(),
	})
```

The `SortByCustom` method lets us specify a `search.SortOrder`.  The `search.SortOrder` is a slice of `*search.Sort` values, which we create using the `search.SortBy()` helper method.  This method takes a single argument which is the extendable `search.TextValueSource` interface.  To refer to a field value indexed with the `Sortable` option, we can use the `search.Field()` helper, which takes the field name as it's argument.

The `*search.Sort` value defaults to ascending order, with missing values last.  These can be flipped using the `Desc()` and `MissingFirst()` methods respectively.

While this form seems quite complex, the power comes from the fact that `search.TextValueSource` can be implemented by your application, to allow sorting by custom functions, which can themselves refer to indexed values.

