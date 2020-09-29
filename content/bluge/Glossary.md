---
title: "Glossary"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 100
---

Analyzer:
: An analyzer transforms input text into a set of terms.  At the time a document is indexed, an analyzer is used to produce the set of terms that should be placed into the index.  At search time, an analyzer can be used to produce search terms for certain types of queries.

Index
: An index is the main unit upon which you can perform a search.  Indexes maintained structured information about the documents that have placed in them, in such a way that searches can be performed on them.

Query
: A query is the part of a search that describes which documents should be returned.  Simple queries can be combined with boolean operators to create more complex queries.

Search Request
: A search request combines a query with additional parameters about the type of search you're performing, such as the number of results you'd like returned, or a set of aggregations to compute over the result set. 

