pips
====

Point-In-Polygon for Shapefiles, works with ESRI Shapefiles that contain polygons.

http://www.esri.com/library/whitepapers/pdfs/shapefile.pdf

ESRI shapefiles are provided free by the US Census.  
They have GIS info for States, Counties, Metropolitan Areas, Places (cities and towns), and congressional districts.
https://www.census.gov/geo/maps-data/data/tiger-line.html

The US States shapefile is included the test_data folder.

Given a GPS coordinate, this library will return the shape and data for the polygon(s) that it intersects with.  

This library works by using a goroutine to check all shape bounding boxes concurrently.  
If a shape's bounding box is intersected, then pips will perform a point in polygon computation for that shape.

With enough shapes and few availible processors, it could be a good idea to use an rtree to index the bounding boxes.
This library would be straightforward to integrate with pips: http://dhconnelly.com/rtreego/

In theory, with enough processing capacity, you could get O(1) performance out of a concurrent bounding box search. 
In a such a system that would allow O(1) performance, an rtree could yeild worse performance.  
For example, a concurrent rtree traversal with a maximum tree depth of D would yeild O(D) worst case performance.

RESTRICTIONS:
This has only been tested with Polygon shapes, all other shapes may cause indeterminate behavior.
This library requires this package: https://github.com/benmullin333/go-shp
