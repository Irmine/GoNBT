package gonbt

import "fmt"

// A byte array holds an array filled with bytes.
// The payload is prefixed with an integer indicating the length.
type ByteArray struct {
	// []byte
	*NamedTag
	values []byte
}

// An int array holds an array filled with int32s.
// The payload is prefixed with an integer indicating the length.
type IntArray struct {
	// []int32
	*NamedTag
	values []int32
}

// A long array holds an array filled with int64s.
// The payload is prefixed with an integer indicating the length.
type LongArray struct {
	// []int64
	*NamedTag
	values []int64
}

func NewByteArray(name string, values []byte) *ByteArray {
	return &ByteArray{NewNamedTag(name, TAG_Byte_Array, nil), values}
}

func NewIntArray(name string, values []int32) *IntArray {
	return &IntArray{NewNamedTag(name, TAG_Int_Array, nil), values}
}

func NewLongArray(name string, values []int64) *LongArray {
	return &LongArray{NewNamedTag(name, TAG_Long_Array, nil), values}
}

func (tag *ByteArray) read(reader *Reader) {
	var length = reader.GetInt()
	tag.values = reader.GetBuffer()[reader.GetOffset():reader.GetOffset()+int(length)]
	reader.SetOffset(reader.GetOffset() + int(length))
}

func (tag *IntArray) read(reader *Reader) {
	var length = reader.GetInt()
	for i := int32(0); i < length; i++ {
		tag.values = append(tag.values, reader.GetInt())
	}
}

func (tag *LongArray) read(reader *Reader) {
	var length = reader.GetInt()
	for i := int32(0); i < length; i++ {
		tag.values = append(tag.values, reader.GetLong())
	}
}

func (tag *ByteArray) write(writer *Writer) {
	writer.PutInt(int32(len(tag.values)))
	for _, value := range tag.values {
		writer.PutByte(value)
	}
}

func (tag *IntArray) write(writer *Writer) {
	writer.PutInt(int32(len(tag.values)))
	for _, value := range tag.values {
		writer.PutInt(value)
	}
}

func (tag *LongArray) write(writer *Writer) {
	writer.PutInt(int32(len(tag.values)))
	for _, value := range tag.values {
		writer.PutLong(value)
	}
}

func (tag *ByteArray) ToString() string {
	return GetTagName(tag.GetType()) + "('" + tag.GetName() + "'): [" + fmt.Sprint(len(tag.values)) + " bytes]\n"
}

func (tag *IntArray) ToString() string {
	return GetTagName(tag.GetType()) + "('" + tag.GetName() + "'): [" + fmt.Sprint(len(tag.values)) + " integers]\n"
}

func (tag *LongArray) ToString() string {
	return GetTagName(tag.GetType()) + "('" + tag.GetName() + "'): [" + fmt.Sprint(len(tag.values)) + " longs]\n"
}

func (tag *ByteArray) Interface() interface{} {
	return tag.values
}

func (tag *IntArray) Interface() interface{} {
	return tag.values
}

func (tag *LongArray) Interface() interface{} {
	return tag.values
}
