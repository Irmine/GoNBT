package GoNBT

type BinaryStream struct {
	IStream
}

func NewStream(buffer []byte, network bool, endian byte) *BinaryStream {
	var stream IStream
	if network && endian == LittleEndian {
		stream = NewLittleEndianNetworkStream(buffer)
	} else {
		if endian == LittleEndian {
			stream = NewLittleEndianStream(buffer)
		} else {
			stream = NewBigEndianStream(buffer)
		}
	}
	return &BinaryStream{stream}
}

type LittleEndianStream struct {
	Offset int
	Buffer []byte
}

func NewLittleEndianStream(buffer []byte) *LittleEndianStream {
	return &LittleEndianStream{0, buffer}
}

func (stream *LittleEndianStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *LittleEndianStream) GetBuffer() []byte { return stream.Buffer }

func (stream *LittleEndianStream) Feof() bool { return stream.Offset >= len(stream.Buffer) - 1 }

func (stream *LittleEndianStream) Get(length int) []byte { return Read(&stream.Buffer, &stream.Offset, length) }

func (stream *LittleEndianStream) GetRemainder() []byte { return Read(&stream.Buffer, &stream.Offset, len(stream.Buffer) - stream.Offset)}

func (stream *LittleEndianStream) PutBool(v bool) { WriteBool(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetBool() bool { return ReadBool(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutByte(v byte) { WriteByte(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetByte() byte { return ReadByte(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutShort(v int16) { WriteLittleShort(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetShort() int16 { return ReadLittleShort(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutInt(v int32) { WriteLittleInt(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetInt() int32 { return ReadLittleInt(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutLong(v int64) { WriteLittleLong(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetLong() int64 { return ReadLittleLong(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutFloat(v float32) { WriteLittleFloat(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetFloat() float32 { return ReadLittleFloat(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutDouble(v float64) { WriteLittleDouble(&stream.Buffer, v) }

func (stream *LittleEndianStream) GetDouble() float64 { return ReadLittleDouble(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianStream) PutString(v string) {
	WriteUnsignedShort(&stream.Buffer, uint16(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *LittleEndianStream) GetString() string {
	var length = ReadUnsignedShort(&stream.Buffer, &stream.Offset)
	return string(Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *LittleEndianStream) PutBytes(bytes []byte) { stream.Buffer = append(stream.Buffer, bytes...) }

func (stream *LittleEndianStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}

type BigEndianStream struct {
	Offset int
	Buffer []byte
}

func NewBigEndianStream(buffer []byte) *BigEndianStream {
	return &BigEndianStream{0, buffer}
}

func (stream *BigEndianStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *BigEndianStream) GetBuffer() []byte { return stream.Buffer }

func (stream *BigEndianStream) Feof() bool { return stream.Offset >= len(stream.Buffer) - 1 }

func (stream *BigEndianStream) Get(length int) []byte { return Read(&stream.Buffer, &stream.Offset, length) }

func (stream *BigEndianStream) GetRemainder() []byte { return Read(&stream.Buffer, &stream.Offset, len(stream.Buffer) - stream.Offset)}

func (stream *BigEndianStream) PutBool(v bool) { WriteBool(&stream.Buffer, v) }

func (stream *BigEndianStream) GetBool() bool { return ReadBool(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutByte(v byte) { WriteByte(&stream.Buffer, v) }

func (stream *BigEndianStream) GetByte() byte { return ReadByte(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutShort(v int16) { WriteShort(&stream.Buffer, v) }

func (stream *BigEndianStream) GetShort() int16 { return ReadShort(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutInt(v int32) { WriteInt(&stream.Buffer, v) }

func (stream *BigEndianStream) GetInt() int32 { return ReadInt(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutLong(v int64) { WriteLong(&stream.Buffer, v) }

func (stream *BigEndianStream) GetLong() int64 { return ReadLong(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutFloat(v float32) { WriteFloat(&stream.Buffer, v) }

func (stream *BigEndianStream) GetFloat() float32 { return ReadFloat(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutDouble(v float64) { WriteDouble(&stream.Buffer, v) }

func (stream *BigEndianStream) GetDouble() float64 { return ReadDouble(&stream.Buffer, &stream.Offset) }

func (stream *BigEndianStream) PutString(v string) {
	WriteUnsignedShort(&stream.Buffer, uint16(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *BigEndianStream) GetString() string {
	var length = ReadUnsignedShort(&stream.Buffer, &stream.Offset)
	return string(Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *BigEndianStream) PutBytes(bytes []byte) { stream.Buffer = append(stream.Buffer, bytes...) }

func (stream *BigEndianStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}

type LittleEndianNetworkStream struct {
	Offset int
	Buffer []byte
}

func NewLittleEndianNetworkStream(buffer []byte) *LittleEndianNetworkStream {
	return &LittleEndianNetworkStream{0, buffer}
}

func (stream *LittleEndianNetworkStream) SetBuffer(Buffer []byte) { stream.Buffer = Buffer }

func (stream *LittleEndianNetworkStream) GetBuffer() []byte { return stream.Buffer }

func (stream *LittleEndianNetworkStream) Feof() bool { return stream.Offset >= len(stream.Buffer) - 1 }

func (stream *LittleEndianNetworkStream) Get(length int) []byte { return Read(&stream.Buffer, &stream.Offset, length) }

func (stream *LittleEndianNetworkStream) GetRemainder() []byte { return Read(&stream.Buffer, &stream.Offset, len(stream.Buffer) - stream.Offset)}

func (stream *LittleEndianNetworkStream) PutBool(v bool) { WriteBool(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetBool() bool { return ReadBool(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutByte(v byte) { WriteByte(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetByte() byte { return ReadByte(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutShort(v int16) { WriteLittleShort(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetShort() int16 { return ReadLittleShort(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutInt(v int32) { WriteVarInt(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetInt() int32 { return ReadVarInt(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutLong(v int64) { WriteVarLong(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetLong() int64 { return ReadVarLong(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutFloat(v float32) { WriteLittleFloat(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetFloat() float32 { return ReadLittleFloat(&stream.Buffer, &stream.Offset) }

func (stream *LittleEndianNetworkStream) PutDouble(v float64) { WriteLittleDouble(&stream.Buffer, v) }

func (stream *LittleEndianNetworkStream) GetDouble() float64 { return ReadLittleDouble(&stream.Buffer, &stream.Offset)}

func (stream *LittleEndianNetworkStream) PutString(v string) {
	WriteUnsignedVarInt(&stream.Buffer, uint32(len(v)))
	stream.Buffer = append(stream.Buffer, []byte(v)...)
}

func (stream *LittleEndianNetworkStream) GetString() string {
	var length = ReadUnsignedVarInt(&stream.Buffer, &stream.Offset)
	return string(Read(&stream.Buffer, &stream.Offset, int(length)))
}

func (stream *LittleEndianNetworkStream) PutBytes(bytes []byte) { stream.Buffer = append(stream.Buffer, bytes...) }

func (stream *LittleEndianNetworkStream) ResetStream() {
	stream.Offset = 0
	stream.Buffer = []byte{}
}