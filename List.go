package GoNBT

import (
	"fmt"
	"strconv"
	"strings"
)

// A List contains an array with tags of the same type.
type List struct {
	*NamedTag
	tags []INamedTag
	tagType byte
}

func NewList(name string, tagType byte, tags []INamedTag) *List {
	return &List{NewNamedTag(name, TAG_List, nil), tags, tagType}
}

func (list *List) Read(reader *NBTReader) {

}


// GetTags returns all tags in this list.
func (list *List) GetTags() []INamedTag {
	return list.tags
}


// GetTagType returns the tag type of this list.
func (list *List) GetTagType() byte {
	return list.tagType
}


// GetTag returns a tag at the given offset in the list.
func (list *List) GetTag(offset int) INamedTag {
	return list.tags[offset]
}


// AddTag Adds a tag to the list.
func (list *List) AddTag(tag INamedTag) {
	list.tags = append(list.tags, tag)
}


// Pop pushes the last tag off the list.
func (list *List) Pop() INamedTag {
	var tag = list.tags[len(list.tags) - 1]
	list.tags = list.tags[:len(list.tags) - 2]
	return tag
}


// Shift pushes the first tag off the list.
func (list *List) Shift() INamedTag {
	var tag = list.tags[0]
	list.tags = list.tags[1:]
	return tag
}


// DeleteAtOffset deletes a tag at the given offset and rearranges the list.
func (list *List) DeleteAtOffset(offset int) {
	if offset > len(list.tags) - 1 || offset < 0 {
		return
	}

	list.tags = append(list.tags[:offset], list.tags[offset + 1:]...)
}


// ToString converts the entire list to a readable string. Nesting level is used to indicate indentation.
func (list *List) toString(nestingLevel int) string {
	var str = strings.Repeat(" ", nestingLevel * 2)
	var entries = " entries"
	if len(list.tags) == 1 {
		entries = " entry"
	}

	str += "TAG_List('" + list.GetName() + " (" + GetTagName(list.tagType) + ")'): " + strconv.Itoa(len(list.tags)) + entries + "\n"
	str += strings.Repeat(" ", nestingLevel * 2) + "{\n"

	for _, tag := range list.tags {
		if list, ok := tag.(*List); ok {
			str += list.toString(nestingLevel + 1)
		} else {
			if compound, ok := tag.(*Compound); ok {
				str += compound.toString(nestingLevel + 1)
			} else {
				str += strings.Repeat(" ", (nestingLevel + 1) * 2)
				str += GetTagName(tag.GetTagType()) + "(None): " + fmt.Sprint(tag.Interface()) + "\n"
			}
		}
	}
	str += strings.Repeat(" ", nestingLevel * 2) + "}\n"
	return str
}