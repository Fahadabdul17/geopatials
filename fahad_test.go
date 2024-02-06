package geopatials

import (
	"fmt"
	"testing"
)

var dbname = "GismongoDB"
var collname = "Postgis"

func TestGeoIntersects(t *testing.T) {
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
	mconn := SetConnection2dsphere("MONGOSTRING", "GismongoDB", "Postgis")
	coordinates := Point{
		Coordinates: []float64{
			107.90230814756461, -7.212792294268141,
		},
	}
	datagedung := Near(mconn, "Postgis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GismongoDB")
	coordinates := Point{
		Coordinates: []float64{
			107.90357575285924, -7.213350872053013,
		},
	}
	datagedung := NearSphere(mconn, "Postgis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "GismongoDB")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{107.90212716477299, -7.213924382944427},
			{107.90395765245472, -7.213931661492197},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
