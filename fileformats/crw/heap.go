package crw

import (
	"encoding/binary"
	"fmt"
	"io"
)

type heap struct {
	bytes   []byte
	records []dataRecord
}

var (
	errTagNotFound = fmt.Errorf("tag not found")
)

func (h *heap) find(tag Type) (dataRecord, error) {
	for _, r := range h.records {
		if r.Type&kIDCodeMask == tag {
			return r, nil
		}
	}

	return dataRecord{}, errTagNotFound
}

func (h *heap) Bytes(record dataRecord) ([]byte, error) {
	if record.Type&kStgFormatMask == kStg_InRecordEntry {
		return record.bytes[2:], nil
	}

	if record.Offset+record.Length > uint32(len(h.bytes)) {
		return nil, io.ErrUnexpectedEOF
	}

	return h.bytes[record.Offset : record.Offset+record.Length], nil
}

func readHeap(data []byte) (*heap, error) {
	offsetTblOffset := binary.LittleEndian.Uint32(data[len(data)-4:])
	records := binary.LittleEndian.Uint16(data[offsetTblOffset:])

	h := &heap{
		bytes:   data,
		records: make([]dataRecord, records),
	}

	for r := uint16(0); r < records; r++ {
		offset := offsetTblOffset + 2 + (10 * uint32(r))

		record, err := readDataRecord(data[offset : offset+10])
		if err != nil {
			return nil, err
		}

		h.records[r] = record
	}

	return h, nil
}
