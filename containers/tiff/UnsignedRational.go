package tiff

import (
	"fmt"
)

// UnsignedRational is a type representing an unsigned rational
// number in a tag.
type UnsignedRational struct {
	Numerator   uint32
	Denominator uint32
}

var _ fmt.Stringer = UnsignedRational{}

// Float returns the float64 value of the rational.
func (r UnsignedRational) Float() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

// String returns a string representation of the rational.
func (r UnsignedRational) String() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}
