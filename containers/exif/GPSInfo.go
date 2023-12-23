package exif

import (
	"github.com/abrander/apexif/containers/tiff"
)

// GPSInfo is a type representing GPS information as saved in
// the EXIF data.
type GPSInfo struct {
	tiff.IFD

	e *Exif
}

// GPSInfo returns the GPSInfo IFD or an error.
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

// Date returns the GPS date as a string or an error.
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
