package blazar

// Attr sets an attribute with its associated value, allowing for flexible HTML customization.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) Attr(name string, value any) *blazarInputWrapper {
	c.MixAttr(name, value)
	return c
}

// Class assigns one or more classnames to an element, linking it to styles defined in a stylesheet.
// This can be called multiple times to assign multiple classnames.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) Class(class ...string) *blazarInputWrapper {
	c.MixClass(class...)
	return c
}

// DataSet allows for storage of custom data specific to individual elements.
// This can be called multiple times to store multiple sets of data, often used for scripting purposes.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) DataSet(name string, value any) *blazarInputWrapper {
	c.MixDataSet(name, value)
	return c
}

// DataSets denotes datasets linked to an element and can store multiple sets of data.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) DataSets(dataSets map[string]any) *blazarInputWrapper {
	c.MixDataSets(dataSets)
	return c
}

// ID assigns a unique identifier to an element.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) ID(id string) *blazarInputWrapper {
	c.MixID(id)
	return c
}

// Style assigns inline CSS styling to an element.
// This can be called multiple times to set multiple CSS styles.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) Style(name, value string) *blazarInputWrapper {
	c.MixStyle(name, value)
	return c
}

// Styles allocates multiple CSS styles to an element.
// This accepts multiple styling definitions.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) Styles(styles map[string]string) *blazarInputWrapper {
	c.MixStyles(styles)
	return c
}

// Title provides additional information about an element, typically displayed as a tooltip.
//
// This is for compatibility with go-app HTML components.
func (c *blazarInputWrapper) Title(title string) *blazarInputWrapper {
	c.MixTitle(title)
	return c
}
