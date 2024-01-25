package crw

import (
	"encoding/binary"
	"io"

	"github.com/abrander/apexif/containers/exif"
	"github.com/abrander/apexif/fileformats"
)

type CRW struct {
	bytes []byte
}

var _ fileformats.FileType = &CRW{}

func Identify(bytes []byte) (fileformats.FileType, error) {
	if len(bytes) < 10 {
		return nil, fileformats.ErrImageNotRecognized
	}

	if bytes[0] != 'I' || bytes[1] != 'I' {
		return nil, fileformats.ErrImageNotRecognized
	}

	if string(bytes[6:14]) != "HEAPCCDR" {
		return nil, fileformats.ErrImageNotRecognized
	}

	return &CRW{bytes: bytes}, nil
}

func (c *CRW) Name() string {
	return "CRW"
}

func (c *CRW) MediaType() string {
	return "image/x-canon-crw"
}

func (c *CRW) Exif() (*exif.Exif, error) {
	root := binary.LittleEndian.Uint32(c.bytes[2:6])

	offset := binary.LittleEndian.Uint32(c.bytes[len(c.bytes)-4:]) + root
	if offset > uint32(len(c.bytes)) {
		return nil, io.ErrUnexpectedEOF
	}

	data := c.bytes[root:]
	heap, err := readHeap(data)
	if err != nil {
		return nil, err
	}

	imageprops, err := heap.find(0xa)
	if err != nil {
		return nil, err
	}

	data = data[imageprops.Offset : imageprops.Offset+imageprops.Length]
	props, err := readHeap(data)
	if err != nil {
		return nil, err
	}

	buh, err := props.find(0xb)
	if err != nil {
		return nil, err
	}

	data = data[buh.Offset : buh.Offset+buh.Length]
	_, err = readHeap(data)

	return nil, err
}
