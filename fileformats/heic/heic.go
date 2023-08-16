package heic

import (
	"github.com/abrander/apexif/containers/bmff"
	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type HEIC struct {
	bytes []byte
}

var _ fileformats.FileType = &HEIC{}

func Identify(data []byte) (fileformats.FileType, error) {
	// Check that the file looks like a HEIC file.
	if len(data) < 12 {
		return nil, fileformats.ErrImageNotRecognized
	}

	if string(data[4:12]) != "ftypheic" {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &HEIC{
		bytes: data,
	}, nil
}

func (h *HEIC) Name() string {
	return "HEIC"
}

func (h *HEIC) MediaType() string {
	return "image/heic"
}

func (h *HEIC) Exif() (*exif.Exif, error) {
	b, err := bmff.Parse(h.bytes)
	if err != nil {
		return nil, err
	}

	exifBytes := b.Iloc("Exif")

	return exif.Parse(exifBytes)
}
