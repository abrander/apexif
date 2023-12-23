package tiff

import (
	"fmt"
)

// SignedRational is a type representing a signed rational number in a tag.
type SignedRational struct {
	Numerator   int32
	Denominator int32
}

var _ fmt.Stringer = SignedRational{}

// Float returns the float64 value of the rational.
func (r SignedRational) Float() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

// String returns the string representation of the rational.
func (r SignedRational) String() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}
