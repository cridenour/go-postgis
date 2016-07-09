package postgis

import (
	"testing"
	"encoding/hex"
	"strings"
)

var wkbPoint = "0101000020E6100000629B56DC3BC9E23F3626D20395B3DB3F"
var wkbLineString = "0102000020E610000007000000F00576DE1CF4C3BFB04AF07E8F6AAF3F187E3526D203D5BF3C9CB02089C7CE3F10DBB4E2DE49D6BF00D0F5C7CB6074BFCC8DB69667E4D2BF20C7711CC771BCBF705A71EF244BCDBF0838349FEB8FC7BFC090F10576DEACBF60CB33724969C8BF00A9F62110AEA23FD03037DA5A9EC1BF"
var wkbPolygon = "0103000020E610000001000000080000000000003CA08F5B40907BF23F88393C40FFFFFFCB978F5B40C77BF23FD4383C40FFFFFF3BCD8F5B40CC7BF29FCE383C40010000C4EA8F5B40887BF2BF9E393C40FEFFFF63F08F5B40697BF29F093A3C40FFFFFF2BEC8F5B405D7BF25F2B3A3C40FFFFFF13D78F5B405D7BF25F2B3A3C400000003CA08F5B40907BF23F88393C40"

func TestWkb(t *testing.T) {
	point := Point{}
	lineString := LineString{}
	polygon := Polygon{}
	var b []byte
	var ok bool

	//scane
	if err := point.Scan([]byte(wkbPoint)); err != nil {
		t.Error(err)
	}
	//value
	p, err := point.Value()
	if err != nil {
		t.Error(err)
	}
	if b, ok = p.([]byte); !ok {
		t.Error("cann't assigment Point value to []byte")
	}
	if hex.EncodeToString(b) != strings.ToLower(wkbPoint) {
		t.Error("read point value error")
	}

	//scane
	if err := lineString.Scan([]byte(wkbLineString)); err != nil {
		t.Error(err)
	}
	//value
	l, err := lineString.Value()
	if err != nil {
		t.Error(err)
	}
	if b, ok = l.([]byte); !ok {
		t.Error("cann't assigment LineString value to []byte")
	}
	if hex.EncodeToString(b) != strings.ToLower(wkbLineString) {
		t.Error("read lineString value error")
	}
	//scane
	if err := polygon.Scan([]byte(wkbPolygon)); err != nil {
		t.Error(err)
	}
	//value
	g, err := polygon.Value()
	if err != nil {
		t.Error(err)
	}

	if b, ok = g.([]byte); !ok {
		t.Error("cann't assigment Polygon value to []byte")
	}
	if hex.EncodeToString(b) != strings.ToLower(wkbPolygon) {
		t.Error("read Polygon value error")
	}
}
