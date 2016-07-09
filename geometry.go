package postgis

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
)

//geo type
//postgis return hex
const (
	GeomPoint uint32 = 0x20000001
	GeomLineString uint32 = 0x20000002
	GeomPolygon uint32 = 0x20000003
	GeomMultiPoint uint32 = 0x20000004
	GeomMultiLineString uint32 = 0x20000005
	GeomMultiPolygon uint32 = 0x20000006
	GeomCollection uint32 = 0x20000007
)

type Geometry interface {
	sql.Scanner
	driver.Valuer
	GetType() uint32
	Write(*bytes.Buffer) error
}

type LineString struct {
	Points []Point
}
type LinearRing struct {
	Lines []LineString
}

type Polygon struct {
	Points []Point
}
type MultiPoint struct {
	Points []Point
}

type MultiLineString []LineString
type MultiPolygon []Polygon
type GeometryCollection []Geometry
