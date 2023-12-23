package tiff

import (
	"errors"
	"fmt"
	"strings"
)

// Entry is a type representing an IFD entry.
type Entry struct {
	Tag         Tag
	Type        Type
	Count       uint32
	ValueOffset [4]byte

	tiff *Tiff
}

var _ fmt.Stringer = &Entry{}

// String returns a string representation of the entry.
func (e Entry) String() string {
	var str string
	var err error

	switch {
	case e.Type == Ascii:
		str, err = e.Ascii()
		str = "Ascii:[" + str + "]"

	case e.Type == Short && e.Count == 1:
		var sh uint16
		sh, err = e.Short()
		if err == nil {
			str = fmt.Sprintf("Short:[%d]", sh)
		}

	case e.Type == Short:
		var sh []uint16
		sh, err = e.ShortSlice()
		if err == nil {
			str = fmt.Sprintf("Short(%d):[%v]", len(sh), sh)
		}

	case e.Type == Long && e.Count == 1:
		var l uint32
		l, err = e.Long()
		if err == nil {
			str = fmt.Sprintf("Long:[%d]", l)
		}

	case e.Type == Long:
		var l []uint32
		l, err = e.LongSlice()
		if err == nil {
			str = fmt.Sprintf("Long(%d):[%v]", len(l), l)
		}

	case e.Type == Rational && e.Count == 1:
		var r UnsignedRational
		r, err = e.Rational()
		if err == nil {
			str = "Rational:[" + r.String() + "]"
		}

	case e.Type == Rational:
		var r []UnsignedRational
		r, err = e.RationalSlice()
		if err == nil {
			str = fmt.Sprintf("Rational(%d):[%v]", len(r), r)
		}

	case e.Type == SRational && e.Count == 1:
		var r SignedRational
		r, err = e.SRational()
		if err == nil {
			str = fmt.Sprintf("SRational:[%s]", r.String())
		}

	case e.Type == SRational:
		var r []SignedRational
		r, err = e.SRationalSlice()
		if err == nil {
			str = fmt.Sprintf("SRational(%d):[%v]", len(r), r)
		}

	default:
		str = fmt.Sprintf("ERR %04x %s(%d): %d/%08x", int(e.Tag), e.Type, e.Count, e.Offset(), e.ValueOffset)
	}

	if err != nil {
		str = " ERR:" + err.Error()
	}

	return fmt.Sprintf("0x%04x ", int(e.Tag)) + str
}

// Value returns the value of the entry. This *can* panic!
func (e Entry) Value() (any, error) {
	switch e.Type {
	case Ascii:
		return e.Ascii()

	case Short:
		if e.Count == 1 {
			return e.Short()
		}
		return e.ShortSlice()

	case Long:
		if e.Count == 1 {
			return e.Long()
		}
		return e.LongSlice()

	case Rational:
		if e.Count == 1 {
			return e.Rational()
		}

		panic("Not implemented")

	default:
		panic("Not implemented")
	}
}

// Offset returns the offset of the entry.
func (e Entry) Offset() int {
	return int(e.tiff.endianness.Uint32(e.ValueOffset[:]))
}

// Ascii returns the ASCII value of the entry.
func (e *Entry) Ascii() (string, error) {
	clean := func(str string) string {
		str = strings.TrimSuffix(str, "\x00")
		str = strings.TrimSpace(str)

		return str
	}

	if e.Type != Ascii {
		return "", errors.New("Not ascii")
	}

	if e.Count <= 4 {
		// String is stored in the value offset
		return clean(string(e.ValueOffset[:e.Count])), nil
	}

	if len(e.tiff.bytes) < e.Offset()+int(e.Count) {
		return "", errors.New("Buffer too small")
	}

	str := string(e.tiff.bytes[e.Offset() : e.Offset()+int(e.Count)])

	return clean(str), nil
}

// Short returns the unsigned short value of the entry.
func (e *Entry) Short() (uint16, error) {
	if e.Type != Short {
		return 0, errors.New("not a short")
	}

	if e.Count != 1 {
		return 0, errors.New("not a single short")
	}

	return e.tiff.endianness.Uint16(e.ValueOffset[:]), nil
}

// ShortSlice returns the unsigned short slice value of the entry.
func (e *Entry) ShortSlice() ([]uint16, error) {
	if e.Type != Short {
		return nil, errors.New("not a short")
	}

	if e.Count <= 2 {
		// Short is stored in the value offset
		shorts := make([]uint16, e.Count)

		for i := 0; i < int(e.Count); i++ {
			shorts[i] = e.tiff.endianness.Uint16(e.ValueOffset[i*2:])
		}

		return shorts, nil
	}

	if len(e.tiff.bytes) < e.Offset()+int(e.Count)*2 {
		return nil, errors.New("buffer too small")
	}

	shorts := make([]uint16, e.Count)

	for i := 0; i < int(e.Count); i++ {
		shorts[i] = e.tiff.endianness.Uint16(e.tiff.bytes[e.Offset()+2*i:])
	}

	return shorts, nil
}

