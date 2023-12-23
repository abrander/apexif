package exif

import (
	"fmt"

	"github.com/abrander/apexif/containers/tiff"
)

// Tag is a type representing an EXIF tag.
type Tag tiff.Tag

const (
	// Tags relating to version.
	ExifVersion     Tag = 0x9000
	FlashpixVersion Tag = 0xa000

	// Tags relating to image data characteristics.
	ColorSpace Tag = 0xa001
	Gamma      Tag = 0xa500

	// Tags relating to image configuration.
	ComponentsConfiguration Tag = 0x9101
	CompressedBitsPerPixel  Tag = 0x9102
	PixelXDimension         Tag = 0xa002
	PixelYDimension         Tag = 0xa003

	// Tags relating to user information.
	MakerNote   Tag = 0x927c
	UserComment Tag = 0x9286

	// Tags relating to related file information.
	RelatedSoundFile Tag = 0xa004

	// Tags relating to date and time.
	DateTimeOriginal    Tag = 0x9003
	DateTimeDigitized   Tag = 0x9004
	OffsetTime          Tag = 0x9010
	OffsetTimeOriginal  Tag = 0x9011
	OffsetTimeDigitized Tag = 0x9012
	SubSecTime          Tag = 0x9290
	SubSecTimeOriginal  Tag = 0x9291
	SubSecTimeDigitized Tag = 0x9292

	// Tags relating to picture-taking conditions.
	ExposureTime              Tag = 0x829a
	FNumber                   Tag = 0x829d
	ExposureProgram           Tag = 0x8822
	SpectralSensitivity       Tag = 0x8824
	PhotographicSensitivity   Tag = 0x8827 // Used to be ISO (!)
	OECF                      Tag = 0x8828
	SensitivityType           Tag = 0x8830
	StandardOutputSensitivity Tag = 0x8831
	RecommendedExposureIndex  Tag = 0x8832
	ISOSpeed                  Tag = 0x8833
	ISOSpeedLatitudeyyy       Tag = 0x8834
	ISOSpeedLatitudezzz       Tag = 0x8835
	ShutterSpeedValue         Tag = 0x9201
	ApertureValue             Tag = 0x9202
	BrightnessValue           Tag = 0x9203
	ExposureBiasValue         Tag = 0x9204
	MaxApertureValue          Tag = 0x9205
	SubjectDistance           Tag = 0x9206
	MeteringMode              Tag = 0x9207
	LightSource               Tag = 0x9208
	Flash                     Tag = 0x9209
	FocalLength               Tag = 0x920a
	SubjectArea               Tag = 0x9214
	FlashEnergy               Tag = 0xa20b
	SpatialFrequencyResponse  Tag = 0xa20c
	FocalPlaneXResolution     Tag = 0xa20e
	FocalPlaneYResolution     Tag = 0xa20f
	FocalPlaneResolutionUnit  Tag = 0xa210
	SubjectLocation           Tag = 0xa214
	ExposureIndex             Tag = 0xa215
	SensingMethod             Tag = 0xa217
	FileSource                Tag = 0xa300
	SceneType                 Tag = 0xa301
	CFAPattern                Tag = 0xa302
	CustomRendered            Tag = 0xa401
	ExposureMode              Tag = 0xa402
	WhiteBalance              Tag = 0xa403
	DigitalZoomRatio          Tag = 0xa404
	FocalLengthIn35mmFilm     Tag = 0xa405
	SceneCaptureType          Tag = 0xa406
	GainControl               Tag = 0xa407
	Contrast                  Tag = 0xa408
	Saturation                Tag = 0xa409
	Sharpness                 Tag = 0xa40a
	DeviceSettingDescription  Tag = 0xa40b
	SubjectDistanceRange      Tag = 0xa40c

	// Tags relating to shooting situation.
	Temperature          Tag = 0x9400
	Humidity             Tag = 0x9401
	Pressure             Tag = 0x9402
	WaterDepth           Tag = 0x9403
	Acceleration         Tag = 0x9404
	CameraElevationAngle Tag = 0x9405

	// Other tags.
	ImageUniqueID     Tag = 0xa420
	CameraOwnerName   Tag = 0xa430
	BodySerialNumber  Tag = 0xa431
	LensSpecification Tag = 0xa432
	LensMake          Tag = 0xa433
	LensModel         Tag = 0xa434
	LensSerialNumber  Tag = 0xa435

	// Tags relating to GPS.
	GPSVersionID         Tag = 0x0000
	GPSLatitudeRef       Tag = 0x0001
	GPSLatitude          Tag = 0x0002
	GPSLongitudeRef      Tag = 0x0003
	GPSLongitude         Tag = 0x0004
	GPSAltitudeRef       Tag = 0x0005
	GPSAltitude          Tag = 0x0006
	GPSTimeStamp         Tag = 0x0007
	GPSSatellites        Tag = 0x0008
	GPSStatus            Tag = 0x0009
	GPSMeasureMode       Tag = 0x000a
	GPSDOP               Tag = 0x000b
	GPSSpeedRef          Tag = 0x000c
	GPSSpeed             Tag = 0x000d
	GPSTrackRef          Tag = 0x000e
	GPSTrack             Tag = 0x000f
	GPSImgDirectionRef   Tag = 0x0010
	GPSImgDirection      Tag = 0x0011
	GPSMapDatum          Tag = 0x0012
	GPSDestLatitudeRef   Tag = 0x0013
	GPSDestLatitude      Tag = 0x0014
	GPSDestLongitudeRef  Tag = 0x0015
	GPSDestLongitude     Tag = 0x0016
	GPSDestBearingRef    Tag = 0x0017
	GPSDestBearing       Tag = 0x0018
	GPSDestDistanceRef   Tag = 0x0019
	GPSDestDistance      Tag = 0x001a
	GPSProcessingMethod  Tag = 0x001b
	GPSAreaInformation   Tag = 0x001c
	GPSDateStamp         Tag = 0x001d
	GPSDifferential      Tag = 0x001e
	GPSHPositioningError Tag = 0x001f

	InteroperabilityIFDPointer Tag = 0xa005
)

