package gonbt

import (
	"bytes"
	"compress/gzip"
	"github.com/irmine/binutils"
)

// The Writer is used to write compounds. (compressed)
// Network can be set to true to compact values, so ints become varints for example.
// EndianType should be either LittleEndian (0) or BigEndian (1), depending on the byte order the NBT should be written.
type Writer struct {
	*BinaryStream
}

// NewWriter returns a new NBT writer.
// Setting network to true will compact values, the endianType specifies byte order.
func NewWriter(network bool, endianType binutils.EndianType) *Writer {
	return &Writer{NewStream([]byte{}, network, endianType)}
}

// WriteUncompressedCompound writes a compound uncompressed.
func (writer *Writer) WriteUncompressedCompound(compound *Compound) {
	writer.PutTag(compound)
	compound.write(writer)
}

// WriteCompressedCompound writes a compound gzip compressed.
func (writer *Writer) WriteCompressedCompound(compound *Compound) {
	writer.WriteUncompressedCompound(compound)

	var buffer = bytes.NewBuffer(writer.GetBuffer())
	var gz = gzip.NewWriter(buffer)
	gz.Write(writer.GetData())
	defer gz.Close()

	writer.PutBytes(buffer.Bytes())
}

// PutTag puts a tag into the buffer.
// This does not yet write payload of the tag.
func (writer *Writer) PutTag(tag INamedTag) {
	writer.PutByte(tag.GetType())
	if tag.GetType() != TAG_End {
		writer.PutString(tag.GetName())
	}
}

// GetData returns the complete buffer/all data that has been written.
func (writer *Writer) GetData() []byte {
	return writer.GetBuffer()
}
