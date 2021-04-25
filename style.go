package frongo

//Style object, containing a selector and a list of modifiers.
//A style needs to be part of a Sheet object to be rendered.
type Style struct {
	Selector string
	Modifier [][]string
}

//Function to create a new Style object from the specified parameters.
//Takes a selector, a list of modifier names, a list of modifier values
//and a delimiter as parameters.
func NewStyle(params ...string) *Style {
	var style = &Style{Selector: "", Modifier: make([][]string, 0)}
	style.Selector, style.Modifier = getParameters(params...)
	return style
}
