package png

import (
	"encoding/binary"
	"errors"

	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type PNG struct {
	bytes []byte
}

const signature string = "\x89PNG\r\n\x1a\n"
const crcSize = 4

var _ fileformats.FileType = &PNG{}

func Identify(data []byte) (fileformats.FileType, error) {
	if len(data) < 2*len(signature) || string(data[:len(signature)]) != signature {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &PNG{
		bytes: data,
	}, nil
}

func (p *PNG) Name() string {
	return "PNG"
}

func (p *PNG) MediaType() string {
	return "image/png"
}

func (p *PNG) Exif() (*exif.Exif, error) {
	offset := len(signature)

	for {
		if offset >= len(p.bytes) {
			return nil, errors.New("no EXIF data found")
		}

		if offset+8 > len(p.bytes) {
			return nil, fileformats.ErrImageNotRecognized
		}

		length := binary.BigEndian.Uint32(p.bytes[offset:])
		chunkType := string(p.bytes[offset+4 : offset+8])

		offset += 8

		switch chunkType {
		case "eXIf":
			return exif.Parse(p.bytes[offset : offset+int(length)])

		default:
			offset += int(length) + crcSize
		}
	}
}
