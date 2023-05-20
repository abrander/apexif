package tiff

import (
	"encoding/binary"
	"errors"
)

type Tiff struct {
	bytes []byte

	endianness binary.ByteOrder
	ifds       []IFD
}

var (
	ErrNotTiff = errors.New("not a TIFF file")

	ErrTagNotFound = errors.New("tag not found")

	ErrIFDNotFound = errors.New("IFD not found")
)

func Parse(data []byte) (*Tiff, error) {
	if len(data) < 8 {
		return nil, ErrNotTiff
	}

	var endianness binary.ByteOrder

	switch {
	case data[0] == 'I' && data[1] == 'I':
		endianness = binary.LittleEndian

	case data[0] == 'M' && data[1] == 'M':
		endianness = binary.BigEndian

	default:
		return nil, ErrNotTiff
	}

	magic := endianness.Uint16(data[2:])
	if magic != 42 {
		return nil, ErrNotTiff
	}

	t := &Tiff{
		bytes:      data,
		endianness: endianness,
	}

	// Get the offset to the first IFD
	ifdOffset := int(endianness.Uint32(data[4:]))

	// Read the IFDs
	t.ifds = []IFD{}
	for {
		ifd, err := t.ReadIFD(ifdOffset)
		if err != nil {
			return nil, err
		}
		ifdCount := int(endianness.Uint16(data[ifdOffset:]))

		t.ifds = append(t.ifds, ifd)

		// Check if there is another IFD
		if len(data) < int(ifdOffset)+2 {
			return nil, errors.New("Buffer too small")
		}

		ifdOffset = int(endianness.Uint32(data[ifdOffset+2+ifdCount*12:]))
		if ifdOffset == 0 {
			break
		}
	}

	return t, nil
}

func (t *Tiff) IFDs() []IFD {
	return t.ifds
}

const AnyIFD = -1

func (t *Tiff) Entry(ifd int, tag Tag) (Entry, error) {
	var zero Entry

	if ifd == AnyIFD {
		for _, ifd := range t.ifds {
			e, err := ifd.Entry(tag)
			if err == nil {
				return e, nil
			}
		}

		return zero, ErrTagNotFound
	}

	if ifd < 0 || ifd >= len(t.ifds) {
		return zero, ErrIFDNotFound
	}

	return t.ifds[ifd].Entry(tag)
}

func (t *Tiff) ReadIFD(offset int) (IFD, error) {
	buf := t.bytes[offset:]

	if len(buf) < 2 {
		return nil, errors.New("buffer too small")
	}

	// Get number of IFDs
	ifdCount := t.endianness.Uint16(buf[0:2])

	if len(t.bytes) < int(ifdCount)*12+2 {
		return nil, errors.New("buffer too small")
	}

	// Read the IFDs
	ifds := make(IFD, ifdCount)

	for i := 0; i < int(ifdCount); i++ {
		ifd := Entry{
			Tag:         Tag(t.endianness.Uint16(buf[i*12+2 : i*12+4])),
			Type:        Type(t.endianness.Uint16(buf[i*12+4 : i*12+6])),
			Count:       t.endianness.Uint32(buf[i*12+6 : i*12+10]),
			ValueOffset: *(*[4]byte)(buf[i*12+10 : i*12+14]),

			tiff: t,
		}

		ifds[i] = ifd
	}

	return ifds, nil
}

func (t *Tiff) Ascii(ifd int, tag Tag) (string, error) {
	e, err := t.Entry(ifd, tag)
	if err != nil {
		return "", err
	}

	return e.Ascii()
}
