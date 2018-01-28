package GoNBT

type IStream interface {
	SetBuffer([]byte)
	GetBuffer() []byte
	GetOffset() int
	SetOffset(int)
	Feof() bool
	Get(int) []byte
	PutBool(bool)
	GetBool() bool
	PutByte(byte)
	GetByte() byte
	PutShort(int16)
	GetShort() int16
	PutInt(int32)
	GetInt() int32
	PutLong(int64)
	GetLong() int64
	PutFloat(float32)
	GetFloat() float32
	PutDouble(float64)
	GetDouble() float64
	PutString(string)
	GetString() string
	PutBytes([]byte)
	ResetStream()
}
