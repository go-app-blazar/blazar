package blazar

import (
	"github.com/go-app-blazar/blazar/mix"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Item() *blazarItem {
	return &blazarItem{}
}

type blazarItem struct {
	app.Compo
	mix.MixEvents
	mix.MixHTML

	IIcon  string
	ILabel string
	ITo    string
}

var _ app.Composer = (*blazarItem)(nil)

func (c *blazarItem) Icon(icon string) *blazarItem {
	c.IIcon = icon
	return c
}

func (c *blazarItem) Label(name string) *blazarItem {
	c.ILabel = name
	return c
}

func (c *blazarItem) To(to string) *blazarItem {
	c.ITo = to
	return c
}

func (c *blazarItem) Render() app.UI {
	var finalItem app.UI
	finalItem = app.A().
		Class("blazar-item").
		Href(c.ITo).
		Body(
			app.If(c.IIcon != "", func() app.UI {
				return app.Span().
					Class("blazar-item__icon").
					Body(
						Icon().Icon(c.IIcon),
					)
			}),
			app.Span().
				Class("blazar-item__name").
				Text(c.ILabel),
		)

	finalItem = c.MixHTML.Wrap(finalItem)
	finalItem = c.MixEvents.Wrap(finalItem)

	return finalItem
}
