---
title: "A plan to release Bleve v1.0.0 and adopt Go modules"
date: 2020-03-16T15:03:41-04:00
draft: false
author: Marty Schoch
---

Bluge Labs has been up and running for a few months now, and I'm happy to finally be able to share one of the major projects I've been working on.  Bleve has seen wide adoption despite never having an official `v1.0.0` release.  Additionally, in the past year Go modules have matured, and the community has increasingly requested proper support from the library.

This is our [proposal](https://github.com/blevesearch/bleve/issues/1350) to address both of these issues.  We have given special attention to keep backwards compatibility with the existing Bleve APIs as well as the existing file formats.  However, this change will break traditional GOPATH builds.  We had hoped to avoid this, but after evaluating the trade-offs we have concluded that this is right decision and the right time to make it.

We value feedback from the community, and hope to move forward on this proposal with their support.