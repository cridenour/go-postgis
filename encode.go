package postgis

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

func encode(buffer *bytes.Buffer) string {
	return hex.EncodeToString(buffer.Bytes())
}

func writeEWKB(g Geometry) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)

	// Set our endianness
	if err := binary.Write(buffer, binary.LittleEndian, wkbNDR); err != nil {
		return nil, err
	}

	if err := binary.Write(buffer, binary.LittleEndian, g.GetType()); err != nil {
		return nil, err
	}

	if err := g.Write(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}
