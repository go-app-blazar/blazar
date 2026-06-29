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
	IClasses []string
	ILabel   string
	IBody    []app.UI
}

var _ app.Composer = (*blazarInputWrapper)(nil)

func (c *blazarInputWrapper) Class(class ...string) *blazarInputWrapper {
	c.IClasses = class
	return c
}

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

	return app.Span().
		Class(append([]string{"blazar-input-wrapper"}, c.IClasses...)...).
		Body(
			body...,
		)
}
