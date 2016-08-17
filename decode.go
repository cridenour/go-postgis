package postgis

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
)

const (
	wkbXDR byte = 0
	wkbNDR byte = 1
)

// Helper functions to read the EWKB (well-known binary) from PostGIS
// Format document available at http://trac.osgeo.org/postgis/browser/trunk/doc/ZMSgeoms.txt

// Since Postgres by default returns hex encoded strings we need to first get bytes back, only used from Scan
func decode(value interface{}) (io.Reader, error) {
	ewkb, err := hex.DecodeString(string(value.([]byte)))
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(ewkb), nil
}

func readEWKB(reader io.Reader, g interface{}) error {
	var byteOrder binary.ByteOrder
	var wkbByteOrder byte
	var wkbType uint32
	var srid uint32

	// Read as Little Endian to attempt to determine byte order
	if err := binary.Read(reader, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	// Decide byte order
	switch wkbByteOrder {
	case wkbXDR:
		byteOrder = binary.BigEndian
	case wkbNDR:
		byteOrder = binary.LittleEndian
	default:
		return errors.New("Unsupported byte order")
	}

	// Determine the geometry type
	if err := binary.Read(reader, byteOrder, &wkbType); err != nil {
		return err
	}
	//Determine the srid
	if err := binary.Read(reader, byteOrder, &srid); err != nil {
		return err
	}
	//diffence geometry type
	switch wkbType {
	case GeomPoint:
		// Decode into point struct
		return binary.Read(reader, byteOrder, g)

	case GeomLineString:
		//Detemine the point count
		var count uint32
		if err := binary.Read(reader, byteOrder, &count); err != nil {
			return err
		}

		// Decode into linestring struct
		g, ok := g.(*LineString)
		if !ok {
			return errors.New("geometry type should be LineString")
		}
		g.Points = make([]Point, count)
		if err := binary.Read(reader, byteOrder, g.Points); err != nil {
			return err
		}

		return nil
	case GeomPolygon:
		//Detemine the LinearRing count of the polygon
		var rings uint32
		if err := binary.Read(reader, byteOrder, &rings); err != nil {
			return err
		}
		//Detemine the count of the polygon
		var count uint32
		if err := binary.Read(reader, byteOrder, &count); err != nil {
			return err
		}
		g, ok := g.(*Polygon)
		if !ok {
			return errors.New("geometry type should be Polygon")
		}

		// Decode into Polygon struct
		g.Points = make([]Point, count)
		if err := binary.Read(reader, byteOrder, g.Points); err != nil {
			return err
		}

		return nil
	default:
		return errors.New("Unsupported Geomtry type")
	}
}
