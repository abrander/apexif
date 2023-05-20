package exif

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/abrander/apexif/containers/tiff"
)

type Exif struct {
	tiff.Tiff

	exifIDFPointer *tiff.IFD
}

var (
	ErrNoExifFound = errors.New("no EXIF data found")
)

const AnyIFD = tiff.AnyIFD

var unread = &tiff.IFD{}

func Parse(data []byte) (*Exif, error) {
	t, err := tiff.Parse(data)
	if err != nil {
		return nil, err
	}

	return &Exif{*t, unread}, nil
}

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

func (e *Exif) Make() (string, error) {
	make, _, err := e.MakeModel()

	return make, err
}

func (e *Exif) Model() (string, error) {
	_, model, err := e.MakeModel()

	return model, err
}

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

type Apex float64

type FStops float64

var log2 = math.Log(2)

func (a Apex) FStops() FStops {
	return FStops(math.Exp(float64(a) * log2 / 2.0))
}

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
