package postgis

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
)

type Geometry interface {
	sql.Scanner
	driver.Valuer
	GetType() uint32
	Write(*bytes.Buffer) error
}
