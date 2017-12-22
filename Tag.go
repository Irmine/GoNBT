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


// NewNamedTag returns a new tag with given name, tag and value.
func NewNamedTag(name string, tagType byte, value interface{}) *NamedTag {
	return &NamedTag{&Tag{tagType, value}, name}
}


// GetTagType returns the tag type of a tag.
func (tag *Tag) GetTagType() byte {
	return tag.tagType
}


// IsOfType checks if the tag has the same type as the given type.
func (tag *Tag) IsOfType(tagType byte) bool {
	return tag.tagType == tagType
}


// IsCompatibleWith checks if the tag has the same type as the given tag.
func (tag *Tag) IsCompatibleWith(namedTag INamedTag) bool {
	return tag.tagType == namedTag.GetTagType()
}


// Interface returns the value of this tag.
func (tag *Tag) Interface() interface{} {
	return tag.value
}


// Read reads data into the tag from the NBT reader.
func (tag *Tag) Read(*NBTReader) {}


// GetName returns the name of the tag.
func (tag *NamedTag) GetName() string {
	return tag.name
}


// ToString converts the tag to readable string.
func (tag *NamedTag) ToString() string {
	return GetTagName(tag.GetTagType()) + "('" + tag.GetName() + "'): " + fmt.Sprint(tag.value) + "\n"
}
