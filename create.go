package frongo

import (
	"html/template"
	"os"
)

//Save the document as a file.
//Takes the filename as parameter.
func (d *Document) Create(filename string) {
	content := template.HTML("")
	for _, c := range d.Containers {
		content += c.Render()
	}
	Create(filename, content)
}

//Save the HTML container as a file.
//Takes the filename as parameter.
func (c *Container) Create(filename string) {
	Create(filename, c.Render())
}

//Save the CSS sheet as a file.
//Takes the filename as parameter.
func (s *Sheet) Create(filename string) {
	Create(filename, s.Render())
}

//Save a html template string as a file.
//Takes the filename and the html template as parameters.
func Create(filename string, html template.HTML) {
	bytes := []byte(html)
	file, _ := os.Create(filename)
	defer file.Close()
	_, _ = file.Write(bytes)
}
