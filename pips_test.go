package pips

import (
	"testing"
)

func TestPointInPolygon(t *testing.T) {
	shapes := getShapes("test_files/tl_2014_us_state.shp")
	if len(shapes) != 56 {
		t.Error("wrong number of shapes")
	}
	point_y := shp.Point{-82, 39}
	shapes_return := FindPoint(&point_y, &shapes)
	if len(shapes_return) != 1 {
		t.Error("Wrong number of points when searching a point in West Virginia")
	}
	if shapes_return[0].field_values[6] != "West Virginia" {
		t.Error("Failed to find a point in West Virginia")
	}
	point_y = shp.Point{-88, 40}
	shapes_return = FindPoint(&point_y, &shapes)
	if len(shapes_return) != 1 {
		t.Error("Wrong number of points when searching a point in Illinois")
	}
	if shapes_return[0].field_values[6] != "Illinois" {
		t.Error("Failed to find a point in Illinois")
	}
}
