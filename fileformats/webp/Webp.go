package webp

import (
	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/containers/riff"
	"github.com/abrander/apexif/fileformats"
)

type Webp struct {
	bytes []byte
}

func Identify(data []byte) (fileformats.FileType, error) {
	if len(data) < 12 {
		return nil, fileformats.ErrImageNotRecognized
	}

	if string(data[:4]) != "RIFF" {
		return nil, fileformats.ErrImageNotRecognized
	}

	if string(data[8:12]) != "WEBP" {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &Webp{
		bytes: data,
	}, nil
}

func (w *Webp) Name() string {
	return "WebP"
}

func (w *Webp) MediaType() string {
	return "image/webp"
}

func (w *Webp) Exif() (*exif.Exif, error) {
	r, err := riff.Parse(w.bytes)
	if err != nil {
		return nil, err
	}

	if r.Riff.Length < 10 {
		return nil, fileformats.ErrImageNotRecognized
	}

	if string(r.Riff.Data[:4]) != "WEBP" {
		return nil, fileformats.ErrImageNotRecognized
	}

	chunks, _ := riff.ReadChunks(r.Riff.Data[4:])
	for _, chunk := range chunks {
		if chunk.Identifier == "EXIF" && chunk.Length > 6 {
			return exif.Parse(chunk.Data[6:])
		}
	}

	return nil, exif.ErrNoExifFound
}
