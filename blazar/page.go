package blazar

import (
	"github.com/go-app-blazar/blazar/slot"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Page() *blazarPage {
	return &blazarPage{
		IStyles: map[string]string{},
	}
}

type blazarPage struct {
	app.Compo
	slot.Slotted

	IClasses []string
	IStyles  map[string]string
}

var _ app.Composer = (*blazarPage)(nil)

func (c *blazarPage) Class(class ...string) *blazarPage {
	c.IClasses = class
	return c
}

func (c *blazarPage) Style(name, value string) *blazarPage {
	c.IStyles[name] = value
	return c
}

func (c *blazarPage) Render() app.UI {
	element := app.Div().
		Class(append([]string{"blazar-page"}, c.IClasses...)...)
	for name, value := range c.IStyles {
		element.Style(name, value)
	}
	return element.
		Body(c.SlotContents()...)
}

func (c *blazarPage) Body(components ...app.UI) *blazarPage {
	c.Slotted.AddSlotContents(components...)
	return c
}
