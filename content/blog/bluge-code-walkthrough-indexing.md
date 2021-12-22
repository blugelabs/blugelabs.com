---
title: "Code Walk-through - Indexing"
date: 2021-12-22T10:03:41-04:00
draft: false
author: Marty Schoch
---

On December 21st we recorded our first code walk-through session.  This session focuses on a very simple application which indexes a single document.  We step into the Bluge code and discuss the important details behind the indexing process.

<!--more-->

{{< youtube 6XXW9-cls1s >}}

## References

- [Main Program](https://gist.github.com/mschoch/c9fb385570a85286357a51555a814e4f)
- https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/writer.go#L226
- https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/segment_plugin.go#L64
- https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/writer.go#L74
- Prepare Interim
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L36
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L112
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L252
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L339
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L401
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L453
- Convert Interim to Final
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L555
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L645
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L693
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L782
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/write.go#L57
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/write.go#L155
  - https://github.com/blugelabs/ice/blob/3993755d45c2912a10025d426f651981446816da/new.go#L87
- Introducer
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/introducer.go#L43
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/introducer.go#L338
- Persister
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/persister.go#L26
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/persister.go#L219
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/persister.go#L311
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/directory_fs.go#L113
  - https://github.com/blugelabs/bluge/blob/f89eff45771cfe7cb151b01bc6204028f3c06be9/index/deletion.go#L38
- External Data Structure Libraries
  - [Roaring Bitmaps](https://github.com/RoaringBitmap/roaring)
  - [Vellum Finite State Transducer](https://github.com/blevesearch/vellum)

