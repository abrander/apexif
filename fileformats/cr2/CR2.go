package cr2

import (
	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type CR2 struct {
	bytes []byte
}

const signature = "II*\x00\x10\x00"

var _ fileformats.FileType = &CR2{}

func Identify(buf []byte) (fileformats.FileType, error) {
	if len(buf) < 10*1024 || string(buf[0:6]) != signature {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &CR2{
		bytes: buf,
	}, nil
}

func (c *CR2) Name() string {
	return "CR2"
}

func (c *CR2) MediaType() string {
	return "image/x-canon-cr2"
}

func (c *CR2) Exif() (*exif.Exif, error) {
	return exif.Parse(c.bytes)
}
