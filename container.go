package frongo

//Container type, used to store HTML tags.
type Container struct {
	Tags         []Tag
	ContainerTag Tag
}

//Make a container out of a list of tags.
//Takes the tag list as parameter.
func (c *Container) Make(tags ...Tag) {
	c.Tags = append(c.Tags, tags...)
}

//Initializes the container with a base tag.
//Takes the base tag name as parameter.
func InitContainer(name string) *Container {
	return &Container{Tags: make([]Tag, 0), ContainerTag: *NewTag(name)}
}

//Put a tag into the container.
//Takes the tag as parameter.
func (c *Container) Put(tag Tag) {
	c.Tags = append(c.Tags, tag)
}
