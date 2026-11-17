package blazar

// Attr sets an attribute with its associated value, allowing for flexible HTML customization.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) Attr(name string, value any) *blazarCard {
	c.MixAttr(name, value)
	return c
}

// Class assigns one or more classnames to an element, linking it to styles defined in a stylesheet.
// This can be called multiple times to assign multiple classnames.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) Class(class ...string) *blazarCard {
	c.MixClass(class...)
	return c
}

// DataSet allows for storage of custom data specific to individual elements.
// This can be called multiple times to store multiple sets of data, often used for scripting purposes.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) DataSet(name string, value any) *blazarCard {
	c.MixDataSet(name, value)
	return c
}

// DataSets denotes datasets linked to an element and can store multiple sets of data.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) DataSets(dataSets map[string]any) *blazarCard {
	c.MixDataSets(dataSets)
	return c
}

// ID assigns a unique identifier to an element.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) ID(id string) *blazarCard {
	c.MixID(id)
	return c
}

// Style assigns inline CSS styling to an element.
// This can be called multiple times to set multiple CSS styles.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) Style(name, value string) *blazarCard {
	c.MixStyle(name, value)
	return c
}

// Styles allocates multiple CSS styles to an element.
// This accepts multiple styling definitions.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) Styles(styles map[string]string) *blazarCard {
	c.MixStyles(styles)
	return c
}

// Title provides additional information about an element, typically displayed as a tooltip.
//
// This is for compatibility with go-app HTML components.
func (c *blazarCard) Title(title string) *blazarCard {
	c.MixTitle(title)
	return c
}
