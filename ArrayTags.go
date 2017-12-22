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