package blazar

import (
	"strings"

	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/go-app-blazar/blazar/mix"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Icon() *blazarIcon {
	return &blazarIcon{}
}

type blazarIcon struct {
	app.Compo
	mix.MixEvents
	mix.MixHTML

	IconValue string
}

var _ app.Composer = (*blazarIcon)(nil)

// Icon sets the name of the icon.
//
// If this does not start with "fa-", it will be prefixed with "fa-".
//
// If this contains a space, the first part will be the icon class and the
// second part will be the font class (default is "fa-solid").
func (c *blazarIcon) Icon(icon string) *blazarIcon {
	c.IconValue = icon
	return c
}

func (c *blazarIcon) On(event string, function func(ctx app.Context, e app.Event)) *blazarIcon {
	c.MixEvents.On(event, function)
	return c
}

func (c *blazarIcon) OnClick(function func(ctx app.Context, e htmlevent.MouseEvent)) *blazarIcon {
	c.MixEvents.OnClick(function)
	return c
}

func (c *blazarIcon) Render() app.UI {
	fontClasses := strings.Split(c.IconValue, " ")

	var finalElement app.UI
	finalElement = c.MixEvents.Wrap(
		app.Span().
			Class("blazar-icon").
			DataSet("icon", c.IconValue).
			Body(
				app.I().
					Class("blazar-icon__icon").
					Class(fontClasses...),
			),
	)

	finalElement = c.MixHTML.Wrap(finalElement)

	return finalElement
}
