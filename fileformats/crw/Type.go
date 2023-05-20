package crw

import (
	"fmt"
)

type Type uint16

const (
	kStgFormatMask  Type = 0xc000
	kDataTypeMask   Type = 0x3800
	kIDCodeMask     Type = 0x07ff
	kTypeIDCodeMask Type = kDataTypeMask | kIDCodeMask

	kStg_InHeapSpace   Type = 0x0000
	kStg_InRecordEntry Type = 0x4000
)

const (
	kDT_BYTE              Type = 0x0000
	kDT_ASCII             Type = 0x0800
	kDT_WORD              Type = 0x1000
	kDT_DWORD             Type = 0x1800
	kDT_BYTE2             Type = 0x2000
	kDT_HeapTypeProperty1 Type = 0x2800
	kDT_HeapTypeProperty2 Type = 0x3000

	kTC_WildCard Type = 0xffff
	kTC_Null     Type = 0x0000
	kTC_Free     Type = 0x0001
	kTC_ExUsed   Type = 0x0002

	kTC_Description       Type = kDT_ASCII | 0x0005
	kTC_ModelName         Type = kDT_ASCII | 0x000a
	kTC_FirmwareVersion   Type = kDT_ASCII | 0x000b
	kTC_ComponentVersion  Type = kDT_ASCII | 0x000c
	kTC_ROMOperationMode  Type = kDT_ASCII | 0x000d
	kTC_OwnerName         Type = kDT_ASCII | 0x0010
	kTC_ImageFileName     Type = kDT_ASCII | 0x0016
	kTC_ThumbnailFileName Type = kDT_ASCII | 0x0017

	kTC_TargetImageType  Type = kDT_WORD | 0x000a
	kTC_SR_ReleaseMethod Type = kDT_WORD | 0x0010
	kTC_SR_ReleaseTiming Type = kDT_WORD | 0x0011
	kTC_ReleaseSetting   Type = kDT_WORD | 0x0016
	kTC_BodySensitivity  Type = kDT_WORD | 0x001c

	kTC_ImageFormat              Type = kDT_DWORD | 0x0003
	kTC_RecordID                 Type = kDT_DWORD | 0x0004
	kTC_SelfTimerTime            Type = kDT_DWORD | 0x0006
	kTC_SR_TargetDistanceSetting Type = kDT_DWORD | 0x0007
	kTC_BodyID                   Type = kDT_DWORD | 0x000b
	kTC_CapturedTime             Type = kDT_DWORD | 0x000e
	kTC_ImageSpec                Type = kDT_DWORD | 0x0010
	kTC_SR_EF                    Type = kDT_DWORD | 0x0013
	kTC_MI_EV                    Type = kDT_DWORD | 0x0014
	kTC_SerialNumber             Type = kDT_DWORD | 0x0017
	kTC_SR_Exposure              Type = kDT_DWORD | 0x0018

	kTC_CameraObject        Type = kDT_HeapTypeProperty1 | 0x0007
	kTC_ShootingRecord      Type = kDT_HeapTypeProperty1 | 0x0002
	kTC_MeasuredInfo        Type = kDT_HeapTypeProperty1 | 0x0003
	kTC_CameraSpecificaiton Type = kDT_HeapTypeProperty2 | 0x0004

	RawData             Type = kDT_BYTE2 | 0x0005
	JpgFromRaw          Type = kDT_BYTE2 | 0x0007
	ThumbnailImage      Type = kDT_BYTE2 | 0x0008
	MeasuredInfo        Type = kDT_HeapTypeProperty2 | 0x0003
	ImageDescription    Type = kDT_HeapTypeProperty1 | 0x0004
	ImageProps          Type = kDT_HeapTypeProperty2 | 0x000a
	ExifInformation     Type = kDT_HeapTypeProperty2 | 0x000b
	CanonFlashInfo      Type = kDT_WORD | 0x0028
	FocalLength         Type = kDT_WORD | 0x0029
	CanonShotInfo       Type = kDT_WORD | 0x002a
	CanonFileInfo       Type = kDT_WORD | 0x0093
	CanonCameraSettings Type = kDT_WORD | 0x002d
	CanonModelID        Type = kDT_DWORD | 0x0034
	SensorInfo          Type = kDT_WORD | 0x0031
	DecoderTable        Type = kDT_DWORD | 0x0035
	CanonAFInfo         Type = kDT_WORD | 0x0038
	ColorSpace          Type = kDT_WORD | 0x00b4
	RawJpgInfo          Type = kDT_WORD | 0x00b5
	SerialNumberFormat  Type = kDT_DWORD | 0x003b
)

