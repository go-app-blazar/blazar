package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Icon() *blazarIcon {
	return &blazarIcon{}
}

type blazarIcon struct {
	app.Compo
	UseEvents
	ClassValue string
	IconValue  string
}

var _ app.Composer = (*blazarIcon)(nil)

func (c *blazarIcon) Class(class string) *blazarIcon {
	c.ClassValue = class
	return c
}

func (c *blazarIcon) Icon(icon string) *blazarIcon {
	c.IconValue = icon
	return c
}

func (c *blazarIcon) On(event string, function func(ctx app.Context, e app.Event)) *blazarIcon {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarIcon) Render() app.UI {
	return c.UseEvents.Wrap(
		app.Span().
			Class("blazar-icon", c.ClassValue).
			DataSet("icon", c.IconValue).
			Body(
				app.I().
					Class("blazar-icon__icon").
					Class("fa-solid").
					Class("fa-" + c.IconValue),
			),
	)
}
