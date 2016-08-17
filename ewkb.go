//ewkb format convert:https://github.com/mc2soft/pq-types/blob/master/postgis.go
//ogc wkb format convert:http://www.cnblogs.com/marsprj/archive/2013/02/08/2909452.html
package postgis

type ewkbHeader struct {
	ByteOrder byte
	WkbType   [4]byte
	SRID      uint32
}

type ewkbPoint struct {
	Header ewkbHeader
	Point  Point
}

type ewkbLineString struct {
	Header     ewkbHeader
	Count      uint32
	LineString LineString
}

type ewkbPolygon struct {
	Header  ewkbHeader
	Rings   uint32
	Count   uint32
	Polygon Polygon
}