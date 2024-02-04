package geopatials

import (
	"fmt"
	"testing"
)

var dbname = "gis"
var collname = "gis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			107.90077136603423, -7.216136075981481,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{107.90241395501647, -7.2159039954805735},
				{107.90241973415237, -7.216093196465536},
				{107.90275492403225, -7.216093196465536},
				{107.90275492403225, -7.2159039954805735},
				{107.90241395501647, -7.2159039954805735},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gis", "gis")
	coordinates := Point{
		Coordinates: []float64{
			107.90230814756461, -7.212792294268141,
		},
	}
	datagedung := Near(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Point{
		Coordinates: []float64{
			107.90357575285924, -7.213350872053013,
		},
	}
	datagedung := NearSphere(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.90212716477299, -7.213924382944427},
			{107.90395765245472, -7.213931661492197},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
