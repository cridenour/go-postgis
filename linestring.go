package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
)

func (l *LineString)Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, l); err != nil {
		return err
	}

	return nil
}
func (l LineString)Value() (driver.Value, error) {
	buffer, err := writeEWKB(&l)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
func (l *LineString)Count() uint32 {
	return uint32(len(l.Points))
}

func (l *LineString)Write(buffer *bytes.Buffer) error {
	//write count
	if err := binary.Write(buffer, binary.LittleEndian, l.Count()); err != nil {
		return err
	}
	for _, point := range l.Points {
		if err := binary.Write(buffer, binary.LittleEndian, point); err != nil {
			return err
		}
	}
	return nil
}

func (l LineString)GetType() uint32 {
	return GeomLineString
}