package frongo

//Search function used to check if the document contains an element.
//Takes the element name as parameter.
func (d *Document) Contains(name string) bool {
	contains := false
	for i := 0; i < len(d.Containers); i++ {
		for j := 0; j < len(d.Containers[i].Tags); j++ {
			if d.Containers[i].Tags[j].HasChild(name, true) {
				contains = true
			}
		}
	}
	return contains
}

//Search function used to get elements from the document.
//Takes the element name as parameter.
func (d *Document) Find(name string) []*Tag {
	var result = make([]*Tag, 0)
	for _, container := range d.Containers {
		for _, tag := range container.Tags {
			result = append(result, tag.Find(name)...)
		}
	}
	return result
}

//Search function used to check if the
//HTML tag has a child of a certain name.
//Takes the element name and a boolean to set recursion as parameters.
func (t *Tag) HasChild(name string, recursive ...bool) bool {
	for _, child := range t.Children {
		if child.Tag == name {
			return true
		} else if recursive != nil && recursive[0] {
			if child.HasChild(name, recursive...) {
				return true
			}
		}
	}
	return false
}

//Search function used to get elements from the tag's children.
//Takes the element name as parameter.
func (t *Tag) Find(name string) []*Tag {
	var result []*Tag
	for _, child := range t.Children {
		var subResult []*Tag
		if child.Tag == name {
			result = append(result, child)
		}
		subResult = child.Find(name)
		for _, childResult := range subResult {
			result = append(result, childResult)
		}
	}
	return result
}

//Search function used to get the element from the tag's children by id.
//Takes the id as parameter.
func (t *Tag) FindById(id string) *Tag {
	var result *Tag
	for _, child := range t.Children {
		if child.HasAttrValue("id", id) {
			return child
		} else {
			result = child.FindById(id)
			if result != nil {
				return result
			}
		}
	}
	return result
}

//Search function used to get an element from the tag's children by it's index.
//Takes the index as parameter.
func (t *Tag) Get(index int) *Tag {
	return t.Children[index]
}

//Search function used to get the last element from the tag's children.
func (t *Tag) Last() *Tag {
	return t.Children[len(t.Children)-1]
}

//Search function used to get the first element from the tag's children.
func (t *Tag) First() *Tag {
	return t.Get(0)
}

//Search function used to check if the tag has the attribute.
//Takes the attribute name as parameter.
func (t *Tag) HasAttr(name string) bool {
	if len(t.Attributes) == 0 {
		return false
	}
	for _, attr := range t.Attributes {
		if attr[0] == name {
			return true
		}
	}
	return false
}

//Search function used to check if the tag has the attribute with the specified value.
//Takes the attribute name and value as parameter.
func (t *Tag) HasAttrValue(name string, value string) bool {
	if len(t.Attributes) == 0 {
		return false
	}
	for _, attr := range t.Attributes {
		if attr[0] == name && attr[1] == value {
			return true
		}
	}
	return false
}
