package GoNBT

import (
	"strconv"
	"strings"
)

type Compound struct {
	*NamedTag
	tags map[string]INamedTag
}

func NewCompound(name string, tags map[string]INamedTag) *Compound {
	return &Compound{NewNamedTag(name, TAG_Compound, nil), tags}
}

/**
 * Reads the data of of the reader into the Compound.
 */
func (compound *Compound) Read(reader *NBTReader) {
	for {
		var tag = reader.GetTag()
		if tag == nil || tag.GetTagType() == TAG_End {
			return
		}
		tag.Read(reader)
	}
}

/**
 * Checks if the compound has a tag with the given name.
 */
func (compound *Compound) HasTag(name string) bool {
	var _, exists = compound.tags[name]
	return exists
}

/**
 * Checks if the compound has a tag with the given name and type.
 */
func (compound *Compound) HasTagWithType(name string, tagType byte) bool {
	if !compound.HasTag(name) {
		return false
	}
	var tag = compound.GetTag(name)
	return tag.IsOfType(tagType)
}

/**
 * Returns a tag with the given name.
 */
func (compound *Compound) GetTag(name string) INamedTag {
	if !compound.HasTag(name) {
		return nil
	}
	return compound.tags[name]
}

/**
 * Sets a tag.
 */
func (compound *Compound) SetTag(tag INamedTag) {
	compound.tags[tag.GetName()] = tag
}

/**
 * Returns all compound tags in a name => tag map.
 */
func (compound *Compound) GetTags() map[string]INamedTag {
	return compound.tags
}

/**
 * Sets a tag with the given name to the given byte.
 */
func (compound *Compound) SetByte(name string, value byte) {
	compound.tags[name] = NewByte(name, value)
}

/**
 * Returns a byte from the tag with the given name.
 * If a byte tag with the name does not exist, it returns the default value.
 */
func (compound *Compound) GetByte(name string, defaultValue byte) byte {
	if !compound.HasTagWithType(name, TAG_Byte) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(byte)
}

/**
 * Sets a tag with the given name to the given int16.
 */
func (compound *Compound) SetShort(name string, value int16) {
	compound.tags[name] = NewShort(name, value)
}

/**
 * Returns a short from the tag with the given name.
 * If a short tag with the name does not exist, it returns the default value.
 */
func (compound *Compound) GetShort(name string, defaultValue int16) int16 {
	if !compound.HasTagWithType(name, TAG_Short) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int16)
}

/**
 * Sets a tag with the given name to the given int32.
 */
func (compound *Compound) SetInt(name string, value int32) {
	compound.tags[name] = NewInt(name, value)
}

/**
 * Returns an int32 in an int tag with the given name.
 */
func (compound *Compound) GetInt(name string, defaultValue int32) int32 {
	if !compound.HasTagWithType(name, TAG_Int) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int32)
}

/**
 * Sets a tag with the given name to the given int64.
 */
func (compound *Compound) SetLong(name string, value int64) {
	compound.tags[name] = NewLong(name, value)
}

/**
 * Returns an int64 in a long tag with the given name.
 */
func (compound *Compound) GetLong(name string, defaultValue int64) int64 {
	if !compound.HasTagWithType(name, TAG_Long) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(int64)
}

/**
 * Sets a tag with the given name to the given float32.
 */
func (compound *Compound) SetFloat(name string, value float32) {
	compound.tags[name] = NewFloat(name, value)
}

/**
 * Returns a float32 in a float tag with the given name.
 */
func (compound *Compound) GetFloat(name string, defaultValue float32) float32 {
	if !compound.HasTagWithType(name, TAG_Float) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float32)
}

/**
 * Sets a tag with the given name to the given float64.
 */
func (compound *Compound) SetDouble(name string, value float64) {
	compound.tags[name] = NewDouble(name, value)
}

/**
 * Returns a float64 in a float tag with the given name.
 */
func (compound *Compound) GetDouble(name string, defaultValue float64) float64 {
	if !compound.HasTagWithType(name, TAG_Double) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(float64)
}

/**
 * Sets a tag with the given name to the given string.
 */
func (compound *Compound) SetString(name string, value string) {
	compound.tags[name] = NewString(name, value)
}

/**
 * Returns a string in a string tag with the given name.
 */
func (compound *Compound) GetString(name string, defaultValue string) string {
	if !compound.HasTagWithType(name, TAG_String) {
		return defaultValue
	}
	return compound.GetTag(name).Interface().(string)
}

/**
 * Sets a list with the given name.
 */
func (compound *Compound) SetList(name string, tagType byte, value []INamedTag) {
	compound.tags[name] = NewList(name, tagType, value)
}

/**
 * Returns a list with the given name and tag type.
 * If a list with that name and/or tag type does not exist, returns nil.
 */
func (compound *Compound) GetList(name string, tagType byte) *List {
	if !compound.HasTagWithType(name, TAG_List) {
		return nil
	}
	var list = compound.GetTag(name).(*List)
	if list.GetTagType() != tagType {
		return nil
	}
	return list
}

/**
 * Sets a compound with the given name.
 */
func (compound *Compound) SetCompound(name string, value map[string]INamedTag) {
	compound.tags[name] = NewCompound(name, value)
}

/**
 * Returns a compound with the given name.
 * If a compound with that name doesn't exist, returns nil.
 */
func (compound *Compound) GetCompound(name string) *Compound {
	if !compound.HasTagWithType(name, TAG_Compound) {
		return nil
	}
	return compound.tags[name].(*Compound)
}

/**
 * Returns the compound as an interface.
 */
func (compound *Compound) Interface() interface{} {
	return compound.tags
}

/**
 * Converts the entire compound to a readable string. Nesting level is used to indicate indentation.
 */
func (compound *Compound) toString(nestingLevel int) string {
	var str = strings.Repeat(" ", nestingLevel * 2)
	var entries = " entries"
	if len(compound.tags) == 1 {
		entries = " entry"
	}

	str += "TAG_Compound('" + compound.GetName() + "'): " + strconv.Itoa(len(compound.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel * 2) + "{\n"

	for _, tag := range compound.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel + 1)
			} else {
				str += strings.Repeat(" ", (nestingLevel + 1) * 2)
				str += tag.ToString()
			}
		}
	}
	str += strings.Repeat(" ", nestingLevel * 2) + "}\n"
	return str
}

/**
 * Converts the entire compound to an uncompressed string.
 */
func (compound *Compound) ToString() string {
	return compound.toString(0)
}
