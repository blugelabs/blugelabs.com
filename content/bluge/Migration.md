---
title: "Migrating from Bleve"
date: 2021-12-12T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 93
---

Many projects may migrate to Bluge from Bleve, here we collect information to assist in that effort.

### How can applications support a Bleve-style query string query?

Bluge has moved support for query strings into a separate projec, which can be found here: [query_string](https://github.com/blugelabs/query_string)

### How can applications support a Bleve-style DocIDQuery?

In Bluge, we can query document ID's directly by using a TermQuery on the field named `_id`.  For more information see [this issue](https://github.com/blugelabs/bluge/issues/74).

### How can application list terms from the term dictionary?

Many of the methods used to work with the term dictionary have been renamed and combined into a single more powerful method.  For more information see [this issue](https://github.com/blugelabs/bluge/issues/79).