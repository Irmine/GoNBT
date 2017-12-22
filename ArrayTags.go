package GoNBT

type ByteArray struct {
	*NamedTag
	values []byte
}

type IntArray struct {
	*NamedTag
	values []int32
}

type LongArray struct {
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