// Long returns the unsigned long (32 bit) value of the entry.
func (e *Entry) Long() (uint32, error) {
	if e.Type != Long {
		return 0, errors.New("not a long")
	}

	if e.Count != 1 {
		return 0, errors.New("not a single long")
	}

	if len(e.tiff.bytes) < e.Offset()+4 {
		return 0, errors.New("buffer too small")
	}

	return e.tiff.endianness.Uint32(e.ValueOffset[:]), nil
}

// LongSlice returns the unsigned long slice value of the entry.
func (e *Entry) LongSlice() ([]uint32, error) {
	if e.Type != Long {
		return nil, errors.New("not a long")
	}

	if e.Count <= 1 {
		// Long is stored in the value offset
		longs := make([]uint32, e.Count)

		for i := 0; i < int(e.Count); i++ {
			longs[i] = e.tiff.endianness.Uint32(e.ValueOffset[i*4:])
		}

		return longs, nil
	}

	if len(e.tiff.bytes) < e.Offset()+int(e.Count)*4 {
		return nil, errors.New("buffer too small")
	}

	longs := make([]uint32, e.Count)

	for i := 0; i < int(e.Count); i++ {
		longs[i] = e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+4*i:])
	}

	return longs, nil
}

// Rational returns the unsigned rational value of the entry.
func (e *Entry) Rational() (UnsignedRational, error) {
	if e.Type != Rational {
		return UnsignedRational{}, errors.New("not a rational")
	}

	if e.Count != 1 {
		return UnsignedRational{}, errors.New("not a single rational")
	}

	if len(e.tiff.bytes) < e.Offset()+8 {
		return UnsignedRational{}, errors.New("buffer too small")
	}

	return UnsignedRational{
		Numerator:   e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset():]),
		Denominator: e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+4:]),
	}, nil
}

// RationalSlice returns the unsigned rational slice value of the entry.
func (e *Entry) RationalSlice() ([]UnsignedRational, error) {
	if e.Type != Rational {
		return nil, errors.New("not a rational")
	}

	if e.Count <= 1 {
		// Rational is stored in the value offset
		ratios := make([]UnsignedRational, e.Count)

		for i := 0; i < int(e.Count); i++ {
			ratios[i] = UnsignedRational{
				Numerator:   e.tiff.endianness.Uint32(e.ValueOffset[i*8:]),
				Denominator: e.tiff.endianness.Uint32(e.ValueOffset[i*8+4:]),
			}
		}

		return ratios, nil
	}

	if len(e.tiff.bytes) < e.Offset()+int(e.Count)*8 {
		return nil, errors.New("buffer too small")
	}

	ratios := make([]UnsignedRational, e.Count)

	for i := 0; i < int(e.Count); i++ {
		ratios[i] = UnsignedRational{
			Numerator:   e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+8*i:]),
			Denominator: e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+8*i+4:]),
		}
	}

	return ratios, nil
}

// SRational returns the signed rational value of the entry.
func (e *Entry) SRational() (SignedRational, error) {
	if e.Type != SRational {
		return SignedRational{}, errors.New("not a rational")
	}

	if e.Count != 1 {
		return SignedRational{}, errors.New("not a single rational")
	}

	if len(e.tiff.bytes) < e.Offset()+8 {
		return SignedRational{}, errors.New("buffer too small")
	}

	return SignedRational{
		Numerator:   int32(e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset():])),
		Denominator: int32(e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+4:])),
	}, nil
}

// SRationalSlice returns the signed rational slice value of the entry.
func (e *Entry) SRationalSlice() ([]SignedRational, error) {
	if e.Type != SRational {
		return nil, errors.New("not a rational")
	}

	if e.Count <= 1 {
		// Rational is stored in the value offset
		ratios := make([]SignedRational, e.Count)

		for i := 0; i < int(e.Count); i++ {
			ratios[i] = SignedRational{
				Numerator:   int32(e.tiff.endianness.Uint32(e.ValueOffset[i*8:])),
				Denominator: int32(e.tiff.endianness.Uint32(e.ValueOffset[i*8+4:])),
			}
		}

		return ratios, nil
	}

	if len(e.tiff.bytes) < e.Offset()+int(e.Count)*8 {
		return nil, errors.New("buffer too small")
	}

	ratios := make([]SignedRational, e.Count)

	for i := 0; i < int(e.Count); i++ {
		ratios[i] = SignedRational{
			Numerator:   int32(e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+8*i:])),
			Denominator: int32(e.tiff.endianness.Uint32(e.tiff.bytes[e.Offset()+8*i+4:])),
		}
	}

	return ratios, nil
}
