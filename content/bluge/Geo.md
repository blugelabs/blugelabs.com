---
title: "Geo Points"
date: 2020-09-15T14:03:41-04:00
draft: false
author: Marty Schoch
menu:
    bluge:
        weight: 60
---

Bluge includes support indexing and searching geo point values.

## Indexing

To start we'll create a document with identifier `a`:

```
doc := bluge.NewDocument("a")
```

Now, let's add a geo point to this document:

```
doc.AddField(bluge.NewGeoPointField("location", -77.0237, 38.9911))
```

The first argument to the `NewGeoPointField` name is the field name `age`.  The second argument is the `float64` longitude value.  The third argument is the `float64` latitude value.

## Searching

Bluge provides several query types to find documents that relate to locations.

### Geo Distance Query

To search for documents containing a geo point value that is less than a specified distance from another point, we use the `NewGeoDistanceQuery`.  Here is an example, if we wanted to search for documents with geo point values within 100 miles of the White House:

```
q := bluge.NewGeoDistanceQuery(-77.036530, 38.897918, "100mi")
```

The first and second arguments are the longitude and latitude of the query point.  The third argument is a string describing the distance.  Example supported distances:
- "5in" "5inches" 
- "7yd" "7yards" 
- "9ft" "9feet" 
- "11km" "11kilometers"
- "3nm" "3nauticalmiles"
- "13mm" "13millimeters"
- "15cm" "15centimeters"
- "17mi" "17miles" 
- "19m" "19meters" 

### Geo Bounding Box Query

To search for documents containing a geo point value that is within a specified bounding box we use the `NewGeoBoundingBoxQuery`.  Here is an example, if we wanted to search for documents within a bounding box enclosing the USA:

```
q := bluge.NewGeoBoundingBoxQuery(-125.0011, 49.5904, -66.9326, 24.9493)
```

The first and second arguments are the top-left longitude and latitude respectively.  The third and fourth arguments are the bottom right longitude and latitude respectively.

### Geo Bounding Polygon Query

To search for documents containing a geo point value that is within an arbitrary polygon we use the `NewGeoBoundingPolygonQuery`.  Here is an example, if we wanted to search for documents within a polygon in the Mountain View area:

```
q := bluge.NewGeoBoundingPolygonQuery([]geo.Point{
	{
		Lon: 77.607749,
		Lat: 12.974872,
	},
	{
		Lon: 77.6101101,
		Lat: 12.971725,
	},
	{
		Lon: 77.606912,
		Lat: 12.972530,
	},
	{
		Lon: 77.603780,
		Lat: 12.975112,
	},
})
```

The argument provides a slice of geo.Point values describing the polygon.