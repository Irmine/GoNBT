package GoNBT

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

// The endianType field indicates whether to write in LittleEndian or BigEndian.
// LittleEndian is reverse byte order, BigEndian is normal byte order.
const (
	LittleEndian = iota
	BigEndian
)

func Read(buffer *[]byte, offset *int, length int) []byte {
	var initialLen = len((*buffer)[*offset:])
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Requested length:", length, ", have length:", initialLen)
		}
	}()

	var b = (*buffer)[*offset:*offset + length]
	*offset += length
	return b
}

func Write(buffer *[]byte, v byte){
	*buffer = append(*buffer, v)
}

func WriteBool(buffer *[]byte, bool bool) {
	if bool {
		WriteByte(buffer, 0x01)
		return
	}
	WriteByte(buffer, 0x00)
}

func ReadBool(buffer *[]byte, offset *int) bool {
	out := Read(buffer, offset, 1)
	return out[0] != 0x00
}

func WriteByte(buffer *[]byte, byte byte) {
	Write(buffer, byte)
}

func ReadByte(buffer *[]byte, offset *int) (byte) {
	out := Read(buffer, offset, 1)
	return byte(out[0])
}

func WriteUnsignedByte(buffer *[]byte, unsigned uint8) {
	WriteByte(buffer, byte(unsigned))
}

func ReadUnsignedByte(buffer *[]byte, offset *int) byte {
	out := Read(buffer, offset, 1)
	return byte(out[0])
}

func WriteShort(buffer *[]byte, signed int16) {
	var i uint
	var v uint
	len2 := uint(2)
	v = uint(len2 * 8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(signed >> v))
		v -= 8
	}
}

func ReadShort(buffer *[]byte, offset *int) int16 {
	var out int
	b := Read(buffer, offset, 2)
	var v uint = 8

	for i := 0; i < len(b); i++ {
		if i == 0 {
			out = int(b[i]) << v
			v -= 8
			break
		}
		out |= int(b[i]) << v
		v -= 8
	}

	return int16(out)
}

func WriteUnsignedShort(buffer *[]byte, int uint16) {
	var i uint
	var v uint
	len2 := uint(2)
	v = uint(len2 * 8) - 8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadUnsignedShort(buffer *[]byte, offset *int) uint16 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 2)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v -= 8
			continue
		}
		out |= int(b[i]) << v
		v -= 8
	}

	return uint16(out)
}

func WriteInt(buffer *[]byte, int int32) {
	var i uint
	var v uint
	len2 := uint(4)
	v = uint(len2 * 8) - 8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadInt(buffer *[]byte, offset *int) int32 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 4)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v -= 8
			continue
		}
		out |= int(b[i]) << v
		v -= 8
	}
	return int32(out)
}

func WriteLong(buffer *[]byte, int int64) {
	var i uint
	var v uint
	len2 := uint(8)
	v = uint(len2 * 8) - 8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadLong(buffer *[]byte, offset *int) int64 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v -= 8
			continue
		}
		out |= int(b[i]) << v
		v -= 8
	}
	return int64(out)
}

func WriteUnsignedLong(buffer *[]byte, int uint64) {
	var i uint
	var v uint
	len2 := uint(8)
	v = uint(len2*8)-8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(int >> v))
		v -= 8
	}
}

func ReadUnsignedLong(buffer *[]byte, offset *int) uint64 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v -= 8
			continue
		}
		out |= int(b[i]) << v
		v -= 8
	}
	return uint64(out)
}

func WriteFloat(buffer *[]byte, float float32) {
	var i uint
	var v uint
	x := math.Float32bits(float)
	len2 := uint(4)
	v = uint(len2 * 8) - 8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(x >> v))
		v -= 8
	}
}

func ReadFloat(buffer *[]byte, offset *int) float32 {
	var v uint
	var i uint
	var out uint32
	b := Read(buffer, offset, 4)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint32(b[i]) << v
			v -= 8
			continue
		}
		out |= uint32(b[i]) << v
		v -= 8
	}
	return math.Float32frombits(out)
}

func WriteDouble(buffer *[]byte, double float64) {
	var i uint
	var v uint
	x := math.Float64bits(double)
	len2 := uint(8)
	v = uint(len2 * 8) - 8
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(x >> v))
		v -= 8
	}
}

func ReadDouble(buffer *[]byte, offset *int) float64 {
	var v uint
	var i uint
	var out uint64
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = uint(len2 * 8) - 8
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint64(b[i]) << v
			v -= 8
			continue
		}
		out |= uint64(b[i]) << v
		v -= 8
	}
	return math.Float64frombits(out)
}

func WriteLittleShort(buffer *[]byte, short int16) {
	var i uint
	len2 := uint(2)
	for i = 0; i < len2 * 8; i += 8 {
		Write(buffer, byte(uint(short) >> i))
	}
}

func ReadLittleShort(buffer *[]byte, offset *int) int16 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 2)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v += 8
			continue
		}
		out |= int(b[i]) << v
		v += 8
	}
	return int16(out)
}

func WriteLittleUnsignedShort(buffer *[]byte, short uint16) {
	var i uint
	len2 := 2
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(short) >> i))
	}
}

