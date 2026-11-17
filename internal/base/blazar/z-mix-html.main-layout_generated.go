package blazar

// Attr sets an attribute with its associated value, allowing for flexible HTML customization.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) Attr(name string, value any) *blazarMainLayout {
	c.MixAttr(name, value)
	return c
}

// Class assigns one or more classnames to an element, linking it to styles defined in a stylesheet.
// This can be called multiple times to assign multiple classnames.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) Class(class ...string) *blazarMainLayout {
	c.MixClass(class...)
	return c
}

// DataSet allows for storage of custom data specific to individual elements.
// This can be called multiple times to store multiple sets of data, often used for scripting purposes.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) DataSet(name string, value any) *blazarMainLayout {
	c.MixDataSet(name, value)
	return c
}

// DataSets denotes datasets linked to an element and can store multiple sets of data.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) DataSets(dataSets map[string]any) *blazarMainLayout {
	c.MixDataSets(dataSets)
	return c
}

// ID assigns a unique identifier to an element.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) ID(id string) *blazarMainLayout {
	c.MixID(id)
	return c
}

// Style assigns inline CSS styling to an element.
// This can be called multiple times to set multiple CSS styles.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) Style(name, value string) *blazarMainLayout {
	c.MixStyle(name, value)
	return c
}

// Styles allocates multiple CSS styles to an element.
// This accepts multiple styling definitions.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) Styles(styles map[string]string) *blazarMainLayout {
	c.MixStyles(styles)
	return c
}

// Title provides additional information about an element, typically displayed as a tooltip.
//
// This is for compatibility with go-app HTML components.
func (c *blazarMainLayout) Title(title string) *blazarMainLayout {
	c.MixTitle(title)
	return c
}
