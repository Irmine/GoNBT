package GoNBT

import "fmt"

type INamedTag interface {
	GetTagType() byte
	ToString() string
	GetName() string
	Interface() interface{}
	IsCompatibleWith(INamedTag) bool
	IsOfType(byte) bool
	Read(*NBTReader)
}

type Tag struct {
	tagType byte
	value interface{}
}

type NamedTag struct {
	*Tag
	name string
}

func NewNamedTag(name string, tagType byte, value interface{}) *NamedTag {
	return &NamedTag{&Tag{tagType, value}, name}
}

/**
 * Returns the tag type of this tag.
 */
func (tag *Tag) GetTagType() byte {
	return tag.tagType
}

/**
 * Checks if the tag has the same type as the given type.
 */
func (tag *Tag) IsOfType(tagType byte) bool {
	return tag.tagType == tagType
}

/**
 * Checks if the tag has the same type as the given tag.
 */
func (tag *Tag) IsCompatibleWith(namedTag INamedTag) bool {
	return tag.tagType == namedTag.GetTagType()
}

/**
 * Returns the value of this tag.
 */
func (tag *Tag) Interface() interface{} {
	return tag.value
}

/**
 * Reads data into the tag.
 */
func (tag *Tag) Read(*NBTReader) {}

/**
 * Returns the name of this named tag.
 */
func (tag *NamedTag) GetName() string {
	return tag.name
}

/**
 * Converts the tag to readable string.
 */
func (tag *NamedTag) ToString() string {
	return GetTagName(tag.GetTagType()) + "('" + tag.GetName() + "'): " + fmt.Sprint(tag.value) + "\n"
}
