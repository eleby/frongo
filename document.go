package frongo

import "strings"

//Document type, used to store HTML containers,
//CSS sheets and a general HTML tag to contain the containers.
type Document struct {
	Containers []Container
	Sheets     []Sheet
	HTMLtag    *Tag
}

//Make a document out of a list of HTML containers.
//Takes the container list as parameter.
func (d *Document) Make(container ...Container) {
	d.Containers = container
}

//Specifies the list of sheets to be used by the document.
//Takes the sheet list as parameter.
func (d *Document) Style(sheet ...Sheet) {
	d.Sheets = sheet
}

//Adds a container to the document's container list.
//Takes the container as parameter.
func (d *Document) Put(container Container) {
	d.Containers = append(d.Containers, container)
}

//Function used to quickly get an initialized document.
func InitDocument() *Document {
	return &Document{HTMLtag: Put("html")}
}

//Function used by tags and styles to get their attributes
//from an ordered list of optional values.
//Takes the HTML class name or CSS selector,
//HTML attribute names or CSS modifier names,
//and HTML attribute values or CSS modifier values, as well as a delimiter.
//Default delimiter is ','.
func getParameters(params ...string) (string, [][]string) {
	mainParam := ""
	arrayNames := make([]string, 0)
	arrayValues := make([]string, 0)
	delim := ","
	if len(params) > 0 {
		mainParam = params[0]
		if len(params) > 2 {
			if len(params) > 3 {
				delim = params[3]
			}
			arrayNames = strings.Split(params[1], delim)
			arrayValues = strings.Split(params[2], delim)
		}
	}
	paramsArray := make([][]string, 0)
	for i := 0; i < len(arrayNames); i++ {
		paramsArray = append(paramsArray, []string{arrayNames[i], arrayValues[i]})
	}
	return mainParam, paramsArray
}
