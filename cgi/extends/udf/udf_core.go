package udf

import "murphyl.com/app/udf/internal/tag"

/* User Defined Function - 用户自定义功能 */

func NewSystemTag(code, name, desc string) *tag.Tag {
	return newTag(tag.TagTypeSystem, code, name, desc)
}

func newTag(tagType tag.TagType, code, name, desc string) *tag.Tag {
	return &tag.Tag{Type: tagType, Code: code, Name: name, Description: desc, Status: tag.TagStatusEnabled}
}
