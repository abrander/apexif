package jpeg

import (
	"encoding/binary"

	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type JPEG struct {
	bytes []byte
}

const (
	SOI  = 0xffd8 // Start of image
	APP1 = 0xffe1 // Exif (mostly)
)

var _ fileformats.FileType = &JPEG{}

func Identify(data []byte) (fileformats.FileType, error) {
	if len(data) < 16 || binary.BigEndian.Uint16(data) != SOI || data[2] != 0xff {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &JPEG{
		bytes: data,
	}, nil
}

func (j *JPEG) Name() string {
	return "JPEG"
}

func (j *JPEG) MediaType() string {
	return "image/jpeg"
}

func (j *JPEG) Exif() (*exif.Exif, error) {
	offset := 2

	for {
		if offset+4 >= len(j.bytes) {
			return nil, exif.ErrNoExifFound
		}

		marker := binary.BigEndian.Uint16(j.bytes[offset:])
		length := binary.BigEndian.Uint16(j.bytes[offset+2:])

		if length == 0 {
			return nil, exif.ErrNoExifFound
		}

		if marker == APP1 {
			if offset+4+int(length) > len(j.bytes) {
				return nil, exif.ErrNoExifFound
			}

			buf := j.bytes[offset+4:]

			if len(buf) >= 6 && string(buf[0:6]) == "Exif\000\000" {
				return exif.Parse(buf[6:])
			}
		}

		offset += int(length) + 2
	}
}
