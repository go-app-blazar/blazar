package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// InputWrapper is a wrapper around an input element.
//
// Use this component to create custom input elements with a label and a body.
func InputWrapper() *blazarInputWrapper {
	return &blazarInputWrapper{}
}

type blazarInputWrapper struct {
	app.Compo
	UseEvents
	MixHTML

	ILabel string
	IBody  []app.UI
}

var _ app.Composer = (*blazarInputWrapper)(nil)

// BEGIN: mixHTML methods

func (c *blazarInputWrapper) Attr(name string, value any) *blazarInputWrapper {
	c.mixAttr(name, value)
	return c
}

func (c *blazarInputWrapper) Class(class ...string) *blazarInputWrapper {
	c.mixClass(class...)
	return c
}

func (c *blazarInputWrapper) DataSet(name string, value any) *blazarInputWrapper {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarInputWrapper) Style(name, value string) *blazarInputWrapper {
	c.mixStyle(name, value)
	return c
}

func (c *blazarInputWrapper) Title(title string) *blazarInputWrapper {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarInputWrapper) Label(label string) *blazarInputWrapper {
	c.ILabel = label
	return c
}

func (c *blazarInputWrapper) Body(body ...app.UI) *blazarInputWrapper {
	c.IBody = body
	return c
}

func (c *blazarInputWrapper) Render() app.UI {
	var body []app.UI
	if c.ILabel != "" {
		body = append(body, app.Span().
			Class("blazar-input-wrapper__label").
			Text(c.ILabel))
	}
	body = append(body, app.Div().
		Class("blazar-input-wrapper__body").
		Body(
			c.IBody...,
		))

	var finalElement app.UI
	finalElement = app.Label().
		Class("blazar-input-wrapper").
		Body(
			body...,
		)

	finalElement = c.MixHTML.Wrap(finalElement)
	return finalElement
}
