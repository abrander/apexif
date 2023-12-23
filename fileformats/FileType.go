package fileformats

import (
	"errors"

	"github.com/abrander/apexif/containers/exif"
)

// FileType is the interface that all file formats must implement.
type FileType interface {
	// Exif returns the EXIF data from the file if found. If not
	// found, nil and ErrNoExifFound is returned.
	Exif() (*exif.Exif, error)

	// Name returns the name of the file format.
	Name() string

	// MediaType returns the MIME type of the file format.
	MediaType() string
}

// Identifier is a function type that can identify a file format. It
// returns the FileType if the file format is recognized, or
// ErrImageNotRecognized if not.
// This should be implemented by all file formats.
type Identifier func(data []byte) (FileType, error)

// ErrImageNotRecognized is returned by the Identify function if
// the file format is not recognized.
var ErrImageNotRecognized = errors.New("image file format not recognized")
