package tiff

import (
	"fmt"
)

// Type is a type representing the type of an IFD entry.
type Type uint16

const (
	Byte      Type = 1
	Ascii     Type = 2
	Short     Type = 3
	Long      Type = 4
	Rational  Type = 5
	Undefined Type = 7
	SLong     Type = 9
	SRational Type = 10
)

var _ fmt.Stringer = Type(0)

// String returns a string representation of the type.
func (t Type) String() string {
	switch t {
	case Byte:
		return "Byte"
	case Ascii:
		return "Ascii"
	case Short:
		return "Short"
	case Long:
		return "Long"
	case Rational:
		return "Rational"
	case Undefined:
		return "Undefined"
	case SLong:
		return "SLong"
	case SRational:
		return "SRational"
	}

	return "Unknown"
}
