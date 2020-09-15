---
title: "Introducing Bluge"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
---

Search is an important part of the Go ecosystem, as evidenced by the [popularity of the Bleve project](https://github.com/blevesearch/bleve/stargazers). Over the summer, I created a new open-source indexing and search library for Go named Bluge.  Bluge is built on the same core technology of Bleve that many companies use to power their search features. I’ve applied my experience building and maintaining Go code for the last eight years to transform Bluge into something I consider a more solid foundation for the future.

Below are some of the more significant changes.  Want more details? [Become a sponsor](https://github.com/sponsors/mschoch) and get an invite to the introductory session on 2020-09-29.

## Removing Code

One of the major challenges in maintaining a library shared by a community is properly defining the scope of the library itself.  In several cases, Bleve attempted to do too much, and found itself with functionality that not all users could agree upon.  In these cases, the application needed to take ownership of these responsibilities to get the behavior they wanted.  The following items have been removed from Bluge:

- Index Mapping
- Query String
- HTTP Handlers
- JSON (Un)Marshaling
- Internal supplemental Key/Value storage
- String-based registry of functionality

## Cleaning up the Codebase

Over time all large codebases accumulate some cruft.  This is an opportunity to review the codebase as a whole and do some housekeeping.  We’ve focused on the following:

- [Fewer, larger packages](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_consider_fewer_larger_packages)
- Use a strongly typed configuration structure
- Leverage static analysis tools to maintain code quality (via [golangci-lint](https://github.com/golangci/golangci-lint))

## Enhancements and New Features

Removing unnecessary code and cleaning up the codebase are nice, but what I think most people will be excited about are the new capabilities:

- A new Directory abstraction separating the core library from the OS
  - This enables an in-memory only index with no other changes
- Multi-process Indexes
  - External processes can open the index read-only and search the latest data persisted to disk
- Online backup
  - Backup your index without closing it
  - Combined with multi-process indexes, this can be done from another process
- [BM25](https://en.wikipedia.org/wiki/Okapi_BM25) Scoring
  - Significantly improved scoring, compared to Bleve
  - With the additional ability to control scoring at the level of each query clause
- Aggregation Framework
  - Bucketing aggregations by terms, numeric range, date range
  - Expanded metrics: count, min, max, avg, weighted-average
  - Cardinality estimation (via [HyperLogLog++](https://en.wikipedia.org/wiki/HyperLogLog))
  - Quantile approximation (via [T-Digest](https://raw.githubusercontent.com/tdunning/t-digest/master/docs/t-digest-paper/histo.pdf))
  - Interfaces allow you to define your own aggregations
  - Attach metrics to buckets at any level
- Source abstraction for Sorting and Aggregating
  - Sort and aggregate over computed values, not just indexed values

## The Plan

Like Bleve, Bluge will be released under the [Apache-2.0 License](https://www.apache.org/licenses/LICENSE-2.0).

This first release will be tagged Developer Preview 1, and is scheduled for 2020-09-29.

While we gather feedback, we have other significant features planned for a second developer preview.  Ultimately, we are roughly targeting a production ready release by the end of the year.

Are you excited to check this out?  Do you plan to download the code?  Please, show your support for my work by [becoming a sponsor](https://github.com/sponsors/mschoch).

Your sponsorship makes it possible for me to continue working on projects like Bluge!

**SPONSOR EXCLUSIVE**: Invitation to attend a series of zoom calls where I will present some Bluge content, and then open up an interactive session.  The first session will be an introduction to Bluge on 2020-09-29.  Don't miss the first session, [become a sponsor right now!](https://github.com/sponsors/mschoch)
