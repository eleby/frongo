package frongo

//Sheet type, used to store CSS styles.
type Sheet struct {
	Styles []Style
}

//Make a sheet out of a list of styles.
//Takes the style list as parameter.
func (s *Sheet) Make(styles ...Style) {
	s.Styles = styles
}

//Put a style into the sheet.
//Takes the style as parameter.
func (s *Sheet) Put(style Style) {
	s.Styles = append(s.Styles, style)
}

//Put a new style into the sheet.
//Takes the selector, modifier names,
//modifier values and delimiter as parameters.
func (s *Sheet) Style(params ...string) *Sheet {
	s.Styles = append(s.Styles, *NewStyle(params...))
	return s
}
