package exif

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/abrander/apexif/containers/tiff"
)

// Exif is a type representing EXIF data from some source.
type Exif struct {
	tiff.Tiff

	exifIDFPointer *tiff.IFD
}

var (
	// ErrNoExifFound is returned if no EXIF data is found.
	ErrNoExifFound = errors.New("no EXIF data found")
)

// AnyIFD is a constant used to indicate that any IFD can be searched.
const AnyIFD = tiff.AnyIFD

// unread is a pointer used to indicate that the IFD has not been read yet.
var unread = &tiff.IFD{}

// Parse parses the given data as EXIF data or returns an error.
func Parse(data []byte) (*Exif, error) {
	t, err := tiff.Parse(data)
	if err != nil {
		return nil, err
	}

	return &Exif{*t, unread}, nil
}

// Entry returns the entry for the given IFD and tag or returns an
// error. AnyIDF can be used to search all IFDs.
func (e *Exif) Entry(ifd int, tag Tag) (tiff.Entry, error) {
	if ifd == AnyIFD {
		if e.exifIDFPointer == unread {
			entry, err := e.Tiff.Entry(AnyIFD, tiff.ExifIDFPointer)
			if err == nil {
				*e.exifIDFPointer, _ = e.Tiff.ReadIFD(entry.Offset())
			}
		}

		if *e.exifIDFPointer != nil {
			entry, err := e.exifIDFPointer.Entry(tiff.Tag(tag))
			if err == nil {
				return entry, nil
			}
		}
	}

	return e.Tiff.Entry(ifd, tiff.Tag(tag))
}

// Time returns the time for the given IFD and tag or returns
// an error. Since EXIF carries no timezone information, the
// location must be passed to set the timezone.
func (e *Exif) Time(ifd int, tag Tag, loc *time.Location) (time.Time, error) {
	entry, err := e.Entry(ifd, tag)
	if err != nil {
		return time.Time{}, err
	}

	str, err := entry.Ascii()
	if err != nil {
		return time.Time{}, err
	}

	return time.ParseInLocation("2006:01:02 15:04:05", str, loc)
}

func parseSubSecTime(str string) (time.Duration, error) {
	if len(str) == 0 {
		return 0, nil
	}

	if len(str) > 9 {
		str = str[:9]
	}

	n, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return time.Duration(n) * time.Microsecond, nil
}

// Time returns the "DateTimeOriginal" time or returns an error.
func (e *Exif) TimeOriginal(loc *time.Location) (time.Time, error) {
	var (
		t   time.Time
		err error
	)

	t, err = e.Time(AnyIFD, DateTimeOriginal, loc)
	if err != nil {
		return time.Time{}, err
	}

	entry, err := e.Entry(AnyIFD, SubSecTimeOriginal)
	if err == nil {
		str, err := entry.Ascii()
		if err == nil {
			n, err := parseSubSecTime(str)
			if err == nil {
				t = t.Add(time.Duration(n))
			}
		}
	}

	return t, nil
}

// Ascii returns the ASCII value for the given IFD and tag or returns
// an error. AnyIFD can be used to search all IFDs.
func (e *Exif) Ascii(ifd int, tag Tag) (string, error) {
	entry, err := e.Entry(ifd, tag)
	if err == nil {
		return entry.Ascii()
	}

	return e.Tiff.Ascii(ifd, tiff.Tag(tag))
}

var translate = map[string]string{
	"CASIO":                 "Casio",
	"SAMSUNG":               "Samsung",
	"samsung":               "Samsung",
	"NIKON CORPORATION":     "Nikon",
	"NIKON":                 "Nikon",
	"Eastman Kodak Company": "Kodak",
	"LG Electronics":        "LG",
	"HMD Global":            "Nokia",
}

// Make returns the make or returns an error. The make will be
// as "normalized" as possible, ie "Eastman Kodak Company" will
// be translated to "Kodak".
func (e *Exif) Make() (string, error) {
	make, _, err := e.MakeModel()

	return make, err
}

// Model returns the model of the camera or returns an error.
func (e *Exif) Model() (string, error) {
	_, model, err := e.MakeModel()

	return model, err
}

// ISO returns the ISO value or returns an error.
func (e *Exif) ISO() (int, error) {
	entry, err := e.Entry(AnyIFD, PhotographicSensitivity)
	if err != nil {
		entry, err = e.Entry(AnyIFD, ISOSpeed)
	}

	if err != nil {
		return 0, err
	}

	sh, err := entry.Short()

	return int(sh), err
}

// Apex is a type representing an APEX value.
type Apex float64

// FStops represent F-stops values.
type FStops float64

var log2 = math.Log(2)

// FStops returns the F-stops value for the APEX value.
func (a Apex) FStops() FStops {
	return FStops(math.Exp(float64(a) * log2 / 2.0))
}

// Aperture returns the aperture value or returns an error.
func (e *Exif) Aperture() (FStops, error) {
	entry, err := e.Entry(AnyIFD, ApertureValue)
	if err != nil {
		return 0, err
	}

	r, err := entry.Rational()
	if err != nil {
		return 0, err
	}

	a := Apex(r.Float())

	return a.FStops(), nil
}

// ExposureTime returns the exposure time or returns an error.
func (e *Exif) ExposureTime() (time.Duration, error) {
	entry, err := e.Entry(AnyIFD, ExposureTime)
	if err != nil {
		return 0, err
	}

	r, err := entry.Rational()
	if err != nil {
		return 0, err
	}

	// We special-case thousands, because it's quite
	// common for mobile cameras.
	if r.Denominator == 1000 {
		return time.Duration(r.Numerator) * time.Millisecond, nil
	}

	return time.Duration(r.Float() * float64(time.Second)), nil
}

// MakeModel returns the make and model of the camera or returns an
// error. The make will be as "normalized" as possible, ie "Eastman
// Kodak Company" will be translated to "Kodak".
func (e *Exif) MakeModel() (string, string, error) {
	make, err := e.Tiff.Ascii(AnyIFD, tiff.Make)
	if err != nil {
		return "", "", err
	}

	if t, ok := translate[make]; ok {
		make = t
	}

	model, err := e.Tiff.Ascii(AnyIFD, tiff.Model)
	if err != nil {
		return make, "", err
	}

	if strings.HasPrefix(strings.ToUpper(model), strings.ToUpper(make)) {
		return make, model[len(make)+1:], nil
	}

	return make, model, nil
}
