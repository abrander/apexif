package tif

import (
	"encoding/binary"

	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type Tif struct {
	bytes []byte
}

func Identify(data []byte) (fileformats.FileType, error) {
	if len(data) < 8 {
		return nil, fileformats.ErrImageNotRecognized
	}

	var endianness binary.ByteOrder

	switch {
	case data[0] == 'I' && data[1] == 'I':
		endianness = binary.LittleEndian

	case data[0] == 'M' && data[1] == 'M':
		endianness = binary.BigEndian

	default:
		return nil, fileformats.ErrImageNotRecognized
	}

	magic := endianness.Uint16(data[2:])
	if magic != 42 {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &Tif{
		bytes: data,
	}, nil
}

func (t *Tif) Name() string {
	return "TIFF"
}

func (t *Tif) MediaType() string {
	return "image/tiff"
}

func (t *Tif) Exif() (*exif.Exif, error) {
	return exif.Parse(t.bytes)
}
