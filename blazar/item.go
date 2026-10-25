package blazar

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

func Item() *blazarItem {
	return &blazarItem{}
}

type blazarItem struct {
	app.Compo
	UseEvents
	MixHTML

	IIcon  string
	ILabel string
	ITo    string
}

var _ app.Composer = (*blazarItem)(nil)

// BEGIN: mixHTML methods

func (c *blazarItem) Attr(name string, value any) *blazarItem {
	c.mixAttr(name, value)
	return c
}

func (c *blazarItem) Class(class ...string) *blazarItem {
	c.mixClass(class...)
	return c
}

func (c *blazarItem) DataSet(name string, value any) *blazarItem {
	c.mixDataSet(name, value)
	return c
}

func (c *blazarItem) Style(name, value string) *blazarItem {
	c.mixStyle(name, value)
	return c
}

func (c *blazarItem) Title(title string) *blazarItem {
	c.mixTitle(title)
	return c
}

// END: mixHTML methods

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

func (c *blazarItem) On(event string, function func(ctx app.Context, e app.Event)) *blazarItem {
	c.UseEvents.On(event, function)
	return c
}

func (c *blazarItem) Render() app.UI {
	var finalItem app.UI
	finalItem = app.A().
		Class("blazar-item").
		Href(c.ITo).
		Body(
			app.Span().
				Class("blazar-item__icon").
				Body(
					app.If(c.IIcon != "", func() app.UI {
						return Icon().Icon(c.IIcon)
					}),
				),
			app.Span().
				Class("blazar-item__name").
				Text(c.ILabel),
		)

	finalItem = c.MixHTML.Wrap(finalItem)
	finalItem = c.UseEvents.Wrap(finalItem)

	return finalItem
}
