---
title: "Indexing"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 20
---

## Config

To start working with a Bluge index, one always begins by creating the appropriate `Config` structure.  To create a default config structure for working with an index stored on the filesystem use the following:

```
config := bluge.DefaultConfig(path)
```

## Writers

Next, to modify an index with this config structure, we need to open a `Writer`, which can be done as follows:

```
writer, err := bluge.OpenWriter(config)
if err != nil {
	log.Fatalf("error opening writer: %v", err)
}
```

Writer's hold an exclusive-lock on their underlying directory which prevents other processes from opening a writer while this one is still open.  This does not affect `Readers` that are already open, and it does not prevent new `Readers` from being opened, but it does mean care care should be taken to close the `Writer` when you done:

```
defer func() {
	err = writer.Close()
	if err != nil {
		log.Fatalf("error closing writer: %v", err)
	}
}()
```

Now that we have an open `Writer`, we can add `Documents` to the index.

In Bluge, a `Document` is simply a collection of `Fields`.  All `Fields` have a string identifier we call the field name.  There is no requirement that `Documents` have `Fields` with any special names, but as a convention it is useful for `Documents` to have a common field named `_id`.  To aid in following this convetion, a helper method exists to create documents with this field:

```
doc := bluge.NewDocument("a")
```

Typically, an application will add other fields to the document, here is a simple example adding a text field:

```
doc.AddField(bluge.NewTextField("name", "bluge")
```

Now our document is ready to be placed into the index.  The most common way to update a document in the index using the `Update` method:

```
err = writer.Update(doc.ID(), doc)
if err != nil {
	log.Fatalf("error updating document: %v", err)
}
```

The `Update` method takes two arguments, the first is an identifier `Term`, and the second is a `Document`.  The identifier `Term` tells Bluge how to identify which documents this will replace.  The `Document` contains a helper method `ID()` to return the identifier used when calling `NewDocument(id)`.  In this example, we are updating the document identified by the field `_id` and value `a`.  By using the `Update` method we ensure there is only ever one document with this identifier.

For advanced users an `Insert` method is offered which only takes the `Document` parameter, however this should only be used in cases where it is known that there is no existing `Document` with the same identifier.

Another important capability is to `Delete` documents from the index.  We can delete the document we just updated using:

```
err = writer.Delete(doc.ID())
if err != nil {
	log.Fatalf("error deleting document: %v", err)
}
```

The `Delete` method takes just one parameter.  Documents matching the identifying `Term` are removed.

## Batches

In Bluge, higher throughput can be achieved by indexing `Documents` in larger batches.  `Batches` also provide atomicity, as you are guaranteed that either all changes in a batch are applied together, or none of them are, and an error is returned.

To create a batch:

```
batch := bluge.NewBatch()
```

Batches offer the same basic operations as the Writer:

```
batch.Insert(doc)
batch.Update(doc.ID(), doc)
batch.Delete(doc)
```

When you are ready to execute a `Batch`:

```
err = indexWriter.Batch(batch)
if err != nil {
	log.Fatalf("error executing batch: %v", err)
}
```

It is important for applications to not operate on the same document multiple times in a batch.  For example, one should not `Update` and `Delete` the same identifier in a single batch.

After a batch has completed execution (Batch() method has returned), a `Batch` can be reused by invoking the `Reset()` method:

```
batch.Reset()
```