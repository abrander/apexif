package tiff

import (
	"fmt"
)

type SignedRational struct {
	Numerator   int32
	Denominator int32
}

func (r SignedRational) Float() float64 {
	return float64(r.Numerator) / float64(r.Denominator)
}

func (r SignedRational) String() string {
	return fmt.Sprintf("%d/%d", r.Numerator, r.Denominator)
}
