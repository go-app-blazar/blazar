package blazar

import (
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Icon() *blazarIcon {
	return &blazarIcon{
		ISolid: true,
	}
}

type blazarIcon struct {
	app.Compo
	UseEvents
	MixHTML

	ISolid    bool
	IconValue string
}

var _ app.Composer = (*blazarIcon)(nil)

// BEGIN: mixHTML methods

func (c *blazarIcon) Attr(name string, value any) *blazarIcon {
	c.mixAttr(name, value)
	return c
}

func (c *blazarIcon) Class(class ...string) *blazarIcon {
	c.mixClass(class...)
	return c
}

func (c *blazarIcon) DataSet(name string, value any) *blazarIcon {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarIcon) Style(name, value string) *blazarIcon {
	c.mixStyle(name, value)
	return c
}

func (c *blazarIcon) Title(title string) *blazarIcon {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

func (c *blazarIcon) Icon(icon string) *blazarIcon {
	c.IconValue = icon
	return c
}

func (c *blazarIcon) Solid(solid bool) *blazarIcon {
	c.ISolid = solid
	return c
}

func (c *blazarIcon) On(event string, function func(ctx app.Context, e app.Event)) *blazarIcon {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarIcon) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.UseEvents.OnClick(function)
	return c
}

func (c *blazarIcon) Render() app.UI {
	solidClass := "fa-regular"
	if c.ISolid {
		solidClass = "fa-solid"
	}
	var finalElement app.UI
	finalElement = c.UseEvents.Wrap(
		app.Span().
			Class("blazar-icon").
			DataSet("icon", c.IconValue).
			Body(
				app.I().
					Class("blazar-icon__icon", solidClass, "fa-"+c.IconValue),
			),
	)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