func ReadLittleUnsignedShort(buffer *[]byte, offset *int) uint16 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 2)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v += 8
			continue
		}
		out |= int(b[i]) << v
		v += 8
	}
	return uint16(out)
}

func WriteLittleInt(buffer *[]byte, int int32) {
	var i uint
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleInt(buffer *[]byte, offset *int) int32 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 4)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v += 8
			continue
		}
		out |= int(b[i]) << v
		v += 8
	}
	return int32(out)
}

func WriteLittleLong(buffer *[]byte, int int64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleLong(buffer *[]byte, offset *int) int64 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v += 8
			continue
		}
		out |= int(b[i]) << v
		v += 8
	}
	return int64(out)
}

func WriteLittleUnsignedLong(buffer *[]byte, int uint64) {
	var i uint
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(uint(int) >> i))
	}
}

func ReadLittleUnsignedLong(buffer *[]byte, offset *int) uint64 {
	var v uint
	var i uint
	var out int
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = int(b[i]) << v
			v += 8
			continue
		}
		out |= int(b[i]) << v
		v += 8
	}
	return uint64(out)
}

func WriteLittleFloat(buffer *[]byte, f float32) {
	var i uint
	x := math.Float32bits(f)
	len2 := 4
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(x >> i))
	}
}

func ReadLittleFloat(buffer *[]byte, offset *int) float32 {
	var v uint
	var i uint
	var out uint32
	b := Read(buffer, offset, 4)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint32(b[i]) << v
			v += 8
			continue
		}
		out |= uint32(b[i]) << v
		v += 8
	}
	return math.Float32frombits(out)
}

func WriteLittleDouble(buffer *[]byte, double float64) {
	var i uint
	x := math.Float64bits(double)
	len2 := 8
	for i = 0; i < uint(len2) * 8; i += 8 {
		Write(buffer, byte(x >> i))
	}
}

func ReadLittleDouble(buffer *[]byte, offset *int) float64 {
	var v uint
	var i uint
	var out uint64
	b := Read(buffer, offset, 8)
	len2 := uint(len(b))
	v = 0
	for i = 0; i < len2; i++ {
		if i == 0 {
			out = uint64(b[i]) << v
			v += 8
			continue
		}
		out |= uint64(b[i]) << v
		v += 8
	}
	return math.Float64frombits(out)
}

func ReadBigEndianTriad(buffer *[]byte, offset *int) uint32 {
	var out uint32
	var b = Read(buffer, offset, 3)
	out = (uint32(b[2]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[0]) & 0x0F) << 16)

	return out
}

func WriteLittleEndianTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint & 0xFF))
	Write(buffer, byte(uint >> 8) & 0xFF)
	Write(buffer, byte(uint >> 16) & 0xFF)
}

func ReadLittleEndianTriad(buffer *[]byte, offset *int) uint32 {
	var b = Read(buffer, offset, 3)

	return (uint32(b[0]) & 0xFF) | ((uint32(b[1]) & 0xFF) << 8) | ((uint32(b[2]) & 0x0F) << 16)
}

func WriteBigEndianTriad(buffer *[]byte, uint uint32) {
	Write(buffer, byte(uint >> 16) & 0xFF)
	Write(buffer, byte(uint >> 8) & 0xFF)
	Write(buffer, byte(uint & 0xFF))
}

func ReadUnsignedVarInt(buffer *[]byte, offset *int) (uint32) {
	var out uint32 = 0
	for v := 0; v < 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= uint32((b & 0x7f) << uint(v))

		if (b & 0x80) == 0 {
			return out
		}
	}

	return 0
}

func WriteUnsignedVarInt(buffer *[]byte, int uint32) {
	if int == 0 {
		WriteByte(buffer, 0)
		return
	}
	var buf = make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(buf, uint64(int))

	buf = bytes.Trim(buf, "\x00")

	for _, b := range buf {
		Write(buffer, b)
	}
}

func WriteVarInt(buffer *[]byte, int int32) {
	if int == 0 {
		WriteByte(buffer, 0)
		return
	}
	var buf = make([]byte, binary.MaxVarintLen32)
	binary.PutVarint(buf, int64(int))

	buf = bytes.Trim(buf, "\x00")

	for _, b := range buf {
		Write(buffer, b)
	}
}

func ReadVarInt(buffer *[]byte, offset *int) int32 {
	var out int32 = 0
	for v := uint(0); v < 35; v += 7 {
		b := int(ReadByte(buffer, offset))
		out |= int32(b << v)

		if (b & 0x80) == 0 {
			return out
		}
	}

	return 0
}

func WriteVarLong(buffer *[]byte, int int64) {
	if int == 0 {
		WriteByte(buffer, 0)
		return
	}

	var bs = make([]byte, 10)
	binary.PutVarint(bs, int)
	bs = bytes.Trim(bs, "\x00")

	for _, b := range bs {
		WriteByte(buffer, b)
	}
}

func ReadVarLong(buffer *[]byte, offset *int) (int64) {
	var varLong, readBytes = binary.Varint((*buffer)[*offset:])
	var newOffset = readBytes + *offset
	offset = &newOffset

	return varLong
}