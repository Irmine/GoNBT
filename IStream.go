package GoNBT

type IStream interface {
	SetBuffer(Buffer []byte)
	GetBuffer() []byte
	Feof() bool
	Get(length int) []byte
	PutBool(v bool)
	GetBool() bool
	PutByte(v byte)
	GetByte() byte
	PutShort(v int16)
	GetShort() int16
	PutInt(v int32)
	GetInt() int32
	PutLong(v int64)
	GetLong() int64
	PutFloat(v float32)
	GetFloat() float32
	PutDouble(v float64)
	GetDouble() float64
	PutString(v string)
	GetString() string
	PutBytes(bytes []byte)
	ResetStream()
}
