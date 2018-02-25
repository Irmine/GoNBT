package gonbt

// An end is used to indicate the end of a compound tag. An end has no payload.
type End struct {
	*NamedTag
}

// A byte has a payload of 1 byte.
type Byte struct {
	// (byte)
	*NamedTag
}

// A short has a payload of 2 bytes.
type Short struct {
	// (int16)
	*NamedTag
}

// An int has a payload of 4 bytes.
type Int struct {
	// (int32)
	*NamedTag
}

// A long has a payload of 8 bytes.
type Long struct {
	// (int64)
	*NamedTag
}

// A float has a payload of 4 bytes.
type Float struct {
	// (float32)
	*NamedTag
}

// A double has a payload of 8 bytes.
type Double struct {
	// (float64)
	*NamedTag
}

// A string has a variable payload, length indicated by a varInt/short, depending on the network field.
type String struct {
	// (string)
	*NamedTag
}

func NewEnd(name string) *End { return &End{NewNamedTag(name, TAG_End, 0)} }

func NewByte(name string, value byte) *Byte { return &Byte{NewNamedTag(name, TAG_Byte, value)} }

func NewShort(name string, value int16) *Short { return &Short{NewNamedTag(name, TAG_Short, value)} }

func NewInt(name string, value int32) *Int { return &Int{NewNamedTag(name, TAG_Int, value)} }

func NewLong(name string, value int64) *Long { return &Long{NewNamedTag(name, TAG_Long, value)} }

func NewFloat(name string, value float32) *Float { return &Float{NewNamedTag(name, TAG_Float, value)} }

func NewDouble(name string, value float64) *Double { return &Double{NewNamedTag(name, TAG_Double, value)} }

func NewString(name, value string) *String { return &String{NewNamedTag(name, TAG_String, value)} }

func (tag *Byte) read(reader *Reader) { tag.value = reader.GetByte() }

func (tag *Short) read(reader *Reader) { tag.value = reader.GetShort() }

func (tag *Int) read(reader *Reader) { tag.value = reader.GetInt() }

func (tag *Long) read(reader *Reader) { tag.value = reader.GetLong() }

func (tag *Float) read(reader *Reader) { tag.value = reader.GetFloat() }

func (tag *Double) read(reader *Reader) { tag.value = reader.GetDouble() }

func (tag *String) read(reader *Reader) { tag.value = reader.GetString() }

func (tag *Byte) write(writer *Writer) { writer.PutByte(tag.value.(byte)) }

func (tag *Short) write(writer *Writer) { writer.PutShort(tag.value.(int16)) }

func (tag *Int) write(writer *Writer) { writer.PutInt(tag.value.(int32)) }

func (tag *Long) write(writer *Writer) { writer.PutLong(tag.value.(int64)) }

func (tag *Float) write(writer *Writer) { writer.PutFloat(tag.value.(float32)) }

func (tag *Double) write(writer *Writer) { writer.PutDouble(tag.value.(float64)) }

func (tag *String) write(writer *Writer) { writer.PutString(tag.value.(string)) }
