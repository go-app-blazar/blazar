package blazar

import (
	"github.com/go-app-blazar/blazar/slot"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Page() *blazarPage {
	return &blazarPage{}
}

type blazarPage struct {
	app.Compo
	slot.Slotted
	MixHTML
}

var _ app.Composer = (*blazarPage)(nil)

// BEGIN: mixHTML methods

func (c *blazarPage) Attr(name string, value any) *blazarPage {
	c.mixAttr(name, value)
	return c
}

func (c *blazarPage) Class(class ...string) *blazarPage {
	c.mixClass(class...)
	return c
}

func (c *blazarPage) DataSet(name string, value any) *blazarPage {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarPage) Style(name, value string) *blazarPage {
	c.mixStyle(name, value)
	return c
}

func (c *blazarPage) Title(title string) *blazarPage {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarPage) Render() app.UI {
	element := app.Div().
		Class("blazar-page").
		Body(c.SlotContents()...)

	var finalElement app.UI
	finalElement = c.MixHTML.Wrap(element)

	return finalElement
}

func (c *blazarPage) Body(components ...app.UI) *blazarPage {
	c.Slotted.AddSlotContents(components...)
	return c
}