var _ fmt.Stringer = Tag(ExifVersion)

// String returns a string representation of the tag.
func (t Tag) String() string {
	m := map[Tag]string{
		ExifVersion:     "ExifVersion",
		FlashpixVersion: "FlashpixVersion",

		ColorSpace: "ColorSpace",
		Gamma:      "Gamma",

		ComponentsConfiguration: "ComponentsConfiguration",
		CompressedBitsPerPixel:  "CompressedBitsPerPixel",
		PixelXDimension:         "PixelXDimension",
		PixelYDimension:         "PixelYDimension",

		MakerNote:   "MakerNote",
		UserComment: "UserComment",

		RelatedSoundFile: "RelatedSoundFile",

		DateTimeOriginal:    "DateTimeOriginal",
		DateTimeDigitized:   "DateTimeDigitized",
		OffsetTime:          "OffsetTime",
		OffsetTimeOriginal:  "OffsetTimeOriginal",
		OffsetTimeDigitized: "OffsetTimeDigitized",
		SubSecTime:          "SubSecTime",
		SubSecTimeOriginal:  "SubSecTimeOriginal",
		SubSecTimeDigitized: "SubSecTimeDigitized",

		ExposureTime:              "ExposureTime",
		FNumber:                   "FNumber",
		ExposureProgram:           "ExposureProgram",
		SpectralSensitivity:       "SpectralSensitivity",
		PhotographicSensitivity:   "PhotographicSensitivity",
		OECF:                      "OECF",
		SensitivityType:           "SensitivityType",
		StandardOutputSensitivity: "StandardOutputSensitivity",
		RecommendedExposureIndex:  "RecommendedExposureIndex",
		ISOSpeed:                  "ISOSpeed",
		ISOSpeedLatitudeyyy:       "ISOSpeedLatitudeyyy",
		ISOSpeedLatitudezzz:       "ISOSpeedLatitudezzz",
		ShutterSpeedValue:         "ShutterSpeedValue",
		ApertureValue:             "ApertureValue",
		BrightnessValue:           "BrightnessValue",
		ExposureBiasValue:         "ExposureBiasValue",
		MaxApertureValue:          "MaxApertureValue",
		SubjectDistance:           "SubjectDistance",
		MeteringMode:              "MeteringMode",
		LightSource:               "LightSource",
		Flash:                     "Flash",
		FocalLength:               "FocalLength",
		SubjectArea:               "SubjectArea",
		FlashEnergy:               "FlashEnergy",
		SpatialFrequencyResponse:  "SpatialFrequencyResponse",
		FocalPlaneXResolution:     "FocalPlaneXResolution",
		FocalPlaneYResolution:     "FocalPlaneYResolution",
		FocalPlaneResolutionUnit:  "FocalPlaneResolutionUnit",
		SubjectLocation:           "SubjectLocation",
		ExposureIndex:             "ExposureIndex",
		SensingMethod:             "SensingMethod",
		FileSource:                "FileSource",
		SceneType:                 "SceneType",
		CFAPattern:                "CFAPattern",
		CustomRendered:            "CustomRendered",
		ExposureMode:              "ExposureMode",
		WhiteBalance:              "WhiteBalance",
		DigitalZoomRatio:          "DigitalZoomRatio",
		FocalLengthIn35mmFilm:     "FocalLengthIn35mmFilm",
		SceneCaptureType:          "SceneCaptureType",
		GainControl:               "GainControl",
		Contrast:                  "Contrast",
		Saturation:                "Saturation",
		Sharpness:                 "Sharpness",
		DeviceSettingDescription:  "DeviceSettingDescription",
		SubjectDistanceRange:      "SubjectDistanceRange",

		Temperature:          "Temperature",
		Humidity:             "Humidity",
		Pressure:             "Pressure",
		WaterDepth:           "WaterDepth",
		Acceleration:         "Acceleration",
		CameraElevationAngle: "CameraElevationAngle",

		ImageUniqueID:     "ImageUniqueID",
		CameraOwnerName:   "CameraOwnerName",
		BodySerialNumber:  "BodySerialNumber",
		LensSpecification: "LensSpecification",
		LensMake:          "LensMake",
		LensModel:         "LensModel",
		LensSerialNumber:  "LensSerialNumber",

		GPSVersionID:         "GPSVersionID",
		GPSLatitudeRef:       "GPSLatitudeRef",
		GPSLatitude:          "GPSLatitude",
		GPSLongitudeRef:      "GPSLongitudeRef",
		GPSLongitude:         "GPSLongitude",
		GPSAltitudeRef:       "GPSAltitudeRef",
		GPSAltitude:          "GPSAltitude",
		GPSTimeStamp:         "GPSTimeStamp",
		GPSSatellites:        "GPSSatellites",
		GPSStatus:            "GPSStatus",
		GPSMeasureMode:       "GPSMeasureMode",
		GPSDOP:               "GPSDOP",
		GPSSpeedRef:          "GPSSpeedRef",
		GPSSpeed:             "GPSSpeed",
		GPSTrackRef:          "GPSTrackRef",
		GPSTrack:             "GPSTrack",
		GPSImgDirectionRef:   "GPSImgDirectionRef",
		GPSImgDirection:      "GPSImgDirection",
		GPSMapDatum:          "GPSMapDatum",
		GPSDestLatitudeRef:   "GPSDestLatitudeRef",
		GPSDestLatitude:      "GPSDestLatitude",
		GPSDestLongitudeRef:  "GPSDestLongitudeRef",
		GPSDestLongitude:     "GPSDestLongitude",
		GPSDestBearingRef:    "GPSDestBearingRef",
		GPSDestBearing:       "GPSDestBearing",
		GPSDestDistanceRef:   "GPSDestDistanceRef",
		GPSDestDistance:      "GPSDestDistance",
		GPSProcessingMethod:  "GPSProcessingMethod",
		GPSAreaInformation:   "GPSAreaInformation",
		GPSDateStamp:         "GPSDateStamp",
		GPSDifferential:      "GPSDifferential",
		GPSHPositioningError: "GPSHPositioningError",

		InteroperabilityIFDPointer: "InteroperabilityIFDPointer",
	}

	if s, ok := m[t]; ok {
		return s
	}

	return tiff.Tag(t).String()
}
