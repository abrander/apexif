package bmff

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/abrander/apexif/fileformats"
)

type Bmff struct {
	bytes     []byte
	iinfItems []item
}

type item struct {
	tag    string
	offset uint32
	length uint32
}

var Debug = false

func debugf(format string, args ...interface{}) {
	if Debug {
		fmt.Fprintf(os.Stderr, "DEBUG: "+format+"\n", args...)
	}
}

func returnErr(err error) error {
	if Debug && err != nil {
		panic(err.Error())
	}

	return err
}

func parseBox(data []byte) (length uint64, boxType string) {
	if len(data) < 8 {
		return 0, ""
	}

	length = uint64(binary.BigEndian.Uint32(data[0:4]))
	boxType = string(data[4:8])

	return
}

func parseFullbox(data []byte) (length uint64, boxType string, version uint8, flags uint32) {
	if len(data) < 12 {
		return
	}

	length, boxType = parseBox(data)

	version = uint8(data[8])

	flags = uint32(data[9])
	flags = flags<<8 | uint32(data[10])
	flags = flags<<8 | uint32(data[11])

	return
}

func Parse(data []byte) (*Bmff, error) {
	if len(data) < 12 {
		return nil, returnErr(fileformats.ErrImageNotRecognized)
	}

	if string(data[4:8]) != "ftyp" {
		return nil, returnErr(fileformats.ErrImageNotRecognized)
	}

	b := &Bmff{data, nil}

	for len(data) > 8 {
		length, err := b.parseBox(data[0:])
		if err != nil {
			return nil, err
		}

		if length == 0 {
			break
		}

		if len(data) < int(length) {
			return nil, returnErr(fileformats.ErrImageNotRecognized)
		}

		data = data[length:]
	}

	return b, nil
}

func (b *Bmff) Iloc(tag string) []byte {
	for _, item := range b.iinfItems {
		if item.tag == tag {
			return b.bytes[item.offset+10 : item.offset+10+item.length]
		}
	}

	return nil
}

func (b *Bmff) parseBox(data []byte) (length uint64, err error) {
	if len(data) < 8 {
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	length, boxType := parseBox(data)
	start := 8

	if len(data) < int(length) {
		length = 0
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	switch length {
	case 0:
		// Box extends to end of file.
		length = uint64(len(data))

	case 1:
		// 64-bit length
		length = binary.BigEndian.Uint64(data[8:16])
		start += 8
	}

	debugf("parseBox: type:%s length:%d", boxType, length)

	if length+1 > uint64(len(data)) {
		length = 0
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	if start >= int(length) {
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	switch boxType {
	case "ftyp":
		length, err = parseFtyp(data[start:length])
		return length + 8, err

	case "meta":
		length, err = b.parseMeta(data[start:length])
		return length + 8, err
	}

	return
}

func parseFtyp(data []byte) (length uint64, err error) {
	if len(data) < 8 {
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	length = uint64(len(data))

	return
}

func parseInfe(data []byte) (tag string) {
	if len(data) < 12 {
		return
	}

	tag = string(data[8:12])

	return
}

func (b *Bmff) parseIinf0(data []byte) {
	if len(data) < 8 {
		return
	}

	items := binary.BigEndian.Uint16(data[0:2])

	debugf("parseIinf0 items:%d", items)

	data = data[2:]

	if b.iinfItems == nil {
		b.iinfItems = make([]item, items)
	}

	if len(b.iinfItems) != int(items) {
		debugf("parseIinf0: item count mismatch: %d != %d", len(b.iinfItems), items)

		return
	}

	for i := range b.iinfItems {
		length, tag := parseBox(data)
		if tag == "" {
			debugf("parseIinf0: pos:%d No tag", i)

			return
		}

		if len(data) < 12+2 {
			debugf("parseIinf0 pos:%d: Too short", i)

			return
		}

		item := binary.BigEndian.Uint16(data[12:])

		if item < 1 || item > uint16(len(b.iinfItems)) {
			debugf("parseIinf0 pos:%d: item:%d Invalid item number", i, item)

			continue
		}

		if length > uint64(len(data)) {
			debugf("parseIinf0 pos:%d: infe: length:%d > len(data):%d", i, length, len(data))

			return
		}

		var infe string
		switch tag {
		case "infe":
			infe = parseInfe(data[8 : 8+length])
			b.iinfItems[i].tag = infe
		}

		debugf("parseIinf0 pos:%d item:%d tag:%s infe:%s", i, item, tag, infe)

		data = data[length:]
	}
}

func (b *Bmff) parseIloc1(data []byte) {
	if len(data) < 8 {
		debugf("parseIloc1: Too short")
		return
	}

	items := binary.BigEndian.Uint16(data[2:4])

	debugf("parseIloc1: items:%d", items)

	if b.iinfItems == nil {
		b.iinfItems = make([]item, items)
	}

	if len(data) < 16*int(items) {
		debugf("parseIloc1: Too short.")

		return
	}

	if len(b.iinfItems) != int(items) {
		debugf("parseIloc1: item count mismatch: %d != %d", len(b.iinfItems), items)

		return
	}

	data = data[4:]

	for i := range b.iinfItems {
		item := binary.BigEndian.Uint16(data[0:2])
		offset := binary.BigEndian.Uint32(data[8:12])
		length := binary.BigEndian.Uint32(data[12:16])

		if item < 1 || item > uint16(len(b.iinfItems)) {
			debugf("parseIloc1 item:%d: Invalid item number: %d", i, item)

			continue
		}

		b.iinfItems[item-1].offset = offset
		b.iinfItems[item-1].length = length

		debugf("parseIloc1 pos:%d item:%d: tag:%s, offset:0x%x, length:0x%x", i, item, b.iinfItems[i].tag, offset, length)

		if len(data) < 16 {
			return
		}

		data = data[16:]
	}
}

func (b *Bmff) parseMeta(data []byte) (length uint64, err error) {
	if len(data) < 8 {
		err = returnErr(fileformats.ErrImageNotRecognized)

		return
	}

	read := uint64(4)
	data = data[4:]

	for len(data) > 12 {
		length, tag, version, tags := parseFullbox(data)

		if length == 0 || length+12 > uint64(len(data)) {
			err = returnErr(fileformats.ErrImageNotRecognized)

			break
		}

		debugf("parseMeta tag:%s length:%d version:%d flags:%x", tag, length, version, tags)

		switch {
		case tag == "iinf" && version == 0:
			b.parseIinf0(data[12 : 12+length])

		case tag == "iloc" && version == 1:
			b.parseIloc1(data[12 : 12+length])
		}

		if uint64(len(data))+1 < length {
			err = returnErr(fileformats.ErrImageNotRecognized)

			break
		}

		read += length

		data = data[length:]
	}

	return read, err
}
