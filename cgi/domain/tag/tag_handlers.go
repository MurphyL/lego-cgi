package tag

import (
	"murphyl.com/lego/dal"
)

func NewSystemTag(code, name, desc string) *Tag {
	return NewTag(TypeSystem, code, name, desc)
}

func NewTag(tagType Type, code, name, desc string) *Tag {
	return &Tag{Type: tagType, Code: code, Name: name, Description: desc, Status: dal.StatusEnabled}
}
