package frongo

import (
	"strings"
)

//Tag object, containing the characteristics of a HTML tag.
//A tag needs to be part of a Container object to be rendered.
type Tag struct {
	Tag        string
	Classes    string
	Attributes [][]string
	Text       string
	Children   []*Tag
	AfterText  string
	Type       Type
}

//Enum used to define if the tag is self-closing or an open tag.
type Type int

const (
	OPENCLOSE   = iota
	SELFCLOSING = iota
)

//Creates and get a tag.
//Takes the tag name as parameter.
func Put(name string) *Tag {
	return &Tag{name, "", make([][]string, 0), "", make([]*Tag, 0), "", OPENCLOSE}
}

//Writes text in the tag, before the eventual children.
//Takes the text to be written as parameter.
func (t *Tag) Write(text string) *Tag {
	t.Text = text
	return t
}

//Writes text after the tag.
//Takes the text to be written as parameter.
func (t *Tag) WriteAfter(text string) *Tag {
	t.AfterText = text
	return t
}

//Adds a tag to the tag's children.
//Takes the new tag's name as parameter.
func (t *Tag) Add(name string) *Tag {
	tag := Tag{name, "", make([][]string, 0), "", make([]*Tag, 0), "", OPENCLOSE}
	t.Children = append(t.Children, &tag)
	return &tag
}

//Removes a class from the tag's class list.
//Takes the class name as parameter.
func (t *Tag) RmClass(name string) *Tag {
	classArray := strings.Split(t.Classes, " ")
	for i := range classArray {
		if classArray[i] == name {
			classArray[i] = ""
		}
	}
	t.Classes = strings.Join(classArray, " ")
	return t
}

//Adds a class to the tag's class list.
//Takes the class name as parameter.
func (t *Tag) Class(name string) *Tag {
	prefix := ""
	if len(t.Classes) > 0 {
		prefix = " "
	}
	t.Classes += prefix + name
	return t
}

//Removes an attribute from the tag.
//Takes the attribute name as parameter.
func (t *Tag) RmAttr(name string) *Tag {
	for i := 0; i < len(t.Attributes); i++ {
		if t.Attributes[i][0] == name {
			t.Attributes = append(t.Attributes[:i], t.Attributes[i+1:]...)
		}
	}
	return t
}

//Adds an id to the tag.
//Takes the id attribute as parameter.
func (t *Tag) Id(id string) *Tag {
	t.Attr("id", id)
	return t
}

//Adds an attribute and its value to the tag.
//Takes the attribute name and value as parameters.
func (t *Tag) Attr(name string, value string) *Tag {
	t.Attributes = append(t.Attributes, []string{name, value})
	return t
}

//Adds attributes and their values to the tag.
//Takes the attribute names and values and an optional delimiter as parameters.
//The default delimiter is ","
func (t *Tag) Attrs(names string, values string, delim ...string) *Tag {
	delimiter := ","
	if len(delim) > 0 {
		delimiter = delim[0]
	}
	attrs := strings.Split(names, delimiter)
	vals := strings.Split(values, delimiter)
	for i := range attrs {
		t.Attributes = append(t.Attributes, []string{attrs[i], vals[i]})
	}
	return t
}

//Defines the tag as a self-closing tag.
func (t *Tag) Close() *Tag {
	t.Type = SELFCLOSING
	return t
}

//Creates and returns a new tag.
//Takes the tag name as parameter, and the tag class,
//attributes names, attributes values and delimiter as
//optional parameters.
func NewTag(tagName string, params ...string) *Tag {
	class, attrs := getParameters(params...)
	return &Tag{tagName, class, attrs, "", make([]*Tag, 0), "", OPENCLOSE}
}

//Creates and returns a new tag added to the tag's children.
//Takes the tag name as parameter, and the tag class,
//attributes names, attributes values and delimiter as
//optional parameters.
func (t *Tag) NewTag(tagName string, params ...string) *Tag {
	tag := t.Add(tagName)
	class, attrs := getParameters(params...)
	tag.Classes = class
	tag.Attributes = attrs
	return tag
}
