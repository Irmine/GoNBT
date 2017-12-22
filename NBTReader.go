package GoNBT

type NBTReader struct {
	*BinaryStream
	network bool
	endianType byte
}

func NewNBTReader(buffer []byte, network bool, endianType byte) *NBTReader {
	return &NBTReader{NewStream(buffer, network, endianType), network, endianType & 0x01}
}

/**
 * Reads the entire buffer into a Compound.
 * Returns nil if the first tag was not a compound.
 */
func (reader *NBTReader) ReadIntoCompound() *Compound {
	var tag = reader.GetTag()
	if compound, ok := tag.(*Compound); ok {
		compound.Read(reader)
	}
	return nil
}

/**
 * Returns the named tag at the current offset.
 * This does not read tag data into the tag.
 */
func (reader *NBTReader) GetTag() INamedTag {
	if reader.Feof() {
		return NewEnd("")
	}
	var tagId = reader.GetByte()
	var tagCheck = GetTagById(tagId, "")
	if tagCheck == nil {
		return nil
	}

	var name = reader.GetString()

	return GetTagById(tagId, name)
}
