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

func readEWKB(reader io.Reader, g Geometry) error {
	var byteOrder binary.ByteOrder
	var wkbByteOrder byte
	var wkbType uint32

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

	// Determine the geometery type
	if err := binary.Read(reader, byteOrder, &wkbType); err != nil {
		return err
	}

	// Decode into our struct
	return binary.Read(reader, byteOrder, g)
}
