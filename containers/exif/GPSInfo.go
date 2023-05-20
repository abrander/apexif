package exif

import (
	"github.com/abrander/apexif/containers/tiff"
)

type GPSInfo struct {
	tiff.IFD

	e *Exif
}

func (e *Exif) GPSInfo() (*GPSInfo, error) {
	entry, err := e.Tiff.Entry(AnyIFD, tiff.GPSInfoIFDPointer)
	if err != nil {
		return nil, err
	}

	ifd, err := e.Tiff.ReadIFD(entry.Offset())
	if err != nil {
		return nil, err
	}

	return &GPSInfo{ifd, e}, nil
}

func (i *GPSInfo) Date() (string, error) {
	entry, err := i.Entry(tiff.Tag(GPSDateStamp))
	if err != nil {
		return "", err
	}

	str, err := entry.Ascii()
	if err != nil {
		return "", err
	}

	return str, nil
}
