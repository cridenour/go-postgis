package postgis

import (
	"database/sql"
	"os"
	"reflect"
)

type Fatalistic interface {
	Fatal(args ...interface{})
}

func openTestConn(t Fatalistic) *sql.DB {
	datname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("PGSSLMODE")

	if datname == "" {
		os.Setenv("PGDATABASE", "pqgotest")
	}

	if sslmode == "" {
		os.Setenv("PGSSLMODE", "disable")
	}

	conn, err := sql.Open("postgres", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("CREATE EXTENSION IF NOT EXISTS postgis")
	if err != nil {
		t.Fatal("PostGIS extension create failed.")
	}

	return conn
}

func compareGeometry(db *sql.DB, g1 Geometry, g2 Geometry) (bool, error) {
	if err := db.QueryRow("SELECT GeomFromEWKB($1);", g1).Scan(g2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(g1, g2), nil
}
