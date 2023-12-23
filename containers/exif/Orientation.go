package exif

import (
	"fmt"

	"github.com/abrander/apexif/containers/tiff"
)

// Orientation is a type representing the orientation of the camera
// when the image was taken.
type Orientation uint16

const (
	Horizontal                Orientation = 1
	MirrorHorizontal          Orientation = 2
	Rotate180                 Orientation = 3
	MirrorVertical            Orientation = 4
	MirrorHorizontalRotate270 Orientation = 5
	Rotate90                  Orientation = 6
	MirrorHorizontalRotate90  Orientation = 7
	Rotate270                 Orientation = 8
)

var _ fmt.Stringer = Orientation(Horizontal)

// String returns a string representation of the orientation.
func (o Orientation) String() string {
	switch o {
	case Horizontal:
		return "Horizontal"
	case MirrorHorizontal:
		return "Mirror horizontal"
	case Rotate180:
		return "Rotate 180"
	case MirrorVertical:
		return "Mirror vertical"
	case MirrorHorizontalRotate270:
		return "Mirror horizontal and rotate 270 CW"
	case Rotate90:
		return "Rotate 90 CW"
	case MirrorHorizontalRotate90:
		return "Mirror horizontal and rotate 90 CW"
	case Rotate270:
		return "Rotate 270 CW"
	default:
		return "Unknown"
	}
}

// Orientation returns the orientation of the camera when the image
// was taken or an error.
func (e *Exif) Orientation() (Orientation, error) {
	entry, err := e.Tiff.Entry(0, tiff.Orientation)
	if err != nil {
		return 0, err
	}

	o, err := entry.Short()

	return Orientation(o), err
}
