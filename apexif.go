package apexif

import (
	"github.com/abrander/apexif/fileformats"

	"github.com/abrander/apexif/fileformats/cr2"
	"github.com/abrander/apexif/fileformats/crw"
	"github.com/abrander/apexif/fileformats/heic"
	"github.com/abrander/apexif/fileformats/jpeg"
	"github.com/abrander/apexif/fileformats/png"
	"github.com/abrander/apexif/fileformats/tif"
	"github.com/abrander/apexif/fileformats/webp"
)

// Identify tries to identify the file format of the given data.
// If the file format is recognized, a FileType is returned.
// If the file format is not recognized, ErrImageNotRecognized
// is returned.
func Identify(data []byte) (fileformats.FileType, error) {
	try := []fileformats.Identifier{
		jpeg.Identify,
		png.Identify,
		heic.Identify,
		webp.Identify,
		cr2.Identify,
		crw.Identify,
		tif.Identify,
	}

	for _, recognizer := range try {
		fileType, err := recognizer(data)
		if err == nil {
			return fileType, nil
		}

		if err != fileformats.ErrImageNotRecognized {
			return nil, err
		}
	}

	return nil, fileformats.ErrImageNotRecognized
}
