package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
)

func (p *Polygon)Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}
//point count
func (p *Polygon)Count() uint32 {
	return uint32(len(p.Points))
}
//default count is 1
func (p *Polygon)LinearRingCount() uint32 {
	return 1
}
func (p *Polygon)Value() (driver.Value, error) {
	buffer, err := writeEWKB(p)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (p *Polygon)Write(buffer *bytes.Buffer) error {
	//write count
	if err := binary.Write(buffer, binary.LittleEndian, p.LinearRingCount()); err != nil {
		return err
	}
	if err := binary.Write(buffer, binary.LittleEndian, p.Count()); err != nil {
		return err
	}
	for _, point := range p.Points {
		if err := binary.Write(buffer, binary.LittleEndian, point); err != nil {
			return err
		}
	}
	return nil
}

func (p *Polygon)GetType() uint32 {
	return GeomPolygon
}