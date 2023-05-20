package tiff

import (
	"fmt"
)

type Tag uint16

const (
	// Tags related to image data structure.
	ImageWidth                Tag = 0x0100
	ImageLength               Tag = 0x0101
	BitsPerSample             Tag = 0x0102
	Compression               Tag = 0x0103
	PhotometricInterpretation Tag = 0x0106
	Orientation               Tag = 0x0112
	SamplesPerPixel           Tag = 0x0115
	PlanarConfiguration       Tag = 0x011C
	YCbCrSubSampling          Tag = 0x0212
	YCbCrPositioning          Tag = 0x0213
	XResolution               Tag = 0x011A
	YResolution               Tag = 0x011B
	ResolutionUnit            Tag = 0x0128

	// Tags related to recording offset.
	StripOffsets                Tag = 0x0111
	RowsPerStrip                Tag = 0x0116
	StripByteCounts             Tag = 0x0117
	JPEGInterchangeFormat       Tag = 0x0201
	JPEGInterchangeFormatLength Tag = 0x0202

	// Tags related to image data characteristics.
	TransferFunction      Tag = 0x012D
	WhitePoint            Tag = 0x013E
	PrimaryChromaticities Tag = 0x013F
	YCbCrCoefficients     Tag = 0x0211
	ReferenceBlackWhite   Tag = 0x0214

	// Other tags.
	Datetime         Tag = 0x0132
	ImageDescription Tag = 0x010E
	Make             Tag = 0x010F
	Model            Tag = 0x0110
	Software         Tag = 0x0131
	Artist           Tag = 0x013B
	Copyright        Tag = 0x8298

	ExifIDFPointer    Tag = 0x8769
	GPSInfoIFDPointer Tag = 0x8825
)

func (t Tag) String() string {
	m := map[Tag]string{
		ImageWidth:                "ImageWidth",
		ImageLength:               "ImageLength",
		BitsPerSample:             "BitsPerSample",
		Compression:               "Compression",
		PhotometricInterpretation: "PhotometricInterpretation",
		Orientation:               "Orientation",
		SamplesPerPixel:           "SamplesPerPixel",
		PlanarConfiguration:       "PlanarConfiguration",
		YCbCrSubSampling:          "YCbCrSubSampling",
		YCbCrPositioning:          "YCbCrPositioning",
		XResolution:               "XResolution",
		YResolution:               "YResolution",
		ResolutionUnit:            "ResolutionUnit",

		StripOffsets:                "StripOffsets",
		RowsPerStrip:                "RowsPerStrip",
		StripByteCounts:             "StripByteCounts",
		JPEGInterchangeFormat:       "JPEGInterchangeFormat",
		JPEGInterchangeFormatLength: "JPEGInterchangeFormatLength",

		TransferFunction:      "TransferFunction",
		WhitePoint:            "WhitePoint",
		PrimaryChromaticities: "PrimaryChromaticities",
		YCbCrCoefficients:     "YCbCrCoefficients",
		ReferenceBlackWhite:   "ReferenceBlackWhite",

		Datetime:         "Datetime",
		ImageDescription: "ImageDescription",
		Make:             "Make",
		Model:            "Model",
		Software:         "Software",
		Artist:           "Artist",
		Copyright:        "Copyright",

		ExifIDFPointer:    "ExifIDFPointer",
		GPSInfoIFDPointer: "GPSInfoIFDPointer",
	}

	if s, ok := m[t]; ok {
		return s
	}

	return fmt.Sprintf("UNKNOWN:%04x", uint16(t))
}
