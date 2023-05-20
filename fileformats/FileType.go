package fileformats

import (
	"errors"

	"github.com/abrander/apexif/containers/exif"
)

type FileType interface {
	Exif() (*exif.Exif, error)
	Name() string
	MediaType() string
}

type Identifier func(data []byte) (FileType, error)

var ErrImageNotRecognized = errors.New("image file format not recognized")
