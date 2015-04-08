package pips

import (
	"github.com/benmullin333/go-shp"
)

//GOMAXPROCS(4)

type PolyData struct {
	poly         shp.Polygon
	index        int
	field_names  []string
	field_values []string
}

func getShapes(filename string) (shapes []PolyData) {
	file, err := shp.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	var field_names []string
	for _, shp_field := range file.Fields() {
		field_names = append(field_names, shp_field.String())
	}

	for file.Next() {
		index, shape := file.Shape()
		poly, ok := shape.(*shp.Polygon)
		if !ok {
			continue
		}
		var field_values []string
		for i := 0; i < len(field_names); i++ {
			field_values = append(field_values, file.ReadAttribute(index, i))
		}
		shapes = append(shapes, PolyData{*poly, index, field_names, field_values})
	}
	return shapes
}

func FindPoint(p *shp.Point, shapes *[]PolyData) (intersections []*PolyData) {
	completions := make(chan int)
	for _, shape := range *shapes {

		go func(shape PolyData, completions chan int, p *shp.Point) {
			if shape.poly.IncludesPoint(*p) {
				intersections = append(intersections, &shape)
			}

			completions <- 1
		}(shape, completions, p)
	}
	for i := 0; i < len(*shapes); i++ {
		<-completions
	}
	return intersections
}
