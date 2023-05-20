package tiff

import (
	"fmt"
)

type UnsignedRational struct {
	Numerator   uint32
	Denominator uint32
}

func (r UnsignedRational) Float() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r UnsignedRational) String() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}