func (t Type) String() string {
	ret := "UNKNOWN "
	switch t & kDataTypeMask {
	case kDT_BYTE:
		ret = "[BYTE] "
	case kDT_ASCII:
		ret = "[ASCII] "
	case kDT_WORD:
		ret = "[WORD] "
	case kDT_DWORD:
		ret = "[DWORD] "
	case kDT_BYTE2:
		ret = "[BYTE2] "
	case kDT_HeapTypeProperty1:
		ret = "[HeapTypeProperty1] "
	case kDT_HeapTypeProperty2:
		ret = "[HeapTypeProperty2] "
	}

	switch t & kTypeIDCodeMask {
	case kTC_WildCard:
		ret += "Wildcard"
	case kTC_Null:
		ret += "Null"
	case kTC_Free:
		ret += "Free"
	case kTC_ExUsed:
		ret += "ExUsed"
	case kTC_Description:
		ret += "Description"
	case kTC_ModelName:
		ret += "ModelName"
	case kTC_FirmwareVersion:
		ret += "FirmwareVersion"
	case kTC_ComponentVersion:
		ret += "ComponentVersion"
	case kTC_ROMOperationMode:
		ret += "ROMOperationMode"
	case kTC_OwnerName:
		ret += "OwnerName"
	case kTC_ImageFileName:
		ret += "ImageFileName"
	case kTC_ThumbnailFileName:
		ret += "ThumbnailFileName"
	case kTC_TargetImageType:
		ret += "TargetImageType"
	case kTC_SR_ReleaseMethod:
		ret += "SR_ReleaseMethod"
	case kTC_SR_ReleaseTiming:
		ret += "SR_ReleaseTiming"
	case kTC_ReleaseSetting:
		ret += "ReleaseSetting"
	case kTC_BodySensitivity:
		ret += "BodySensitivity"
	case kTC_ImageFormat:
		ret += "ImageFormat"
	case kTC_RecordID:
		ret += "RecordID"
	case kTC_SelfTimerTime:
		ret += "SelfTimerTime"
	case kTC_SR_TargetDistanceSetting:
		ret += "SR_TargetDistanceSetting"
	case kTC_BodyID:
		ret += "BodyID"
	case kTC_CapturedTime:
		ret += "CapturedTime"
	case kTC_ImageSpec:
		ret += "ImageSpec"
	case kTC_SR_EF:
		ret += "SR_EF"
	case kTC_MI_EV:
		ret += "MI_EV"
	case kTC_SerialNumber:
		ret += "SerialNumber"
	case kTC_SR_Exposure:
		ret += "SR_Exposure"
	case kTC_CameraObject:
		ret += "CameraObject"
	case kTC_ShootingRecord:
		ret += "ShootingRecord"

	case RawData:
		ret += "RawData"
	case JpgFromRaw:
		ret += "JpgFromRaw"
	case ThumbnailImage:
		ret += "ThumbnailImage"
	case MeasuredInfo:
		ret += "MeasuredInfo"
	case ImageDescription:
		ret += "ImageDescription"
	case ImageProps:
		ret += "ImageProps"
	case ExifInformation:
		ret += "ExifInformation"
	case CanonFlashInfo:
		ret += "CanonFlashInfo"
	case FocalLength:
		ret += "FocalLength"
	case CanonShotInfo:
		ret += "CanonShotInfo"
	case CanonFileInfo:
		ret += "CanonFileInfo"
	case CanonCameraSettings:
		ret += "CanonCameraSettings"
	case CanonModelID:
		ret += "CanonModelID"
	case SensorInfo:
		ret += "SensorInfo"
	case DecoderTable:
		ret += "DecoderTable"
	case CanonAFInfo:
		ret += "CanonAFInfo"
	case ColorSpace:
		ret += "ColorSpace"
	case RawJpgInfo:
		ret += "RawJpgInfo"
	case SerialNumberFormat:
		ret += "SerialNumberFormat"

	default:
		ret += fmt.Sprintf("\033[35mUNKNOWN\033[0m (%04x)", uint16(t))
	}

	return ret
}
