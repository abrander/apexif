package riff

import (
	"encoding/binary"
	"fmt"
)

type Riff struct {
	Riff Chunk

	chunks []Chunk
}

type Chunk struct {
	Identifier string
	Length     uint32
	Data       []byte
}

func ReadChunk(data []byte, offset uint32) (Chunk, uint32, error) {
	if len(data[offset:]) < 8 {
		return Chunk{}, offset, fmt.Errorf("not enough data for chunk header, got %d bytes", len(data[offset:]))
	}

	identifier := string(data[offset : offset+4])
	length := binary.LittleEndian.Uint32(data[offset+4 : offset+8])

	if len(data[offset+8:]) < int(length) {
		return Chunk{}, offset, fmt.Errorf("not enough data for chunk")
	}

	chunkdata := data[offset+8 : offset+8+length]

	chunk := Chunk{
		Identifier: identifier,
		Length:     length,
		Data:       chunkdata,
	}

	if length%2 == 1 {
		length++
	}

	offset += 8 + length

	return chunk, offset, nil
}

func ReadChunks(data []byte) ([]Chunk, error) {
	offset := uint32(0)

	var chunks []Chunk
	var chunk Chunk
	var err error

	for offset < uint32(len(data)) {
		chunk, offset, err = ReadChunk(data, offset)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk)
	}

	return chunks, nil
}

func (c Chunk) String() string {
	return fmt.Sprintf("%s (%d bytes)", c.Identifier, c.Length)
}

func Parse(data []byte) (*Riff, error) {
	chunks, err := ReadChunks(data)
	if err != nil {
		return nil, err
	}

	r := &Riff{
		chunks: chunks,
	}

	for _, chunk := range chunks {
		if chunk.Identifier == "RIFF" {
			if r.Riff.Data != nil {
				return nil, fmt.Errorf("multiple RIFF chunks found in RIFF container")
			}

			r.Riff = chunk
		}
	}

	// RIFF containers must have exactly one RIFF chunk.
	if r.Riff.Data == nil {
		return nil, fmt.Errorf("no RIFF chunk found in RIFF container")
	}

	return r, nil
}

func (r *Riff) Chunks() []Chunk {
	return r.chunks
}
