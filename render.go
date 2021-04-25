package frongo

import (
	"html/template"
)

//Renders a document as a template.HTML using its elements.
func (d *Document) Render() template.HTML {
	var content template.HTML
	content = "<!doctype html>\n"
	content += "<" + template.HTML(d.HTMLtag.Tag) + " "
	content += template.HTML(d.HTMLtag.renderClass())
	content += template.HTML(d.HTMLtag.renderAttr())
	content += ">"
	for i := 0; i < len(d.Containers); i++ {
		content += d.Containers[i].Render()
	}
	content += "<style>"
	for i := 0; i < len(d.Sheets); i++ {
		content += d.Sheets[i].Render()
	}
	content += "</style>"
	content += "</" + template.HTML(d.HTMLtag.Tag) + ">"
	return content
}

//Renders a container as a template.HTML using its tags.
func (c *Container) Render() template.HTML {
	var content string
	if c.ContainerTag.Tag != "" {
		content += "<" + c.ContainerTag.Tag + " "
		content += c.ContainerTag.renderClass()
		content += c.ContainerTag.renderAttr()
		content += ">"
	}
	for i := 0; i < len(c.Tags); i++ {
		content += c.Tags[i].makeTag()
	}
	if c.ContainerTag.Tag != "" {
		content += "</" + c.ContainerTag.Tag + ">"
	}
	return template.HTML(content)
}

//Renders a sheet as a template.HTML using its styles.
func (s *Sheet) Render() template.HTML {
	var content string
	for i := 0; i < len(s.Styles); i++ {
		content += s.Styles[i].makeStyle()
	}
	return template.HTML(content)
}

//Renders a CSS style.
func (s *Style) makeStyle() string {
	str := s.Selector + "{"
	str += s.renderModifier()
	str += "}"
	return str
}

//Renders a style's modifier line.
func (s *Style) renderModifier() string {
	var mod string
	for i := 0; i < len(s.Modifier); i++ {
		mod += s.Modifier[i][0]
		mod += ":"
		mod += s.Modifier[i][1]
		mod += ";"
	}
	return mod
}

//Renders a HTML tag.
func (t *Tag) makeTag() string {
	str := "<" + t.Tag + t.renderClass() + t.renderAttr() + ">"
	if t.Type != SELFCLOSING {
		str += t.Text
		for i := range t.Children {
			str += t.Children[i].makeTag()
		}
		str += "</" + t.Tag + ">"
		str += t.AfterText
	}
	return str
}

//Renders a tag's class.
func (t *Tag) renderClass() string {
	var class string
	if t.Classes != "" {
		class = " class=\"" + t.Classes + "\""
	}
	return class
}

//Renders a tag's attributes.
func (t *Tag) renderAttr() string {
	var attr string
	for i := 0; i < len(t.Attributes); i++ {
		attr += " "
		attr += t.Attributes[i][0]
		attr += "=\""
		attr += t.Attributes[i][1]
		attr += "\""
	}
	return attr
}
