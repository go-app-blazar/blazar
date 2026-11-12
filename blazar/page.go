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
