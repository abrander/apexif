package crw

import (
	"encoding/binary"
	"fmt"
	"io"
)

type dataRecord struct {
	bytes  []byte
	Type   Type
	Offset uint32
	Length uint32
}

func (r dataRecord) String() string {
	return fmt.Sprintf("Type: %s, Offset: %d, Length: %d", r.Type, r.Offset, r.Length)
}

func readDataRecord(data []byte) (dataRecord, error) {
	if len(data) < 10 {
		return dataRecord{}, io.ErrUnexpectedEOF
	}

	var dr dataRecord

	dr.bytes = data[:10]
	dr.Type = Type(binary.LittleEndian.Uint16(data[0:]))
	dr.Length = binary.LittleEndian.Uint32(data[2:])
	dr.Offset = binary.LittleEndian.Uint32(data[6:])

	if (dr.Type & kStgFormatMask) == kStg_InRecordEntry {
		dr.Offset = 2
	}

	dr.Type &= 0x3fff

	return dr, nil
}
