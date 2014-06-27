package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
)

// Structs respresenting varying types of points
type Point struct {
	X, Y float64
}

type PointZ struct {
	X, Y, Z float64
}

type PointM struct {
	X, Y, M float64
}

type PointZM struct {
	X, Y, Z, M float64
}

type PointS struct {
	SRID int32
	X, Y float64
}

type PointZS struct {
	SRID    int32
	X, Y, Z float64
}

type PointMS struct {
	SRID    int32
	X, Y, M float64
}

type PointZMS struct {
	SRID       int32
	X, Y, Z, M float64
}

/** Point functions **/
func (p *Point) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p Point) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p Point) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p Point) GetType() uint32 {
	return 1
}

/** PointZ functions **/
func (p *PointZ) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZ) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZ) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZ) GetType() uint32 {
	return 0x80000001
}

/** PointM functions **/
func (p *PointM) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointM) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointM) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointM) GetType() uint32 {
	return 0x40000001
}

/** PointZM functions **/
func (p *PointZM) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZM) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZM) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZM) GetType() uint32 {
	return 0xC0000001
}

/** PointS functions **/
func (p *PointS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointS) GetType() uint32 {
	return 0x20000001
}

/** PointZS functions **/
func (p *PointZS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZS) GetType() uint32 {
	return 0xA0000001
}

/** PointMS functions **/
func (p *PointMS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointMS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointMS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointMS) GetType() uint32 {
	return 0x60000001
}

/** PointZMS functions **/
func (p *PointZMS) Scan(value interface{}) error {
	reader, err := decode(value)
	if err != nil {
		return err
	}

	if err = readEWKB(reader, p); err != nil {
		return err
	}

	return nil
}

func (p PointZMS) Value() (driver.Value, error) {
	buffer, err := writeEWKB(&p)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (p PointZMS) Write(buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.LittleEndian, &p)
	return err
}

func (p PointZMS) GetType() uint32 {
	return 0xE0000001
}
