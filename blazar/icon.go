package blazar

import (
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
	IClasses  []string
	ISolid    bool
	IconValue string
}

var _ app.Composer = (*blazarIcon)(nil)

func (c *blazarIcon) Class(class ...string) *blazarIcon {
	c.IClasses = class
	return c
}

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

func (c *blazarIcon) Render() app.UI {
	solidClass := "fa-regular"
	if c.ISolid {
		solidClass = "fa-solid"
	}
	return c.UseEvents.Wrap(
		app.Span().
			Class(append([]string{"blazar-icon"}, c.IClasses...)...).
			DataSet("icon", c.IconValue).
			Body(
				app.I().
					Class("blazar-icon__icon", solidClass, "fa-"+c.IconValue),
			),
	)
}
