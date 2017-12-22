package GoNBT

type End struct {
	*NamedTag
}

type Byte struct {
	*NamedTag
}

type Short struct {
	*NamedTag
}

type Int struct {
	*NamedTag
}

type Long struct {
	*NamedTag
}

type Float struct {
	*NamedTag
}

type Double struct {
	*NamedTag
}

type String struct {
	*NamedTag
}

func NewEnd(name string) *End {
	return &End{NewNamedTag(name, TAG_End, 0)}
}

func NewByte(name string, value byte) *Byte {
	return &Byte{NewNamedTag(name, TAG_Byte, value)}
}

func NewShort(name string, value int16) *Short {
	return &Short{NewNamedTag(name, TAG_Short, value)}
}

func NewInt(name string, value int32) *Int {
	return &Int{NewNamedTag(name, TAG_Int, value)}
}

func NewLong(name string, value int64) *Long {
	return &Long{NewNamedTag(name, TAG_Long, value)}
}

func NewFloat(name string, value float32) *Float {
	return &Float{NewNamedTag(name, TAG_Float, value)}
}

func NewDouble(name string, value float64) *Double {
	return &Double{NewNamedTag(name, TAG_Double, value)}
}

func NewString(name, value string) *String {
	return &String{NewNamedTag(name, TAG_String, value)}
}

/**
 * Reads a byte into the tag.
 */
func (tag *Byte) Read(reader *NBTReader) {
	tag.value = reader.GetByte()
}

/**
 * Reads an int16 into the short tag.
 */
func (tag *Short) Read(reader *NBTReader) {
	tag.value = reader.GetShort()
}

/**
 * Reads a (var)int32 into the tag.
 */
func (tag *Int) Read(reader *NBTReader) {
	tag.value = reader.GetInt()
}

/**
 * Reads a (var)int64 into the tag.
 */
func (tag *Long) Read(reader *NBTReader) {
	tag.value = reader.GetLong()
}

/**
 * Reads a float32 into the tag.
 */
func (tag *Float) Read(reader *NBTReader) {
	tag.value = reader.GetFloat()
}
/**
 * Reads a float64 into the tag.
 */
func (tag *Double) Read(reader *NBTReader) {
	tag.value = reader.GetDouble()
}

/**
 * Reads a string into the tag.
 */
func (tag *String) Read(reader *NBTReader) {
	tag.value = reader.GetString()
}