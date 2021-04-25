package frongo

//Preset to make a title tag.
//Takes the title as parameter.
func Title(title string) *Tag {
	return &Tag{"title", "", make([][]string, 0), title, make([]*Tag, 0), "", OPENCLOSE}
}

//Preset to make a div tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func Div(params ...string) *Tag {
	return NewTag("div", params...)
}

//Preset to add a div tag to the tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func (t *Tag) Div(params ...string) *Tag {
	return t.NewTag("div", params...)
}

//Preset to make an img tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func Img(params ...string) *Tag {
	return NewTag("img", params...).Close()
}

//Preset to add an img tag to the tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func (t *Tag) Img(params ...string) *Tag {
	return t.NewTag("img", params...).Close()
}

//Preset to make a p tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func P(params ...string) *Tag {
	return NewTag("p", params...)
}

//Preset to add a p tag to the tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func (t *Tag) P(params ...string) *Tag {
	return t.NewTag("p", params...)
}

//Preset to make a span tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func Span(params ...string) *Tag {
	return NewTag("span", params...)
}

//Preset to add a span tag to the tag.
//Takes class, attribute names, attribute values,
//delimiter as optional parameters.
func (t *Tag) Span(params ...string) *Tag {
	return t.NewTag("span", params...)
}

//Preset to make a br tag.
func Br() *Tag {
	return NewTag("br").Close()
}

//Preset to add a br tag to the tag.
func (t *Tag) Br() *Tag {
	return t.Add("br").Close()
}
