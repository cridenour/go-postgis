package postgis

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestPoint(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := Point{-84.5014, 39.1064}
	var newPoint Point

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("Point did not return itself through PostGIS.")
	}
}

func TestPointZ(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointZ{-84.5014, 39.1064, 167.9448}
	var newPoint PointZ

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointZ did not return itself through PostGIS.")
	}
}

func TestPointM(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointM{-84.5014, 39.1064, 1.0}
	var newPoint PointM

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointM did not return itself through PostGIS.")
	}
}

func TestPointZM(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointZM{-84.5014, 39.1064, 167.9448, 1.0}
	var newPoint PointZM

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointZM did not return itself through PostGIS.")
	}
}

func TestPointS(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointS{4326, -84.5014, 39.1064}
	var newPoint PointS

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointS did not return itself through PostGIS.")
	}
}

func TestPointZS(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointZS{4326, -84.5014, 39.1064, 167.9448}
	var newPoint PointZS

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointZS did not return itself through PostGIS.")
	}
}

func TestPointMS(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointMS{4326, -84.5014, 39.1064, 1.0}
	var newPoint PointMS

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointMS did not return itself through PostGIS.")
	}
}

func TestPointZMS(t *testing.T) {
	db := openTestConn(t)
	defer db.Close()

	point := PointZMS{4326, -84.5014, 39.1064, 167.9448, 1.0}
	var newPoint PointZMS

	if ok, err := compareGeometry(db, &point, &newPoint); !ok || err != nil {
		t.Error("PointZMS did not return itself through PostGIS.")
	}
}
